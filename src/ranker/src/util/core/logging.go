package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Logger struct{}

type Log struct {
	Level    string      `json:"level"`
	Location string      `json:"location"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	AppId    string      `json:"appId"`
}

func (l Logger) isAlive() bool {
	constants := GetConstants()
	_, err := http.Get(constants.LogsUrl)
	return err == nil
}

func (l Logger) send(level string, location string, message string, data interface{}) bool {
	constants := GetConstants()

	for !l.isAlive() {
		time.Sleep(1 * time.Second)
	}

	body := Log{
		Level:    level,
		Location: location,
		Message:  message,
		Data:     data,
		AppId:    constants.LogsAppName,
	}

	payload, marshalErr := json.Marshal(body)

	if marshalErr != nil {
		return false
	}

	req, err := http.NewRequest("POST", constants.LogsAppName+"/logs", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "apiKey "+constants.LogsApiKey)

	if err != nil {
		return false
	}

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return false
	}

	defer res.Body.Close()

	return true
}

func (l Logger) Debug(location string, message string, data interface{}) bool {
	return l.send("debug", location, message, data)
}

func (l Logger) Info(location string, message string, data interface{}) bool {
	return l.send("info", location, message, data)
}

func (l Logger) Warn(location string, message string, data interface{}) bool {
	return l.send("warn", location, message, data)
}

func (l Logger) Error(location string, message string, data interface{}) bool {
	return l.send("error", location, message, data)
}

func (l Logger) Critical(location string, message string, data interface{}) bool {
	return l.send("critical", location, message, data)
}

func GetLogger() Logger {
	return Logger{}
}
