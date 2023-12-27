package events

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type personEventsConsumer struct {
	r      *kafka.Reader
	logger *logrus.Logger
	repo   CastsRepository
}

type personDeletedEvent struct {
	ID int32 `json:"person_id"`
}

const (
	personDeletedTopic = "person_deleted"
)

func NewPersonEventsConsumer(cfg KafkaReaderConfig, logger *logrus.Logger,
	repo CastsRepository) *personEventsConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          cfg.Brokers,
		GroupTopics:      []string{personDeletedTopic},
		GroupID:          cfg.GroupID,
		Logger:           logger,
		ReadBatchTimeout: cfg.ReadBatchTimeout,
	})

	return &personEventsConsumer{
		r:      r,
		logger: logger,
		repo:   repo,
	}
}

func (c *personEventsConsumer) Run(ctx context.Context) {
	for {
		select {
		default:
			c.Concume(ctx)
		case <-ctx.Done():
			c.logger.Info("person events consumer shutting down")
			c.r.Close()
			c.logger.Info("person events consumer shutted down")
			return
		}
	}
}

func (c *personEventsConsumer) Concume(ctx context.Context) {
	message, err := c.r.FetchMessage(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}

	switch message.Topic {
	case personDeletedTopic:
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

func (c *personEventsConsumer) ConcumeDeletedEvent(ctx context.Context, msg kafka.Message) error {
	var event personDeletedEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		return err
	}
	return c.repo.RemoveActorFromCasts(ctx, event.ID)
}
