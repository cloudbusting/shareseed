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

package convert

import (
	"encoding/hex"
	"github.com/tyler-smith/go-bip39"
)


func toMnemonic(entropy string) (string, error) {
	bytes, err := hex.DecodeString(entropy)
	if err != nil {
		return "", err
	} else {
		mnemonic, err := bip39.NewMnemonic(bytes)
		if err != nil {
			return "", err
		} else {
			return mnemonic, nil
		}
	}
}

func toEntropy(mnemonic string) ([]byte, error) {
	inverted, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	} else {
		return inverted, nil
	}
}