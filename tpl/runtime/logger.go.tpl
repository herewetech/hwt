/*
 * Copyright (C) ###__PROJ_AUTHOR__### - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file logger.go
 * @package runtime
 * @author ###__PROJ_AUTHOR__###
 * @since ###__TODAY__###
 */

package runtime

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LoggerRaw, _ = zap.NewProduction(
	zap.AddStacktrace(zapcore.ErrorLevel),
)
var Logger = LoggerRaw.Sugar()

func InitLogger() error {
	loggerCfg := zap.NewProductionConfig()
	if Config.Debug {
		loggerCfg.Level.SetLevel(zapcore.DebugLevel)
	} else {
		loggerCfg.Level.SetLevel(zapcore.ErrorLevel)
	}

	LoggerRaw = zap.Must(loggerCfg.Build())
	Logger = LoggerRaw.Sugar()

	// Nothing over an error
	return nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
