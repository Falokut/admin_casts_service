package events

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type eventsConsumer struct {
	r      *kafka.Reader
	logger *logrus.Logger
	repo   CastsRepository
}

const (
	movieDeletedTopic  = "movie_deleted"
	personDeletedTopic = "person_deleted"
)

func NewEventsConsumer(cfg KafkaReaderConfig, logger *logrus.Logger, repo CastsRepository) *eventsConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          cfg.Brokers,
		GroupTopics:      []string{movieDeletedTopic, personDeletedTopic},
		GroupID:          cfg.GroupID,
		Logger:           logger,
		ReadBatchTimeout: cfg.ReadBatchTimeout,
	})

	return &eventsConsumer{
		r:      r,
		logger: logger,
		repo:   repo,
	}
}

func (c *eventsConsumer) Run(ctx context.Context) {
	for {
		select {
		default:
			c.Concume(ctx)
		case <-ctx.Done():
			c.logger.Info("events consumer shutting down")
			c.r.Close()
			c.logger.Info("events consumer shutted down")
			return
		}
	}
}

func (c *eventsConsumer) Concume(ctx context.Context) {
	message, err := c.r.FetchMessage(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}

	switch message.Topic {
	case personDeletedTopic:
		err = c.ConcumePersonDeletedEvent(ctx, &message)
	case movieDeletedTopic:
		err = c.ConcumeMovieDeletedEvent(ctx, &message)
	}

	if err != nil {
		c.logger.Error(err)
		return
	}

	if err := c.r.CommitMessages(ctx, message); err != nil {
		c.logger.Error(err)
	}
}

func (c *eventsConsumer) ConcumeMovieDeletedEvent(ctx context.Context, msg *kafka.Message) error {
	var event struct {
		ID int32 `json:"movie_id"`
	}
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		return err
	}
	return c.repo.DeleteCast(ctx, event.ID)
}

func (c *eventsConsumer) ConcumePersonDeletedEvent(ctx context.Context, msg *kafka.Message) error {
	var event struct {
		ID int32 `json:"person_id"`
	}
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		return err
	}

	return c.repo.RemovePersonFromCasts(ctx, event.ID)
}
