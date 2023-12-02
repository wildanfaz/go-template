package pkg

import "github.com/sirupsen/logrus"

func InitLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	log.SetReportCaller(true)

	return log
}
