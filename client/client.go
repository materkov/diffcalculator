package client

import (
	"context"
	"github.com/go-kit/kit/log"
)

type DiffCalculator interface {
	Calculate(ctx context.Context, sourceID string, items []Item) error
}

type Item struct {
	ID   int
	Data interface{}
}

var Std DiffCalculator

func init() {
	var err error
	if Std, err = New(); err != nil {
		log.Fatalf("Error creating DiffCalculator client: %s", err)
	}
}
