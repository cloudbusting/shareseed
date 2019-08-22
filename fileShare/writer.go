package fileShare

import (
	"fmt"
	"os"
)

func MakeFiles(parts int, threshold int, prefix string, cryptocurrency string, shares []string) error {
	template := partsGuidance(parts) + "\n" + cryptocurrencyGuidance(cryptocurrency) + "\n" + thresholdGuidance(threshold) + `
# The tool used to create the shared secrets and to recombine them can be found at https://github.com/cloudbusting/shareseed
# A BIP39 seed is used to derive public/private key pairs, in this case to secure cryptocurrency (e.g. Bitcoin)
# If you have sufficient shares to recombine and recover the secret, you should initialise a hardware wallet, using the seed to recover the addresses and stored funds
part=%d
secret=%s`

	for i := 0; i < parts; i++ {
		if err := writePart(template, prefix, i+1, parts, shares[i]); err != nil {
			return err
		}
	}

	return nil
}

func partsGuidance(parts int) string {
	return fmt.Sprintf("# This file contains a shared secret, one of %d parts comprising a BIP39 seed.", parts)
}

func cryptocurrencyGuidance(cryptocurrency string) string {
	var cryptocurrencyGuidance string
	if cryptocurrency == "" {
		cryptocurrencyGuidance = fmt.Sprintf("# Information to identity the applicable cryptocurrency was not provided. It may be Bitcoin, or potentially many cryptocurrencies.")
	} else {
		cryptocurrencyGuidance = fmt.Sprintf("# The seed applied to the cryptocurrency '%s'", cryptocurrency)
	}
	return cryptocurrencyGuidance
}

func thresholdGuidance(threshold int) string {
	var partsMaybePlural string
	if threshold <= 2 {
		partsMaybePlural = "part"
	} else {
		partsMaybePlural = "parts"
	}
	return fmt.Sprintf("# To recover the secret, at least %d other %s of the share set (%d in total) must be combined.", threshold-1, partsMaybePlural, threshold)
}

func writePart(template string, prefix string, part int, parts int, share string) error {
	if file, err := os.Create(fmt.Sprintf("%s-%d-of-%d.share.txt", prefix, part, parts)); err != nil {
		return err
	} else {
		defer file.Close()
		if _, err := fmt.Fprintf(file, template, part, share); err != nil {
			return err
		}
		return file.Sync()
	}
}
