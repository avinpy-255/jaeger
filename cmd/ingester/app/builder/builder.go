// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package builder

import (
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/cmd/ingester/app"
	"github.com/jaegertracing/jaeger/cmd/ingester/app/consumer"
	"github.com/jaegertracing/jaeger/cmd/ingester/app/processor"
	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/storage/v1/api/spanstore"
)

// CreateConsumer creates a new span consumer for the ingester
func CreateConsumer(logger *zap.Logger, metricsFactory metrics.Factory, spanWriter spanstore.Writer, options app.Options) (*consumer.Consumer, error) {
	spParams := processor.SpanProcessorParams{
		Writer: spanWriter,
	}
	proc := processor.NewSpanProcessor(spParams)

	return consumer.New(logger, metricsFactory, options, proc)
}
