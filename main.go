package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"go.opentelemetry.io/otel"
	apiTrace "go.opentelemetry.io/otel/trace"
)

const name = "rolldice"

var tracer apiTrace.Tracer

func main() {
	ctx := context.Background()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	tracer = otel.Tracer(name)

	res := rolldice(ctx)
	fmt.Printf("got %d\n", res)
}

func rolldice(ctx context.Context) int {
	ctx, span := tracer.Start(ctx, "roll", apiTrace.WithSpanKind(apiTrace.SpanKindServer))
	defer span.End()

	// The ctx returned by tracer.Start() has trace information.
	// Pass the same ctx here so that any spans created in the called function
	// get linked to the current span.
	return rolldiceWork(ctx)
}

func rolldiceWork(ctx context.Context) int {
	ctx, span := tracer.Start(ctx, "roll_work")
	defer span.End()

	roll := 1 + rand.Intn(6)
	return roll
}
