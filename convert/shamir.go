package convert

import (
	"encoding/hex"
	"github.com/hashicorp/vault/shamir"
)

func Share(mnemonic string, parts int, threshold int) ([]string, error) {
	if rawEntropy, err := toEntropy(mnemonic); err != nil {
		return nil, err
	} else {
		if shares, err := shamir.Split(rawEntropy, parts, threshold); err != nil {
			return nil, err
		} else {
			seedShares := make([]string, parts)
			for part, share := range shares {
				hShare := hex.EncodeToString(share)
				if seed, err := toMnemonic(hShare[2:]); err != nil {
					return nil, err
				} else {
					seedShares[part] = hShare[:2] + " " + seed
				}
			}
			return seedShares, nil
		}
	}
}

func Combine(parts []string) (string, error) {
	byteParts := make([][]byte, len(parts))
	for i, part := range parts {
		if postfixBytes, err := toEntropy(part[3:]); err != nil {
			return "", err
		} else {
			if prefixBytes, err := hex.DecodeString(part[:2]); err != nil {
				return "", err
			} else {
				byteParts[i] = append(prefixBytes, postfixBytes...)
			}
		}
	}

	if secret, err := shamir.Combine(byteParts); err != nil {
		return "", err
	} else {
		return toMnemonic(hex.EncodeToString(secret))
	}
}
