package utils

import (
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	log.Logger = log.With().Caller().Logger()
}
