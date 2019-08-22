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