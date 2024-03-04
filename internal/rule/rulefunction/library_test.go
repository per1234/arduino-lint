// This file is part of Arduino Lint.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License, either
// version 3 of the License, or (at your option) any later version.
// This license covers the main part of Arduino Lint.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package rulefunction

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/arduino/arduino-lint/internal/project"
	"github.com/arduino/arduino-lint/internal/project/projectdata"
	"github.com/arduino/arduino-lint/internal/project/projecttype"
	"github.com/arduino/arduino-lint/internal/rule/ruleresult"
	"github.com/arduino/go-paths-helper"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var librariesTestDataPath *paths.Path

func init() {
	workingDirectory, _ := os.Getwd()
	librariesTestDataPath = paths.New(workingDirectory, "testdata", "libraries")
}

type libraryRuleFunctionTestTable struct {
	testName            string
	libraryFolderName   string
	expectedRuleResult  ruleresult.Type
	expectedOutputQuery string
}

func checkLibraryRuleFunction(ruleFunction Type, testTables []libraryRuleFunctionTestTable, t *testing.T) {
	for _, testTable := range testTables {
		expectedOutputRegexp := regexp.MustCompile(testTable.expectedOutputQuery)

		testProject := project.Type{
			Path:             librariesTestDataPath.Join(testTable.libraryFolderName),
			ProjectType:      projecttype.Library,
			SuperprojectType: projecttype.Library,
		}

		projectdata.Initialize(testProject)

		result, output := ruleFunction()
		assert.Equal(t, testTable.expectedRuleResult, result, testTable.testName)
		assert.True(t, expectedOutputRegexp.MatchString(output), fmt.Sprintf("%s (output: %s, assertion regex: %s)", testTable.testName, output, testTable.expectedOutputQuery))
	}
}

func TestLibraryInvalid(t *testing.T) {
	testTables := []libraryRuleFunctionTestTable{
		{"Invalid library.properties", "InvalidLibraryProperties", ruleresult.Fail, ""},
		{"Invalid flat layout", "FlatWithoutHeader", ruleresult.Fail, ""},
		{"Invalid recursive layout", "RecursiveWithoutLibraryProperties", ruleresult.Fail, ""},
		{"Valid library", "Recursive", ruleresult.Pass, ""},
	}

	checkLibraryRuleFunction(LibraryInvalid, testTables, t)
}
