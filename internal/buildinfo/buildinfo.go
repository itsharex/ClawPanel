package buildinfo

import (
	"os"
	"strings"
)

var (
	Version = "dev"
	Edition = "pro"
)

func NormalizedEdition() string {
	edition := strings.TrimSpace(strings.ToLower(os.Getenv("CLAWPANEL_EDITION")))
	if edition == "" {
		edition = strings.TrimSpace(strings.ToLower(Edition))
	}
	switch edition {
	case "lite":
		return "lite"
	default:
		return "pro"
	}
}

func IsLite() bool {
	return NormalizedEdition() == "lite"
}
