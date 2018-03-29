package client

import (
	"context"
	stdlog "log"
)

type DiffCalculator interface {
	Calculate(ctx context.Context, sourceID string, items []Item) error
}

type Item struct {
	ID   string
	Data interface{}
}

var Std DiffCalculator

func init() {
	var err error
	if Std, err = New(); err != nil {
		stdlog.Fatalf("Error creating DiffCalculator client: %s", err)
	}
}
