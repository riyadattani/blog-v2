package main

import (
	"context"
	"os"
	"os/signal"
)

const (
	exitCodeInterrupt = 2
)

// https://pace.dev/blog/2020/02/17/repond-to-ctrl-c-interrupt-signals-gracefully-with-context-in-golang-by-mat-ryer.html
// this creates an application context which will emit a cancel signal when an interrupt it sent or the application ends.
func listenForCancellationAndAddToContext() (ctx context.Context, done func()) {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	done = func() {
		signal.Stop(signalChan)
		cancel()
	}

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()

	return ctx, done
}
