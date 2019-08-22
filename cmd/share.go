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
Optionally creates share parts files. You might then write these files to USB drives (with this executable for good measure).

Warning:
Specifying the mnemonic seed phrase using the --mnemonic flag may leave the phrase in the shell history, which could have serious security risks. Consider HISTCONTROL settings, or redirect input.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mnemonic, _ := cmd.Flags().GetString("mnemonic")
		parts, _ := cmd.Flags().GetInt("parts")
		threshold, _ := cmd.Flags().GetInt("threshold")
		cryptocurrency, _ := cmd.Flags().GetString("cryptocurrency")
		quiet, _ := cmd.Flags().GetBool("quiet")
		fileparts, _ := cmd.Flags().GetString("fileparts")

		if err := validParams(parts, threshold, quiet, fileparts); err != nil {
			return err
		}

		return executeShare(mnemonic, parts, threshold, fileparts, cryptocurrency, quiet)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().IntP("parts", "p", 5, "The number of parts to produce")
	shareCmd.Flags().IntP("threshold", "t", 3, "The number of parts required in combination to reproduce the BIP39 mnemonic")
	shareCmd.Flags().StringP("cryptocurrency", "c", "", "The cryptocurrency (or a list) that the seed is applicable to (e.g 'Bitcoin'). For output file information only.")
	shareCmd.Flags().StringP("fileparts", "f", "", "Write parts into separate files per part, named for the prefix (e.g. 'BTC') and part number")
	shareCmd.Flags().BoolP("quiet", "q", false, "Do not write shares to standard output")
	shareCmd.Flags().StringP("mnemonic", "m", "", "The BIP39 mnemonic phrase. Omit to read from stdin, i.e. pipe ('|') or redirect ('<') from a file")
}

func validParams(parts int, threshold int, quiet bool, fileparts string) error {
	if parts < 2 || parts > 255 {
		return ErrInvalidPartsSize
	}
	if threshold < 1 || threshold > 255 {
		return ErrInvalidThresholdSize
	}
	if threshold > parts {
		return ErrThresholdGreaterThanParts
	}
	if quiet && fileparts == "" {
		return ErrNoOutputOptions
	}
	return nil
}

func executeShare(mnemonic string, parts int, threshold int, fileparts string, cryptocurrency string, quiet bool) error {
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

		if !quiet {
			for part, share := range shares {
				fmt.Printf("%s-%d-of-%d: %s\n", fileparts, part+1, parts, share)
			}
			fmt.Println()
		}
		if fileparts != "" {
			fmt.Printf("Files for each part have been created using prefix '%s'\n", fileparts)
			return fileShare.MakeFiles(parts, threshold, fileparts, cryptocurrency, shares)
		} else {
			fmt.Println("Record and store each share separately")
		}
	}

	return nil
}
