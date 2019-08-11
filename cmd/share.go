/*
Copyright © 2019 Andy Hitchman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/cloudbusting/shareseed/convert"
	"github.com/cloudbusting/shareseed/fileShare"
	"os"

	"github.com/spf13/cobra"
)

var (
	ErrInvalidPartsSize          = errors.New("parts must be between 2 and 255")
	ErrInvalidThresholdSize      = errors.New("threshold must be between 1 and 255")
	ErrThresholdGreaterThanParts = errors.New("threshold cannot be greater than parts")
	ErrNoOutputOptions           = errors.New("no output would be generated. quiet and fileparts must be set appropriately")
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Take a BIP39 mnemonic phrase and produce Shamir Secret Share mnemonic phrases",
	Long: `Share
-----
Inputs:
Provide a BIP39 mnemonic phrase (from stdin, or as a parameter---not recommended, see below).
Specify the number of parts to share the secret in, and the number of parts required to successfully recombine (the threshold).

Outputs:
A set of numbered parts comprising a 2-digit hex number and seed words. The hex number and seed words represent the shareable part. 
Optionally creates share parts and/or QR code files. You would write these to USB drives, for example (with this executable for good measure).

Note that specifying the mnemonic seed phrase using the --mnemonic flag may leave the phrase in the shell history, which could have serious security risks. Consider HISTCONTROL settings, or redirect input.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic, _ := cmd.Flags().GetString("mnemonic")
		parts, _ := cmd.Flags().GetInt("parts")
		threshold, _ := cmd.Flags().GetInt("threshold")
		prefix, _ := cmd.Flags().GetString("prefix")
		device, _ := cmd.Flags().GetString("device")
		quiet, _ := cmd.Flags().GetBool("quiet")
		fileparts, _ := cmd.Flags().GetBool("fileparts")

		if err := validParams(parts, threshold, quiet, fileparts); err != nil {
			return err
		}

		return execute(mnemonic, parts, threshold, prefix, device, fileparts)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().IntP("parts", "p", 5, "The number of parts to produce")
	shareCmd.Flags().IntP("threshold", "t", 3, "The number of parts required in combination to reproduce the BIP39 mnemonic")
	shareCmd.Flags().String("prefix", "", "A prefix for the shared part number. Helps identify part sets, e.g. 'BTC'")
	shareCmd.Flags().String("device", "", "A meaningful identifier for the source device, e.g. 'ColdCard 1'")
	shareCmd.Flags().BoolP("fileparts", "f", false, "Write parts into separate files per part, named for the prefix and part number")
	shareCmd.Flags().BoolP("quiet", "q", false, "Do not write shares to standard output")
	shareCmd.Flags().StringP("mnemonic", "m", "", "The BIP39 mnemonic phrase. Omit to read from pipe (stdin), or redirect ('<')")
}

func validParams(parts int, threshold int, quiet bool, fileparts bool) error {
	if parts < 2 || parts > 255 {
		return ErrInvalidPartsSize
	}
	if threshold < 1 || threshold > 255 {
		return ErrInvalidThresholdSize
	}
	if threshold > parts {
		return ErrThresholdGreaterThanParts
	}
	if quiet && !fileparts {
		return ErrNoOutputOptions
	}
	return nil
}

func execute(mnemonic string, parts int, threshold int, prefix string, device string, fileparts bool) error {
	if mnemonic == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			mnemonic += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	if shares, err := convert.Share(mnemonic, parts, threshold); err != nil {
		return err
	} else {
		fmt.Printf("Sharing seed in %d parts, requiring %d shares to recover secret\n\n", parts, threshold)

		for part, share := range shares {
			fmt.Printf("%s-%d-of-%d: %s\n", prefix, part+1, parts, share)
		}

		if fileparts {
			fmt.Println("\nFiles for each part have been created:")
			return fileShare.MakeFiles(parts, threshold, prefix, device, shares)

		} else {
			fmt.Println("\nRecord and store each share separately")
		}
	}

	return nil
}