package main

import (
	logrus "github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("HelloWorld!!")
	logrus.Warning("Warning")
	logrus.Error("Error")
	logrus.Fatal("Fatal")
	logrus.Debug("debug")

}
