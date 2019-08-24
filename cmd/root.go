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
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shareseed",
	Short: "Securely share a BIP39 mnemonic seed in multiple parts, and recombine from a subset of parts.",
	Long: `
Use Shamir Secret Sharing to protect BIP39 mnemonic seeds by sharing in parts, and to recover from those parts.

Warnings
--------
Test this utility with an unimportant seed to share and recombine parts to understand its function. 
(Valid test mnemonic seeds can be created at https://iancoleman.io/bip39/).

To reduce the possible risk of a malware infection capturing the inputs or outputs of this utility, it is strongly 
recommended that you run this tool from a secure OS, such as a live boot of tails.

Typing or storing your seed phrase in plain text on your normal desktop OS is extremely dangerous.
Additionally, storing the output parts of this utility together is no safer than simply storing your seed in plain text.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
