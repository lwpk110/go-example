package trace

import (
	"fmt"
	"os"
	"strconv"
	"time"

	sotlp "github.com/gunsluo/go-example/opentelemetry/demo/pkg/otlp"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/collector/translator/conventions"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/exporters/otlp"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	// environment variable names
	envServiceName   = "TRACE_SERVICE_NAME"
	envEnabled       = "TRACE_ENABLED"
	envAgentEndpoint = "OTLP_AGENT_EDNPOINT"

	OTEL_AGENT_ENDPOINT
)

// Configuration configures and creates Tracer
type Configuration struct {
	// ServiceName specifies the service name to use on the tracer.
	// Can be provided via environment variable named TRACE_SERVICE_NAME
	ServiceName string

	// AgentEndpoint is agent address of collecting trace message, host:port
	// Can be provided via environment variable named OTLP_AGENT_EDNPOINT
	AgentEndpoint string

	// Enabled can be provided via environment variable named TRACE_ENABLED
	Enabled bool
}

// FromEnv uses environment variables and overrides existing tracer's Configuration
func (c *Configuration) FromEnv() (*Configuration, error) {
	if e := os.Getenv(envServiceName); e != "" {
		c.ServiceName = e
	}

	if e := os.Getenv(envEnabled); e != "" {
		if value, err := strconv.ParseBool(e); err == nil {
			c.Enabled = value
		} else {
			return nil, fmt.Errorf("cannot parse env var %s=%s, %w", envEnabled, e, err)
		}
	}

	if e := os.Getenv(envAgentEndpoint); e != "" {
		c.AgentEndpoint = e
	}

	return c, nil
}

// Option is a function that sets some option on the client.
type Option func(c *Options)

// Options control behavior of the client.
type Options struct {
	ServiceName string
	Logger      *zap.Logger
}

// ServiceName can be provided to service name.
func ServiceName(serviceName string) Option {
	return func(c *Options) {
		c.ServiceName = serviceName
	}
}

// Logger can be provided to logger.
func Logger(logger *zap.Logger) Option {
	return func(c *Options) {
		c.Logger = logger
	}
}

// NewTracer create a new tracer
func (c *Configuration) NewTracer(options ...Option) (trace.Tracer, error) {
	if !c.Enabled {
		return &trace.NoopTracer{}, nil
	}

	if c.AgentEndpoint == "" {
		return nil, fmt.Errorf("missing agent address, please set environment variable %s", envAgentEndpoint)
	}

	opts := applyOptions(options...)
	if opts.ServiceName != "" {
		c.ServiceName = opts.ServiceName
	}

	exporter := sotlp.SingletonExporter()
	if exporter == nil {
		exp, err := otlp.NewExporter(otlp.WithInsecure(),
			otlp.WithAddress(c.AgentEndpoint),
			otlp.WithReconnectionPeriod(time.Minute),
			otlp.WithGRPCDialOption(grpc.WithTimeout(5*time.Second)))
		if err != nil {
			return nil, fmt.Errorf("failed to create the collector exporter: %w", err)
		}
		exporter = exp
		sotlp.SetExporter(exporter)
		opts.Logger.With(zap.String("agentEndpoint", c.AgentEndpoint)).Info("success to otlp agent")
	}
	// exporter.Stop()

	tags := []label.KeyValue{
		// the service name used to display traces
		label.Key(conventions.AttributeServiceName).String(c.ServiceName),
	}
	if hostname, err := os.Hostname(); err == nil {
		tags = append(tags, label.Key(conventions.AttributeHostName).String(hostname))
	}
	if ip, err := HostIP(); err == nil {
		tags = append(tags, label.Key(conventions.AttributeNetHostIP).String(ip.String()))
	}

	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithResource(resource.New(tags...)),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		return nil, fmt.Errorf("failed to creating trace provider: %w", err)
	}

	return tp.Tracer(c.ServiceName), nil
}

func applyOptions(options ...Option) Options {
	opts := Options{Logger: zap.NewNop()}

	for _, option := range options {
		option(&opts)
	}

	return opts
}

func (c *Configuration) NewHttpMiddleware(options ...HttpOption) (*HttpMiddleware, error) {
	//if !c.Enabled {
	//	return &HttpMiddleware{}, nil
	//}

	tracer, err := c.NewTracer(ServiceName(c.ServiceName))
	if err != nil {
		return &HttpMiddleware{}, err
	}

	m, err := NewHttpMiddleware(tracer, options...)
	if err != nil {
		return m, err
	}

	return m, nil
}

func (c *Configuration) NewTransport(options ...TransportOption) (*Transport, error) {
	tracer, err := c.NewTracer(ServiceName(c.ServiceName))
	if err != nil {
		return &Transport{}, err
	}

	t, err := NewTransport(tracer, options...)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (c *Configuration) OpenDB(dsn string) (*sqlx.DB, error) {
	tracer, err := c.NewTracer(ServiceName(c.ServiceName))
	if err != nil {
		return OpenDB(nil, dsn)
	}

	db, err := OpenDB(tracer, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
