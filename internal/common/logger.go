package common

import (
	"todo/internal/config"
	"todo/pkg/logging"
)

func GetLogger() *logging.Logger {
	mode := logging.Prod

	switch cfg := config.GetConfig(); cfg.Project.Mode {
	case "local":
		mode = logging.Local
	case "dev":
		mode = logging.Dev
	}

	return logging.GetLogger(mode)
}
