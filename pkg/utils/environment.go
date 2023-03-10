package utils

import (
	"github.com/elvenworks/go_unleash/pkg/domain/memory"
)

func GetPrefix() string {
	switch memory.APP_ENV_AppEnvironment {
	case "development":
		return "development"
	case "staging":
		return "staging"
	}
	return ""
}
