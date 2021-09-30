package datapoints

import (
	"context"
	"fmt"
	"franz-size/kafka"
	"time"

	"github.com/nakabonne/tstorage"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

type DataPoints struct {
	Cfg     Config
	kafka   *kafka.Service
	storage tstorage.Storage
	logger  *zap.Logger
}

func New(cfg Config, kafka *kafka.Service, logger *zap.Logger) *DataPoints {

	storage, _ := tstorage.NewStorage(
		tstorage.WithTimestampPrecision(tstorage.Seconds),
	)
	defer storage.Close()

	return &DataPoints{
		Cfg:     cfg,
		kafka:   kafka,
		storage: storage,
		logger:  logger,
	}
}

func (d *DataPoints) GetMetric(metric string, labels []tstorage.Label, start int64, end int64) {
	points, _ := d.storage.Select(metric, labels, 1600000000, 1600000001)
	for _, p := range points {
		fmt.Printf("timestamp: %v, value: %v\n", p.Timestamp, p.Value)
	}
}

func (d *DataPoints) GetAndStoreMetadata(ctx context.Context, client *kgo.Client) error {
	metadata, err := d.kafka.GetMetadata(ctx, client, []string{"test", "orders"})
	if err != nil {
		return err
	}

	return nil
}

func (d *DataPoints) PollAndStore(ctx context.Context, client *kgo.Client) {

	pollMetadata := time.NewTicker(d.Cfg.ProbeInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-pollMetadata.C:
			err := d.GetAndStoreMetadata(ctx, client)
			if err != nil {
				d.logger.Error("failed to validate end-to-end topic", zap.Error(err))
			}
		}
	}

	metric := "size"
	labels := []tstorage.Label{
		{Name: "topic", Value: "orders"},
	}

	_ = d.storage.InsertRows([]tstorage.Row{
		{
			Metric:    metric,
			Labels:    labels,
			DataPoint: tstorage.DataPoint{Timestamp: time.Now().Unix(), Value: topic.Bytes()},
		},
	})
}

// func (s *Service) startConsumeMessages(ctx context.Context, initializedCh chan<- bool) {
// 	client := s.client

// 	s.logger.Info("Starting to consume end-to-end topic",
// 		zap.String("topic_name", s.config.TopicManagement.Name),
// 		zap.String("group_id", s.groupId))

// 	isInitialized := false
// 	for {
// 		fetches := client.PollFetches(ctx)
// 		if !isInitialized {
// 			isInitialized = true
// 			initializedCh <- true
// 		}

// 		// Log all errors and continue afterwards as we might get errors and still have some fetch results
// 		errors := fetches.Errors()
// 		for _, err := range errors {
// 			s.logger.Error("kafka fetch error",
// 				zap.String("topic", err.Topic),
// 				zap.Int32("partition", err.Partition),
// 				zap.Error(err.Err))
// 		}

// 		fetches.EachRecord(s.processMessage)
// 	}
// }
