/*
Copyright 2019 The Knative Authors

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

// error.go helps with error handling

package main

const (
	// Minimal ratio of results to be counted as valid results for each testcase, this is an arbitrary number
	requiredRatio = 0.8
	// Don't do anything if found more than 5 tests flaky, or 1% tests flaky, whichever comes first
	countThreshold   = 5
	percentThreshold = 0.01

	org = "knative"
)
