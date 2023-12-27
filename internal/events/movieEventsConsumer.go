package events

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type movieEventsConsumer struct {
	r      *kafka.Reader
	logger *logrus.Logger
	repo   CastsRepository
}

type movieDeletedEvent struct {
	ID int32 `json:"movie_id"`
}

const (
	movieDeletedTopic = "movie_deleted"
)

func NewMovieEventsConsumer(cfg KafkaReaderConfig, logger *logrus.Logger, repo CastsRepository) *movieEventsConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          cfg.Brokers,
		GroupTopics:      []string{movieDeletedTopic},
		GroupID:          cfg.GroupID,
		Logger:           logger,
		ReadBatchTimeout: cfg.ReadBatchTimeout,
	})

	return &movieEventsConsumer{
		r:      r,
		logger: logger,
		repo:   repo,
	}
}

func (c *movieEventsConsumer) Run(ctx context.Context) {
	for {
		select {
		default:
			c.Concume(ctx)
		case <-ctx.Done():
			c.logger.Info("movie events consumer shutting down")
			c.r.Close()
			c.logger.Info("movie events consumer shutted down")
			return
		}
	}
}

func (c *movieEventsConsumer) Concume(ctx context.Context) {
	message, err := c.r.FetchMessage(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}

	switch message.Topic {
	case movieDeletedTopic:
		err = c.ConcumeDeletedEvent(ctx, message)
	}
	if err != nil {
		c.logger.Error(err)
		return
	}

	if err := c.r.CommitMessages(ctx, message); err != nil {
		c.logger.Error(err)
	}
}

func (c *movieEventsConsumer) ConcumeDeletedEvent(ctx context.Context, msg kafka.Message) error {
	var event movieDeletedEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		return err
	}
	return c.repo.DeleteCast(ctx, event.ID)
}
