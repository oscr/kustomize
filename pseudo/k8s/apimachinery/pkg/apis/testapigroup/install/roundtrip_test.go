/*
Copyright 2017 The Kubernetes Authors.

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

package install

import (
	"testing"

	"sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/api/apitesting/roundtrip"
	testapigroupfuzzer "sigs.k8s.io/kustomize/pseudo/k8s/apimachinery/pkg/apis/testapigroup/fuzzer"
)

func TestRoundTrip(t *testing.T) {
	roundtrip.RoundTripTestForAPIGroup(t, Install, testapigroupfuzzer.Funcs)
	roundtrip.RoundTripProtobufTestForAPIGroup(t, Install, testapigroupfuzzer.Funcs)
}