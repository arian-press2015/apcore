package logger

import (
	"apcore/config"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
	var entry map[string]interface{}
	err = json.Unmarshal(p, &entry)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal log entry: %w", err)
	}

	entry["host"] = hook.host
	entry["timestamp"] = time.Now()

	_, err = hook.client.Index().
		Index(hook.index).
		BodyJson(entry).
		Do(context.Background())

	if err != nil {
		return 0, fmt.Errorf("failed to index log entry: %w", err)
	}

	return len(p), nil
}

func NewElasticClient(cfg *config.Config) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.Elastic.Url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
		elastic.SetBasicAuth(cfg.Elastic.Username, cfg.Elastic.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create elastic client: %w", err)
	}
	return client, nil
}

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(client *elastic.Client, cfg *config.Config) (*Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	elasticHook := NewElasticHook(client, cfg.Elastic.Index, "localhost")
	elasticCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(elasticHook), zapcore.DebugLevel)

	fileSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     1, // days
		Compress:   true,
	})

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileSyncer,
		zapcore.DebugLevel,
	)

	core := zapcore.NewTee(consoleCore, elasticCore, fileCore)

	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
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
