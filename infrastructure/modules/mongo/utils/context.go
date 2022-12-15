package utils

import (
	"context"
	"time"
)

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return ctx, cancel
}