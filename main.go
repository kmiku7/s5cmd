package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kmiku7/s5cmd/command"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := command.Main(ctx, os.Args); err != nil {
		os.Exit(1)
	}
}
