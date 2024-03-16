package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

// для заменяемости логгера возвращаться должен интерфейс
// Как дз - реализуй интерфейс логгера

func SetupLogger(logFilePath string) (*logrus.Logger, *os.File, error) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger, file, nil
}
