package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

var logger *zap.Logger

func init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create elastic client: %v", err))
	}

	consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	elasticHook := NewElasticHook(client, "apcore_logs", "localhost")
	elasticCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(elasticHook), zapcore.DebugLevel)

	core := zapcore.NewTee(consoleCore, elasticCore)

	logger = zap.New(core, zap.AddCaller())
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Sync() {
	_ = logger.Sync()
}
