package fileShare

import (
	"fmt"
	"os"
)

func MakeFiles(parts int, threshold int, prefix string, device string, shares []string) error {
	var partsGuidance, deviceGuidance, thresholdGuidance string

	partsGuidance = fmt.Sprintf("This file contains a shared secret, one of %d parts comprising a BIP39 seed.", parts)

	if device == "" {
		deviceGuidance = fmt.Sprintf("Information to identity the source device was not provided. It is likely a small USB device or a something that looks like a calculator")
	} else {
		deviceGuidance = fmt.Sprintf("The seed originated in device '%s'", device)
	}

	thresholdGuidance = fmt.Sprintf("To recover the secret, %d other parts of the share must be combined (using this tool)", threshold)

	for i := 0; i < parts; i++ {
		template := `---
information:
  aboutThisFile: ` + partsGuidance + `
  sourceDevice: ` + deviceGuidance + `
  recoveryAdvice: ` + thresholdGuidance + `
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
