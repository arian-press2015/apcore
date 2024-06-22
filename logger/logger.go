package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Options(
	fx.Provide(NewElasticClient),
	fx.Provide(NewLogger),
)

type ElasticHook struct {
	client *elastic.Client
	index  string
	host   string
}

func NewElasticHook(client *elastic.Client, index string, host string) *ElasticHook {
	return &ElasticHook{
		client: client,
		index:  index,
		host:   host,
	}
}

func (hook *ElasticHook) Write(p []byte) (n int, err error) {
	entry := map[string]interface{}{
		"message":   string(p),
		"host":      hook.host,
		"timestamp": time.Now(),
	}

	_, err = hook.client.Index().
		Index(hook.index).
		BodyJson(entry).
		Do(context.Background())

	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func NewElasticClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create elastic client: %w", err)
	}
	return client, nil
}

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(client *elastic.Client) (*Logger, error) {
	consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	elasticHook := NewElasticHook(client, "apcore_logs", "localhost")
	elasticCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(elasticHook), zapcore.DebugLevel)

	core := zapcore.NewTee(consoleCore, elasticCore)

	zapLogger := zap.New(core, zap.AddCaller())
	return &Logger{zapLogger}, nil
}

func (l *Logger) Sync() {
	_ = l.zapLogger.Sync()
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}
