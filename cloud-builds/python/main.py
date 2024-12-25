from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello():
	return "Hello, world!, CI build V1.0"

if __name__ == "__main__":
	app.run(debug=True)
