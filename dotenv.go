package dotenv

import (
	"os"
	"strings"
)

// Get the value for a key
func GetEnv(key string) string {
	return os.Getenv(key)
}

// Load the .env, and set the key and value if they don't exist
func LoadFile(path string) error {
	envMap, err := readFile(path)
	curEnv := make(map[string]bool)
	env := os.Environ()
	for _, k := range env {
		split := strings.Split(k, "=")
		curEnv[split[0]] = true
	}
	for key, value := range envMap {
		if !curEnv[key] {
			err = os.Setenv(key, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Load the .env file, set the key and his value, if a key already exists it changes the value of the key
func OverloadFile(path string) error {
	envMap, err := readFile(path)
	if err != nil {
		return err
	}
	for key, value := range envMap {
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// Define the value for the key
func SetEnv(key string, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filename string) (map[string]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	s := strings.SplitN(string(file), "\r\n", -1)
	envMap := make(map[string]string)
	for _, i := range s {
		split := strings.Split(i, "=")
		envMap[split[0]] = split[1]
	}
	return envMap, nil
}
