/*
 * Copyright © 2019 Andy Hitchman
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fileShare

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrNoMatchingFiles      = errors.New("no matching file parts")
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
		return getSecretFromReader(file)
	}
}

func getSecretFromReader(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
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
	return strings.Join(strings.Fields(strings.TrimPrefix(secret, "secret=")), " "), nil
}
