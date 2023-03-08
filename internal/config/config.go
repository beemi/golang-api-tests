package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
)

var pattern, _ = regexp.Compile(`{{([^}]+)}}`)

var workdir = "."

// init() is a magic method called when this package is loaded.
func init() {
	var exist bool

	if w, exists := os.LookupEnv("WORK_DIR"); exists {
		workdir = w
	}
	fmt.Println("config.init() called")
	// Only auto init if it is in real deployment
	if _, exist = os.LookupEnv("ENV"); !exist {
		panic("ENV is not set, please set ENV to dev, test or prod")
	}

	path := workdir + "/configs"
	viper.AddConfigPath(path)
	viper.SetConfigName(os.Getenv("ENV"))
	log.Printf("work dir config.init() called")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Full file path %v", path)
		log.Printf("Failed to read %s configuration", os.Getenv("ENV"))
	}
	log.Printf("Configuration: ENV = %s, WORK_DIR = %s", os.Getenv("ENV"),
		workdir,
	)
}

func bindEnv(input map[string]string, strict bool) {
	for k, v := range input {
		if strict {
			if _, ok := os.LookupEnv(v); !ok {
				panic(fmt.Errorf("required env var not set %s", v))
			}
		}
		err := viper.BindEnv(k, v)
		if err != nil {
			panic(errors.WithMessage(err, fmt.Sprintf("Failed to bind environment variable %s", v)))
		}
	}
}

// Pwd returns the current work dir.
func Pwd() string {
	return workdir
}

// PostCodeIOEndPoint returns the postcode io host
func PostCodeIOEndPoint() string {
	return viper.GetString("postcode.host")
}

// createVarMap picks out {{var}} from in and returns a map of var -> viper(var)
func createVarMap(in string) (map[string]string, error) {
	matches := pattern.FindAllStringSubmatch(in, -1)
	seen := make(map[string]bool)
	result := make(map[string]string)
	for _, captures := range matches {
		if len(captures) >= 2 && !seen[captures[1]] {
			if viper.IsSet(captures[1]) {
				result[captures[0]] = viper.GetString(captures[1])
				seen[captures[1]] = true
			} else {
				return nil, fmt.Errorf("Replacement config key %s is undefined", captures[0])
			}
		}
	}
	return result, nil
}
