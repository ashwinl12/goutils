package goutils

import "go.uber.org/zap"

func LogError(logger *zap.Logger, msg string, obj map[string]interface{}) {
	fields := make([]zap.Field, 0, len(obj))
	for key, val := range obj {
		fields = append(fields, zap.Any(key, val))
	}
	logger.Error(msg, fields...)
}

func LogInfo(logger *zap.Logger, msg string, obj map[string]interface{}) {
	fields := make([]zap.Field, 0, len(obj))
	for key, val := range obj {
		fields = append(fields, zap.Any(key, val))
	}
	logger.Info(msg, fields...)
}

func LogDebug(logger *zap.Logger, msg string, obj map[string]interface{}) {
	fields := make([]zap.Field, 0, len(obj))
	for key, val := range obj {
		fields = append(fields, zap.Any(key, val))
	}
	logger.Debug(msg, fields...)
}

func LogPanic(logger *zap.Logger, msg string, obj map[string]interface{}) {
	fields := make([]zap.Field, 0, len(obj))
	for key, val := range obj {
		fields = append(fields, zap.Any(key, val))
	}
	logger.Panic(msg, fields...)
}