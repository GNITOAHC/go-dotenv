package dotenv

import (
	"fmt"
	"os"
)

// Must checks if the environment variables in the slice m are set.
func Must(m []string, paths ...string) error {
	allPath := append([]string{".env"}, paths...)
	Load(allPath...)

	for _, _m := range m {
		if os.Getenv(_m) == "" {
			panic(fmt.Sprintf("Environment variable %s is not set", _m))
		}
	}

	return nil
}
