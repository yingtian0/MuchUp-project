package server

import (
	"MuchUp/backend/config"
	"MuchUp/backend/pkg/logger"
)

func StartGRPCserver(cfg *config.Config, appLogger logger.Logger) {
	_ = cfg
	_ = appLogger
}
