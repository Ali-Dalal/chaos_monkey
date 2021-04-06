package utils

import "github.com/sirupsen/logrus"

func InitLogrusLog()  {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
}

func LogInfo(args ... interface{})  {
	logrus.Info(args...)
}

func LogError(args ...interface{})  {
	logrus.Error(args...)
}
