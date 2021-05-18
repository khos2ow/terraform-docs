/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package version

import (
	"fmt"
	"runtime"
	"time"
)

// current version
const dev = "v0.13.0"

// Provisioned by ldflags
var (
	version    string
	commitHash string
	buildDate  string
)

// Load defaults for info variables
func init() {
	if version == "" {
		version = dev
	}
	if version == "v0.0.0-" { // building in a directory which is not a git repository
		version = dev
	}
	if commitHash == "" {
		commitHash = "HEAD"
	}
	if buildDate == "" {
		buildDate = time.Now().Format("2006-01-02")
	}
}

// Short return the version of the binary
func Short() string {
	return version
}

// Full return the full version of the binary including commit hash and build date
func Full() string {
	// if !strings.HasSuffix(version, commitHash) {
	// 	version += " " + commitHash
	// }
	return fmt.Sprintf(
		"%s %s/%s (%s %s)",
		version,
		runtime.GOOS,
		runtime.GOARCH,
		commitHash,
		buildDate,
	)
}
