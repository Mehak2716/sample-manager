package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Mehak2716/sample-manager/cmd/pkg/app"
	"github.com/swiggy-private/gocommons/log"
)

func main() {
	// Setup the main application context that will be notified on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// Setup the error channel that will be used to notify the main loop of any errors
	errch := make(chan error, 1)

	// Start the main application
	app.Start(ctx, errch)

	// Wait till the context is done (by SIGINT or SIGTERM) or an error is received
	// from the goroutines
	select {
	case <-ctx.Done():
	case err := <-errch:
		log.Fatalw(ctx, "main loop failed", "err", err)
		close(errch)
	}
}
