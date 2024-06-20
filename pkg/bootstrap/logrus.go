package bootstrap

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogrusLogger ...
func NewLogrusLogger() (*logrus.Logger, error) {

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp:  true,
		DisableHTMLEscape: true,
		// PrettyPrint:       true,
	})
	return logger, nil
}
