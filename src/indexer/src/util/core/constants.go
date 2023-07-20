package core

import "os"

type Constants struct {
	LogsAppName string
	LogsApiKey  string
	AppEnv      string
	LogsUrl     string
	MongoUri    string
	BertUri     string

	StorageBaseDir        string
	StorageDocsDir        string
	StorageAverageDocsDir string
	StorageClustersDir    string
	StorageClusterMapPath string

	ClusterCount int
}

func GetConstants() Constants {
	const StorageBaseDir = "./offline-index"
	return Constants{
		LogsAppName: os.Getenv("LOGS_APP_NAME"),
		LogsApiKey:  os.Getenv("LOGS_KEY"),
		AppEnv:      os.Getenv("APP_ENV"),
		MongoUri:    os.Getenv("MONGO_URI"),
		LogsUrl:     "http://logs:8080",
		BertUri:     "http://bert:5000",

		StorageBaseDir:        StorageBaseDir,
		StorageDocsDir:        StorageBaseDir + "/documents",
		StorageAverageDocsDir: StorageBaseDir + "/average-documents",
		StorageClustersDir:    StorageBaseDir + "/clusters",
		StorageClusterMapPath: StorageBaseDir + "/clusters/clusterMap.index",

		ClusterCount: 100,
	}
}
