package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sqlCloudDNS            string
	err                    error
	db                     *gorm.DB
	dbPwd                  string
	instanceConnectionName string
	usePrivate             string
	dbUser                 string
)

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func connectWithConnector() (*sql.DB, error) {
	ctx, cacnel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cacnel()
	go func() {
		<-ctx.Done()
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("context deadline exceeded")
			panic("context deadline exceeded")
		}
	}()
	dbName := "auth"

	d, err := cloudsqlconn.NewDialer(ctx)
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)
	log.Println("dbURI=", dbURI)
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}

func connect() error {
	sqlDB, err := connectWithConnector()
	db, err = gorm.Open(gmysql.New(gmysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Println(err, "connect GORM error")
		return err
	}

	return nil
}

func init() {
	sqlCloudDNS, err = getSecretKey("projects/243723541767/secrets/sql_dns/versions/latest")
	dbPwd, err = getSecretKey("projects/243723541767/secrets/sql_pw/versions/latest")
	instanceConnectionName, err = getSecretKey("projects/243723541767/secrets/sql_cname/versions/latest")
	usePrivate, err = getSecretKey("projects/243723541767/secrets/sql_ip_private/versions/latest")
	dbUser, err = getSecretKey("projects/243723541767/secrets/sql_user/versions/latest")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		err := connect()
		if err != nil {
			fmt.Fprint(w, "connect error ", err.Error())
			return
		}

		fmt.Fprint(w, "connect success")
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := new(User)
	err = json.Unmarshal(b, user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = db.Create(user).Error
	if err != nil {
		http.Error(w, "Create user error", http.StatusInternalServerError)
		return
	}

	log.Println("create user success", user)
	fmt.Fprintln(w, "create user success", user)
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path=", r.URL.Path)
	switch r.URL.Path {
	case "/create-sql":
		createUserHandler(w, r)
		return
	}

	fmt.Fprint(w, "Hello, World!", r.URL.Path)
}

func getSecretKey(key string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create kms client: %w", err)
	}
	defer client.Close()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: key,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", err
	}

	// Print the secret payload.
	//
	// WARNING: Do not print the secret in a production environment - this
	// snippet is showing how to access the secret material.
	log.Printf("key: %s\n Plaintext: %s", key, result.Payload.Data)
	return string(result.Payload.Data), nil
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
