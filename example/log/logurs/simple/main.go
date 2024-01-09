package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{"animal": "walrus"}).Info("A walrus appears")
	// INFO[0000] A walrus appears                              animal=walrus

}
