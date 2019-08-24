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
	"math/rand"
	"testing"
)

func pickParts(shares []string, threshold int) []string {
	picks := make([]string, threshold)
	for i := 0; i < threshold; i++ {
		pick := rand.Intn(threshold - i)
		picks[i] = shares[pick]
		shares = append(shares[:pick], shares[pick+1:]...)
	}
	return picks
}

func TestMakeShares_ProducesRequestedParts(t *testing.T) {
	mnemonic := "void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold"

	for expectedParts := 2; expectedParts <= 255; expectedParts++ {
		var randomThreshold int
		if expectedParts == 2 {
			randomThreshold = 2
		} else {
			randomThreshold = rand.Intn(expectedParts-2) + 2
		}
		got, err := Share(mnemonic, expectedParts, randomThreshold)
		if err != nil {
			t.Errorf("makeShares should not error")
		}
		if len(got) != expectedParts {
			t.Errorf("Expected %d produced got to equal %d expected parts", len(got), expectedParts)
		}
	}
}

type shamirExample struct {
	mnemonic        string
	parts           int
	threshold       int
	sharesAvailable int
}

var recoverableExamples = []shamirExample{
	{
		mnemonic:        "gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog",
		parts:           10,
		threshold:       4,
		sharesAvailable: 4,
	},
	{
		mnemonic:        "hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length",
		parts:           6,
		threshold:       3,
		sharesAvailable: 4,
	},
	{
		mnemonic:        "scheme spot photo card baby mountain device kick cradle pact join borrow",
		parts:           25,
		threshold:       20,
		sharesAvailable: 20,
	},
	{
		mnemonic:        "horn tenant knee talent sponsor spell gate clip pulse soap slush warm silver nephew swap uncle crack brave",
		parts:           15,
		threshold:       15,
		sharesAvailable: 15,
	},
	{
		mnemonic:        "panda eyebrow bullet gorilla call smoke muffin taste mesh discover soft ostrich alcohol speed nation flash devote level hobby quick inner drive ghost inside",
		parts:           2,
		threshold:       2,
		sharesAvailable: 2,
	},
	{
		mnemonic:        "cat swing flag economy stadium alone churn speed unique patch report train",
		parts:           3,
		threshold:       2,
		sharesAvailable: 2,
	},
	{
		mnemonic:        "light rule cinnamon wrap drastic word pride squirrel upgrade then income fatal apart sustain crack supply proud access",
		parts:           3,
		threshold:       2,
		sharesAvailable: 3,
	},
	{
		mnemonic:        "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform",
		parts:           5,
		threshold:       2,
		sharesAvailable: 2,
	},
	{
		mnemonic:        "vessel ladder alter error federal sibling chat ability sun glass valve picture",
		parts:           9,
		threshold:       2,
		sharesAvailable: 9,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 31,
	},
	{
		mnemonic:        "void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold",
		parts:           11,
		threshold:       7,
		sharesAvailable: 10,
	},
}

func TestCombine_WithSufficientShares(t *testing.T) {
	for _, example := range recoverableExamples {
		if shares, err := Share(example.mnemonic, example.parts, example.threshold); err != nil {
			t.Error(err)
		} else {
			parts := pickParts(shares, example.sharesAvailable)

			if got, err := Combine(parts); err != nil {
				t.Error(err)
			} else {
				if got != example.mnemonic {
					t.Errorf("Expected recovered seed '%s' to equal example '%s'", got, example.mnemonic)
				}
			}
		}
	}
}

var insufficientSharesExamples = []shamirExample{
	{
		mnemonic:        "void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold",
		parts:           6,
		threshold:       4,
		sharesAvailable: 3,
	},
	{
		mnemonic:        "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform",
		parts:           5,
		threshold:       2,
		sharesAvailable: 0,
	},
	{
		mnemonic:        "vessel ladder alter error federal sibling chat ability sun glass valve picture",
		parts:           9,
		threshold:       2,
		sharesAvailable: 1,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 29,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 28,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 27,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 20,
	},
	{
		mnemonic:        "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		parts:           100,
		threshold:       30,
		sharesAvailable: 2,
	},

}

func TestCombine_WithInsufficientShares(t *testing.T) {
	for _, example := range insufficientSharesExamples {
		if shares, err := Share(example.mnemonic, example.parts, example.threshold); err != nil {
			t.Error(err)
		} else {
			parts := pickParts(shares, example.sharesAvailable)

			got, err := Combine(parts)
			if got == example.mnemonic {
				t.Errorf("Expected to not recover correct seed %s with insufficient shares available", example.mnemonic)
			} else {
				if got == "" && err == nil {
					t.Errorf("Expected failure as insufficient shares with threshold of %d and %d available", example.threshold, example.sharesAvailable)
				}
			}
		}
	}
}
