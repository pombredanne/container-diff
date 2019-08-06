/*
Copyright 2018 Google, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import "fmt"

// Bump this on release
var version = "v0.15.0"

// When built using `make` this is overridden via -ldflags
var gitVersion = "(unknown)"

// returns just the vX.Y.Z version suitable for `make release`
func GetShortVersion() string {
	return version
}

func GetVersion() string {
	return fmt.Sprintf("%s built from git %s", version, gitVersion)
}
