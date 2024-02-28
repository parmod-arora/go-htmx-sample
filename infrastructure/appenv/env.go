package appenv

import "os"

func GetWithDefault(name string, defaultValue string) string {
	result, ok := os.LookupEnv(name)
	if ok {
		return result
	}
	return defaultValue
}
