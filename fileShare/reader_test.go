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
	},	{
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
