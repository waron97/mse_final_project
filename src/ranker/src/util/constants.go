package util

import "os"

type Constants struct {
	logsAppName string
	logsApiKey  string
	appEnv      string
	logsUrl     string
	mongoUri    string
}

func GetConstants() Constants {
	return Constants{
		logsAppName: os.Getenv("LOGS_APP_NAME"),
		logsApiKey:  os.Getenv("LOGS_KEY"),
		appEnv:      os.Getenv("APP_ENV"),
		mongoUri:    os.Getenv("MONGO_URI"),
		logsUrl:     "http://logs:8080",
	}
}
