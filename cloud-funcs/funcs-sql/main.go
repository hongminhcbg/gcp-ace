package helloworld

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

// IMPORTANT: Not recommended for production use.
// need to add trusted network IP to cloud SQL to connect
// connect with unix socket client
func connectSQLWithPasswordAndPublicIP(w http.ResponseWriter, r *http.Request) {
	password := os.Getenv("SQL_PW")
	user := os.Getenv("SQL_UNAME")
	instanceName := os.Getenv("INSTANCE_NAME")
	socketPath := fmt.Sprintf("/cloudsql/%s", instanceName) // get from gcloud sql instances describe

	connectionString := fmt.Sprintf("%s:%s@unix(%s)/users?parseTime=true", user, password, socketPath)
	log.Println("Connecting to SQL with public IP and password", connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error opening database: %v", err)))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error ping database: %v", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Connected to SQL with public IP and password"))
}

func connectSQLWithSQLProxy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("connect with sql proxy"))
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func helloHTTP(w http.ResponseWriter, r *http.Request) {
	sqlType := r.URL.Query().Get("sql_type")
	switch sqlType {
	case "public_ip":
		connectSQLWithPasswordAndPublicIP(w, r)
	default:
		connectSQLWithSQLProxy(w, r)
	}
}
