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
	"testing"
)

type bip39example struct {
	entropy  string
	mnemonic string
}

var validConvertExamples = []bip39example{
	{
		entropy:  "00000000000000000000000000000000",
		mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
	},
	{
		entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
		mnemonic: "legal winner thank year wave sausage worth useful legal winner thank yellow",
	},
	{
		entropy:  "80808080808080808080808080808080",
		mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage above",
	},
	{
		entropy:  "ffffffffffffffffffffffffffffffff",
		mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong",
	},
	{
		entropy:  "000000000000000000000000000000000000000000000000",
		mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent",
	},
	{
		entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
		mnemonic: "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will",
	},
	{
		entropy:  "808080808080808080808080808080808080808080808080",
		mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always",
	},
	{
		entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffff",
		mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when",
	},
	{
		entropy:  "0000000000000000000000000000000000000000000000000000000000000000",
		mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art",
	},
	{
		entropy:  "7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f",
		mnemonic: "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title",
	},
	{
		entropy:  "8080808080808080808080808080808080808080808080808080808080808080",
		mnemonic: "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless",
	},
	{
		entropy:  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		mnemonic: "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote",
	},
	{
		entropy:  "9e885d952ad362caeb4efe34a8e91bd2",
		mnemonic: "ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic",
	},
	{
		entropy:  "6610b25967cdcca9d59875f5cb50b0ea75433311869e930b",
		mnemonic: "gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog",
	},
	{
		entropy:  "68a79eaca2324873eacc50cb9c6eca8cc68ea5d936f98787c60c7ebc74e6ce7c",
		mnemonic: "hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length",
	},
	{
		entropy:  "c0ba5a8e914111210f2bd131f3d5e08d",
		mnemonic: "scheme spot photo card baby mountain device kick cradle pact join borrow",
	},
	{
		entropy:  "6d9be1ee6ebd27a258115aad99b7317b9c8d28b6d76431c3",
		mnemonic: "horn tenant knee talent sponsor spell gate clip pulse soap slush warm silver nephew swap uncle crack brave",
	},
	{
		entropy:  "9f6a2878b2520799a44ef18bc7df394e7061a224d2c33cd015b157d746869863",
		mnemonic: "panda eyebrow bullet gorilla call smoke muffin taste mesh discover soft ostrich alcohol speed nation flash devote level hobby quick inner drive ghost inside",
	},
	{
		entropy:  "23db8160a31d3e0dca3688ed941adbf3",
		mnemonic: "cat swing flag economy stadium alone churn speed unique patch report train",
	},
	{
		entropy:  "8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0",
		mnemonic: "light rule cinnamon wrap drastic word pride squirrel upgrade then income fatal apart sustain crack supply proud access",
	},
	{
		entropy:  "066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad",
		mnemonic: "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform",
	},
	{
		entropy:  "f30f8c1da665478f49b001d94c5fc452",
		mnemonic: "vessel ladder alter error federal sibling chat ability sun glass valve picture",
	},
	{
		entropy:  "c10ec20dc3cd9f652c7fac2f1230f7a3c828389a14392f05",
		mnemonic: "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
	},
	{
		entropy:  "f585c11aec520db57dd353c69554b21a89b20fb0650966fa0a9d6f74fd989d8f",
		mnemonic: "void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold",
	},
}

func TestToEntropy(t *testing.T) {
	for _, example := range validConvertExamples {
		entropy, err := toEntropy(example.mnemonic)
		if err != nil {
			t.Errorf("ToEntropy should not error")
		}
		got := hex.EncodeToString(entropy)
		if got != example.entropy {
			t.Errorf("Expected entropy %s to equual example %s", got, example.entropy)
		}
	}
}

func TestToMnemonic(t *testing.T) {
	for _, example := range validConvertExamples {
		got, err := toMnemonic(example.entropy)
		if err != nil {
			t.Errorf("ToMnemonic should not error")
		}
		if got != example.mnemonic {
			t.Errorf("Expected got %s to equual example %s", got, example.mnemonic)
		}
	}
}
