package main

import (
	"context"
	"kafka-pet/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatalf("couldn't start app: %v", err)
	}
}
