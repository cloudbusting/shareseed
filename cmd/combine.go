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

package cmd

import (
	"fmt"
	"github.com/cloudbusting/shareseed/convert"
	"github.com/cloudbusting/shareseed/fileShare"
	"github.com/spf13/cobra"
)

// combineCmd represents the combine command
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Combine a quorum of shared parts to reproduce the original secret BIP39 mnemonic phrase",
	Long:  `Combine
-------
Inputs:
Either a set of file shares containing shared secrets, or a set of shared secrets passed directly in parameters.

Outputs:
Prints the recovered seed to standard output. You must provide sufficient shares to recover the correct seed. 
i.e. you must provide at least shares equal to required threshold; this is stated in file shares, but you may not
have this information if you are working from other media. 

Note:
Different wallets may treat the same phrase differently, meaning that if the phrase was created with one wallet
and restored in another, that wallet may show no coins. In this case try to establish the wallet the seed was
created from and refer to https://walletsrecovery.org/ for guidance.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		secrets, _ := cmd.Flags().GetStringSlice("secret")
		filePattern, _ := cmd.Flags().GetString("filepattern")

		return executeCombine(secrets, filePattern)
	},
}

func init() {
	rootCmd.AddCommand(combineCmd)
	combineCmd.Flags().StringSliceP("secret", "s", []string{}, "Provide a set of secret shares. Quote each secret. Comma separate or provide flag multiple times")
	combineCmd.Flags().StringP("filepattern", "f", "", "The pattern matching of a set of shares. e.g. 'BTC*.share.txt'. Quote in shells that automatically expand globs")
}

func executeCombine(secrets []string, filePattern string) error {
	if filePattern != "" {
		if fileSecrets, err := fileShare.FilesToSecrets(filePattern); err != nil {
			return err
		} else {
			secrets := append(secrets, fileSecrets...)
			if seed, err := convert.Combine(secrets); err != nil {
				return err
			} else {
				fmt.Println(seed)
			}
		}
	}
	return nil
}