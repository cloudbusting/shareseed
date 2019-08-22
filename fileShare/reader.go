package fileShare

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrNoMatchingFiles = errors.New("no matching file parts")
	ErrSecretNotFoundInFile = errors.New("secret not found in file")
)

func FilesToSecrets(filePattern string) ([]string, error) {
	if matches, err := filepath.Glob(filePattern); err != nil {
		return nil, err
	} else {
		if len(matches) < 1 {
			return nil, ErrNoMatchingFiles
		}
		secrets := make([]string, len(matches))
		for i, file := range matches {
			if secret, err := getSecretFromFile(file); err != nil {
				return nil, err
			} else {
				secrets[i] = secret
			}
		}
		return secrets, nil
	}
}

func getSecretFromFile(name string) (string, error) {
	if file, err := os.Open(name); err != nil {
		return "", err
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		secret := ""

		for scanner.Scan() {
			if strings.HasPrefix(strings.TrimSpace(scanner.Text()), "secret=") {
				secret = scanner.Text()
				break
			}
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		if secret == "" {
			return "", ErrSecretNotFoundInFile
		}

		return strings.TrimSpace(strings.TrimPrefix(secret,"secret=")), nil
	}
}
