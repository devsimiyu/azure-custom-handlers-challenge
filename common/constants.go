package common

import "os"

var (
	MongoUri string
)

func init() {
	if uri, ok := os.LookupEnv("MONGO_URI"); ok {
		MongoUri = uri
	} else {
		MongoUri = "mongodb+srv://root:4d3wZAMQwPvTJ8m@champeschoolcluster.6d0cant.mongodb.net/?retryWrites=true&w=majority&appName=ChampeSchoolCluster"
	}
}
