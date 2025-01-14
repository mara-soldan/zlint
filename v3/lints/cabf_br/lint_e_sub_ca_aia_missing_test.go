package cabf_br

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func Test_SubCaAiaMissing(t *testing.T) {

	var tests = []struct {
		name      string
		inputPath string

		expected lint.LintStatus
	}{
		{
			name:      "pass - cert valid",
			inputPath: "subCAAIAValid.pem",

			expected: lint.Pass,
		},
		{
			name:      "not effective - test case for CABF_BR 1.7.1 version of lint",
			inputPath: "subCAAIAMissingPostCABFBR171.pem",

			expected: lint.NE,
		},
		{
			name:      "error - intermediate cert missing AIA",
			inputPath: "subCAAIAMissing.pem",

			expected: lint.Error,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			out := test.TestLint("e_sub_ca_aia_missing", testCase.inputPath)

			if out.Status != testCase.expected {
				t.Errorf("%s: expected %s, got %s", testCase.inputPath, testCase.expected, out.Status)
			}
		})
	}
}
