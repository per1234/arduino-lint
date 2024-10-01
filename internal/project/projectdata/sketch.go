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

package projectdata

import (
	"strings"

	"github.com/arduino/arduino-cli/arduino/sketch"
	"github.com/arduino/arduino-lint/internal/project"
)

// InitializeForSketch gathers the check data for the specified sketch project.
func InitializeForSketch(project project.Type) {
	loadedSketch, sketchLoadError = sketch.New(ProjectPath())
	// TODO: Better error detection
	if sketchLoadError != nil && strings.Contains(sketchLoadError.Error(), "importing sketch metadata") {
		metadataLoadError = sketchLoadError
	}
}

var sketchLoadError error

// SketchLoadError returns the error output from Arduino CLI loading the sketch.
func SketchLoadError() error {
	return sketchLoadError
}

var loadedSketch *sketch.Sketch

// Sketch returns the sketch object generated by Arduino CLI.
func Sketch() *sketch.Sketch {
	return loadedSketch
}

var metadataLoadError error

// MetadataLoadError returns the error produced during load of the sketch metadata.
func MetadataLoadError() error {
	return metadataLoadError
}
