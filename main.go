package main

import (
	"context"
	"os/signal"
	"syscall"
)

var (
	Version = "N/A"
	Commit  = "N/A"
	BuiltAt = "N/A"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	_ = newCLI(ctx).Execute()
}
