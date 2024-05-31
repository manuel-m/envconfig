package envconfig

import (
	"fmt"
	"os"
)

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

const ModeDefault = Development

var Mode = ModeDefault

func Load() error {
	if value, exists := os.LookupEnv("ENV"); exists {
		if value == Development || value == Staging || value == Production {
			Mode = value
			return nil
		}
		return fmt.Errorf("invalid mode: %s", value)

	}

	return nil

}
