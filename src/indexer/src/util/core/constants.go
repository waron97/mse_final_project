package core

import "os"

type Constants struct {
	LogsAppName string
	LogsApiKey  string
	AppEnv      string
	LogsUrl     string
	MongoUri    string

	StorageBaseDir        string
	StorageDocsDir        string
	StorageAverageDocsDir string
}

func GetConstants() Constants {
	const StorageBaseDir = "./offline-index"
	return Constants{
		LogsAppName: os.Getenv("LOGS_APP_NAME"),
		LogsApiKey:  os.Getenv("LOGS_KEY"),
		AppEnv:      os.Getenv("APP_ENV"),
		MongoUri:    os.Getenv("MONGO_URI"),
		LogsUrl:     "http://logs:8080",

		StorageBaseDir:        StorageBaseDir,
		StorageDocsDir:        StorageBaseDir + "/documents",
		StorageAverageDocsDir: StorageBaseDir + "/average-documents",
	}
}
