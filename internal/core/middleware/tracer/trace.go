package tracer

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"io"
	"strings"
)

// this file is used to replace ginOTel but also can use by ginOTel

const (
	TracerKey   = "otel-tracer"
	TraceCtxKey = "X-Tracer-ID"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.GetTracerProvider().Tracer("gin-rush-template")
		spanName := c.Request.Method + " " + c.Request.URL.Path
		c.Set(TracerKey, tracer)
		savedCtx := c.Request.Context()
		defer func() {
			c.Request = c.Request.WithContext(savedCtx)
		}()

		ctx := otel.GetTextMapPropagator().Extract(savedCtx, propagation.HeaderCarrier(c.Request.Header))
		opts := []oteltrace.SpanStartOption{
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			oteltrace.WithAttributes(
				semconv.HTTPRequestMethodKey.String(c.Request.Method),
				semconv.URLPath(c.Request.URL.Path),
				semconv.HostID(c.Request.Host),
				semconv.NetworkProtocolName(c.Request.Proto),
				semconv.HostIPKey.String(c.ClientIP()),
			),
		}

		ctx, span := tracer.Start(ctx, spanName, opts...)

		for name, values := range c.Request.Header {
			span.SetAttributes(attribute.String("http.header."+name, strings.Join(values, ", ")))
		}

		traceID := span.SpanContext().TraceID().String()
		c.Writer.Header().Set(TraceCtxKey, traceID)
		//hub.Log = hub.Log.With(zap.String(TraceCtxKey, traceID))
		//hub.Log.Info("-------------------------------------------")
		defer span.End()

		var body []byte
		if c.Request.Body != nil {
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(c.Request.Body)
			if err != nil {
				span.SetStatus(codes.Error, "Failed to read request body")
			}
			body = buf.Bytes()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		span.SetAttributes(attribute.String("http.request.body", string(body)))

		c.Request = c.Request.WithContext(ctx)
		c.Next()
		status := c.Writer.Status()
		span.SetAttributes(attribute.Int("http.status_code", status))
	}
}
