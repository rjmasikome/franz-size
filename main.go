package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"franz-size/endpoints"
	"franz-size/kafka"
	"franz-size/logging"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

func main() {
	startupLogger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to create startup logger: %w", err))
	}

	cfg, err := newConfig(startupLogger)
	if err != nil {
		startupLogger.Fatal("failed to parse config", zap.Error(err))
	}

	logger := logging.NewLogger(cfg.Logger)
	if err != nil {
		startupLogger.Fatal("failed to create new logger", zap.Error(err))
	}

	// Setup context that cancels when the application receives an interrupt signal
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			logger.Info("received a signal, going to shut down this service")
			cancel()
		case <-ctx.Done():
		}
	}()

	kgoOpts := []kgo.Opt{
		// TODO: this should be defined on the config
		kgo.ConsumeTopics("__consumer_offsets"),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	}

	// Create kafka service
	kafkaSvc := kafka.NewService(cfg.Kafka, logger)
	_, err = kafkaSvc.CreateAndTestClient(ctx, logger, kgoOpts)
	if err != nil {
		logger.Fatal("failed to setup minion service", zap.Error(err))
	}

	// TODO: Put this in one service module
	endpoints := endpoints.NewService(cfg.Endpoints, logger)
	endpoints.Propagate()
}
