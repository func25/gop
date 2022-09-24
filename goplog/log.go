package goplog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	_FEATURE = "feature"
	_OP      = "op"
)

func Log(feature string, oper string, subscribers ...string) zerolog.Logger {
	logger := log.Logger.With().Str(_FEATURE, feature).Str(_OP, oper)
	for i := range subscribers {
		logger = logger.Bool(subscribers[i], true)
	}

	return logger.Logger()
}
