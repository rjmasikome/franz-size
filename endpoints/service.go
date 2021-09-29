package endpoints

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Service struct {
	Cfg         Config
	logger      *zap.Logger
	utils       *Utils
	controllers *Controllers
}

func NewService(cfg Config, logger *zap.Logger) *Service {

	utils := NewUtils(cfg, logger)
	controllers := NewControllers(cfg, logger, utils)

	return &Service{
		Cfg:         cfg,
		utils:       utils,
		controllers: controllers,
		logger:      logger.Named("kafka_service"),
	}
}

func (s *Service) Propagate() {
	router := mux.NewRouter()

	//Pre-flight Endpoints
	router.PathPrefix("/").HandlerFunc(s.controllers.Ok()).Methods("OPTIONS")
	router.PathPrefix("/").HandlerFunc(s.controllers.Ok()).Methods("GET")

	// Unprotected Endpoints
	router.HandleFunc("/metrics", s.controllers.Ok()).Methods("GET")
}
