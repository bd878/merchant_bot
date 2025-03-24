package merchant_bot

import "go.uber.org/zap"

var log *zap.SugaredLogger = (*zap.SugaredLogger)(nil)

func init() {
	log = NewProduction()
}

func NewProduction(options ...zap.Option) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return zap.Must(config.Build(options...)).Sugar()
}

func NewDevelopment(options ...zap.Option) *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return zap.Must(config.Build(options...)).Sugar()
}

func Log() *zap.SugaredLogger {
	return log
}