package endpoints

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Utils struct {
	cfg    Config
	logger *zap.Logger
}

func NewUtils(cfg Config, logger *zap.Logger) *Utils {
	return &Utils{
		cfg:    cfg,
		logger: logger,
	}
}

func (u *Utils) ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
