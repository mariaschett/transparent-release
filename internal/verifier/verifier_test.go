// Copyright 2022 The Project Oak Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package verify

import (
	"os"
	"testing"

	"github.com/project-oak/transparent-release/internal/testutil"
)

const schemaExamplePath = "schema/amber-slsa-buildtype/v1/example.json"

func TestReproducibleProvenanceVerifier(t *testing.T) {
	// The path to provenance is specified relative to the root of the repo, so we need to go one level up.
	// Get the current directory before that to restore the path at the end of the test.
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("couldn't get current directory: %v", err)
	}
	defer testutil.Chdir(t, currentDir)
	testutil.Chdir(t, "../../")
	verifier := ReproducibleProvenanceVerifier{}

	if err := verifier.Verify(schemaExamplePath); err != nil {
		t.Fatalf("couldn't verify the provenance file: %v", err)
	}
}

func TestAmberProvenanceMetadataVerifier(t *testing.T) {
	// The path to provenance is specified relative to the root of the repo, so we need to go one level up.
	// Get the current directory before that to restore the path at the end of the test.
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("couldn't get current directory: %v", err)
	}
	defer testutil.Chdir(t, currentDir)
	testutil.Chdir(t, "../../")
	verifier := AmberProvenanceMetadataVerifier{}

	if err := verifier.Verify(schemaExamplePath); err != nil {
		t.Fatalf("couldn't verify the provenance file: %v", err)
	}
}
