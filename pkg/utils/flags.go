package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func SetEnvFlags(fs *flag.FlagSet, prefix string) (err error) {
	cmdFlags := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) {
		cmdFlags[f.Name] = true
	})

	fs.VisitAll(func(f *flag.Flag) {
		if !cmdFlags[f.Name] {
			key := strings.ToUpper(prefix + "_" + strings.Replace(f.Name, "-", "_", -1))
			val := os.Getenv(key)
			if val != "" {
				errs := fs.Set(f.Name, val)
				if errs != nil {
					err = fmt.Errorf("invalid value %q for %s: %v", val, key, errs)
				}
			}
		}
	})
	return err
}
