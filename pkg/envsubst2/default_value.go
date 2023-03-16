package envsubst2

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var defaultValues = make(map[string]string, 0)

func InitDefaultValues(flag *Flag) error {
	if flag.DefaultValuesFile == "" {
		return nil
	}

	data, err := os.ReadFile(flag.DefaultValuesFile)
	if err != nil {
		return fmt.Errorf("read default values file failed: %w", err)
	}

	err = yaml.Unmarshal(data, defaultValues)
	if err != nil {
		return fmt.Errorf("unmarshal default values file failed: %w", err)
	}

	// fmt.Println(defaultValues)

	return nil
}

func getDefaultValue(key string) (string, bool) {

	if len(defaultValues) == 0 {
		return "", false
	}

	val, ok := defaultValues[key]
	return val, ok
}
