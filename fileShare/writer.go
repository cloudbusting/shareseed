package fileShare

import (
	"fmt"
	"os"
)

func MakeFiles(parts int, threshold int, prefix string, cryptocurrency string, shares []string) error {
	var partsGuidance, cryptocurrencyGuidance, thresholdGuidance string

	partsGuidance = fmt.Sprintf("This file contains a shared secret, one of %d parts comprising a BIP39 seed.", parts)

	if cryptocurrency == "" {
		cryptocurrencyGuidance = fmt.Sprintf("Information to identity the applicable cryptocurrency was not provided. It may be Bitcoin, or potenially many cryptocurrencies.")
	} else {
		cryptocurrencyGuidance = fmt.Sprintf("The seed applied to the cryptocurrency '%s'", cryptocurrency)
	}

	var partsMaybePlural string
	if threshold <= 2 {
		partsMaybePlural = "part"
	} else {
		partsMaybePlural = "parts"
	}
	thresholdGuidance = fmt.Sprintf("To recover the secret, at least %d other %s of the share set (%d in total) must be combined", threshold-1, partsMaybePlural, threshold)

	for i := 0; i < parts; i++ {
		template := `---
information:
  aboutThisFile: ` + partsGuidance + `
  cryptocurrency: ` + cryptocurrencyGuidance + `
  recoveryAdvice: ` + thresholdGuidance + `
  toolLocation: The tool used to create the shared secrets and to recombine them can be found at https://github.com/cloudbusting/shareseed
  whatIsASeed: A BIP39 seed is used to derive public/private key pairs, in this case to secure cryptocurreny (e.g. Bitcoin)
  whatShouldIDoWithTheRecoveredSecret: If you have sufficient shares to recombine and recover the secret, you should initialise a hardware wallet, using the seed to recover the addresses and stored funds
share: %s`
		if err := writePart(template, prefix, i+1, parts, shares[i]); err != nil {
			return err
		}
	}

	return nil
}

func writePart(template string, prefix string, part int, parts int, share string) error {
	if file, err := os.Create(fmt.Sprintf("%s-%d-of-%d.share.yaml", prefix, part, parts)); err != nil {
		return err
	} else {
		defer file.Close()
		if _, err := fmt.Fprintf(file, template, share); err != nil {
			return err
		}
		return file.Sync()
	}
}
