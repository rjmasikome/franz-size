package endpoints

import (
	"net/http"

	"go.uber.org/zap"
)

type Controllers struct {
	cfg    Config
	utils  *Utils
	logger *zap.Logger
}

func NewControllers(cfg Config, logger *zap.Logger, utils *Utils) *Controllers {
	return &Controllers{
		cfg:    cfg,
		utils:  utils,
		logger: logger,
	}
}

func (c *Controllers) Ok() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.utils.ResponseJSON(w, "ok")
	}
}
