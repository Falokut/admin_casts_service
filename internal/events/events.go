package events

import (
	"context"
	"time"
)

type CastsRepository interface {
	DeleteCast(ctx context.Context, id int32) error
	RemoveActorFromCasts(ctx context.Context, actorID int32) (err error)
}

type KafkaWriterConfig struct {
	Brokers []string
}

type KafkaReaderConfig struct {
	Brokers          []string
	GroupID          string
	ReadBatchTimeout time.Duration
}
