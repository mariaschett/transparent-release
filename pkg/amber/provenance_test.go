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

package amber

import (
	"fmt"
	"os"
	"testing"

	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
	"github.com/project-oak/transparent-release/internal/testutil"
)

const (
	provenanceExamplePath    = "schema/amber-slsa-buildtype/v1/example.json"
	wantSHA1HexDigitLength   = 40
	wantSHA256HexDigitLength = 64
)

func TestExampleProvenance(t *testing.T) {
	// The path to provenance is specified relative to the root of the repo, so we need to go one level up.
	// Get the current directory before that to restore the path at the end of the test.
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("couldn't get current directory: %v", err)
	}
	defer testutil.Chdir(t, currentDir)
	testutil.Chdir(t, "../../")

	// Parses the provenance and validates it against the schema.
	provenance, err := ParseProvenanceFile(provenanceExamplePath)
	if err != nil {
		t.Fatalf("Failed to parse example provenance: %v", err)
	}

	predicate := provenance.Predicate.(slsa.ProvenancePredicate)
	buildConfig := predicate.BuildConfig.(BuildConfig)

	// Check that the provenance parses correctly
	testutil.AssertEq(t, "repoURL", predicate.Materials[1].URI, "https://github.com/project-oak/oak")
	testutil.AssertEq(t, "commitHash length", len(predicate.Materials[1].Digest["sha1"]), wantSHA1HexDigitLength)
	testutil.AssertEq(t, "builderImageID length", len(predicate.Materials[0].Digest["sha256"]), wantSHA256HexDigitLength)
	testutil.AssertEq(t, "builderImageURI", predicate.Materials[0].URI, fmt.Sprintf("gcr.io/oak-ci/oak@sha256:%s", predicate.Materials[0].Digest["sha256"]))
	testutil.AssertEq(t, "subjectName", provenance.Subject[0].Name, "oak_functions_loader")
	testutil.AssertNonEmpty(t, "command[0]", buildConfig.Command[0])
	testutil.AssertNonEmpty(t, "command[1]", buildConfig.Command[1])
	testutil.AssertNonEmpty(t, "builderId", predicate.Builder.ID)
}
