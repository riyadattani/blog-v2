package metrics

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
)

const (
	greetingsCounterName = "go_walking_skeleton_total_greetings"

	statusKey     = attribute.Key("status")
	statusSuccess = "success"
	statusFail    = "fail"
)

type client struct {
	greetingsCounter syncint64.Counter
}

func NewClient() (*client, error) {
	meter := global.MeterProvider().Meter("blog-v2/app-metrics-client")
	greetingsCounter, err := meter.SyncInt64().Counter(greetingsCounterName, instrument.WithDescription("A metric that records the number of greetings"))
	if err != nil {
		return nil, fmt.Errorf("failed to create %s metric - %w", greetingsCounterName, err)
	}

	return &client{greetingsCounter: greetingsCounter}, nil
}

func (c *client) GreetingSucceeded(ctx context.Context) {
	c.greetingsCounter.Add(ctx, 1, statusKey.String(statusSuccess))
}

func (c *client) GreetingFailed(ctx context.Context) {
	c.greetingsCounter.Add(ctx, 1, statusKey.String(statusFail))
}
