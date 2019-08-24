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
	"bytes"
	"testing"
)

type secretExample struct {
	expectedSecret string
	expectedErr    error
	given          string
}

var examples = []secretExample{
	{
		expectedSecret: "abcdef",
		expectedErr:    nil,
		given: `
# comment 1
# comment 2
secret=abcdef`,
	},
	{
		expectedSecret: "abcdef",
		expectedErr:    nil,
		given: `
# comment 1
# comment 2
secret=    abcdef      
part=3`,
	},
	{
		expectedSecret: "cheese toast sardines eggs salmon sausage chips beans ketchup",
		expectedErr:    nil,
		given: `
# comment 1
# comment 2
secret=cheese toast sardines eggs salmon sausage chips beans ketchup`,
	},
	{
		expectedSecret: "cheese toast sardines eggs salmon sausage chips beans ketchup",
		expectedErr:    nil,
		given: `
# comment 1
# comment 2
secret=   cheese      toast			 sardines eggs salmon            sausage chips beans ketchup           `,
	}, {
		expectedSecret: "",
		expectedErr:    ErrSecretNotFoundInFile,
		given: `
# comment 1
# comment 2`,
	},
}

func TestGetSecretFromReader_ShouldReturnSecret(t *testing.T) {
	for _, example := range examples {
		actual, err := getSecretFromReader(bytes.NewBufferString(example.given))

		if actual != example.expectedSecret {
			t.Errorf("expectedSecret %s\n\n    but got\n\n%s", example.expectedSecret, actual)
		}
		if err != example.expectedErr {
			t.Errorf("expectedSecret err %s\n\n	but got\n\n%s", example.expectedErr, err)
		}
	}
}
