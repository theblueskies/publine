package api

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// Service defines the interface to get rates for a given time range
type DBClean interface {
	DeleteExpiredItems(context.Context, events.CloudWatchEvent) error
}
