package common

import "os"

var (
	MongoUri string
)

func init() {
	if uri, ok := os.LookupEnv("MONGO_URI"); ok {
		MongoUri = uri
	} else {
		MongoUri = "mongodb+srv://<username>:<password>@<host>/?retryWrites=true&w=majority&appName=<cluster>"
	}
}
