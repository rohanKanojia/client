// Copyright 2019 The Knative Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build e2e

package e2e

import (
	"fmt"
	"strings"
	"testing"
)

var (
	e env
	k kn
)

const (
	KnDefaultTestImage string = "gcr.io/knative-samples/helloworld-go"
)

func Setup(t *testing.T) func(t *testing.T) {
	e = buildEnv(t)
	k = kn{t, e.Namespace, Logger{}}
	CreateTestNamespace(t, e.Namespace)
	return Teardown
}

func Teardown(t *testing.T) {
	DeleteTestNamespace(t, e.Namespace)
}

func TestBasicWorkflow(t *testing.T) {
	teardown := Setup(t)
	defer teardown(t)

	testServiceGetEmpty(t, k)
	testServiceCreate(t, k, "hello")
	testServiceGet(t, k, "hello")
	testServiceDescribe(t, k, "hello")
	testServiceDelete(t, k, "hello")
	testServiceGetEmpty(t, k)
}

// Private test functions

func testServiceGetEmpty(t *testing.T, k kn) {
	out, err := k.RunWithOpts([]string{"service", "get"}, runOpts{NoNamespace: false})
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error executing 'kn service get' command. Error: %s", err.Error()))
	}

	if !strings.Contains(out, "No resources found.") {
		t.Fatalf("Expected output 'No resources found.' Instead found:\n%s\n", out)
	}
}

func testServiceCreate(t *testing.T, k kn, serviceName string) {
	out, err := k.RunWithOpts([]string{"service", "create",
		fmt.Sprintf("%s", serviceName),
		"--image", KnDefaultTestImage}, runOpts{NoNamespace: false})
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error executing 'kn service create' command. Error: %s", err.Error()))
	}

	if !strings.Contains(out, fmt.Sprintf("Service '%s' successfully created in namespace '%s'.", serviceName, k.namespace)) {
		t.Fatalf(fmt.Sprintf("Expected to find: Service '%s' successfully created in namespace '%s'. Instead found:\n%s\n", serviceName, k.namespace, out))
	}
}

func testServiceGet(t *testing.T, k kn, serviceName string) {
	out, err := k.RunWithOpts([]string{"service", "get", serviceName}, runOpts{NoNamespace: false})
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error executing 'kn service get %s' command. Error: %s", serviceName, err.Error()))
	}

	expectedOutput := fmt.Sprintf("%s", serviceName)
	if !strings.Contains(out, expectedOutput) {
		t.Fatalf("Expected output incorrect, expecting to include:\n%s\n Instead found:\n%s\n", expectedOutput, out)
	}
}

func testServiceDescribe(t *testing.T, k kn, serviceName string) {
	out, err := k.RunWithOpts([]string{"service", "describe", serviceName}, runOpts{NoNamespace: false})
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error executing 'kn service describe' command. Error: %s", err.Error()))
	}

	expectedOutputHeader := `apiVersion: knative.dev/v1alpha1
kind: Service
metadata:`
	if !strings.Contains(out, expectedOutputHeader) {
		t.Fatalf(fmt.Sprintf("Expected output incorrect, expecting to include:\n%s\n Instead found:\n%s\n", expectedOutputHeader, out))
	}

	expectedOutput := `generation: 1
  name: %s
  namespace: %s`
	expectedOutput = fmt.Sprintf(expectedOutput, serviceName, k.namespace)
	if !strings.Contains(out, expectedOutput) {
		t.Fatalf(fmt.Sprintf("Expected output incorrect, expecting to include:\n%s\n Instead found:\n%s\n", expectedOutput, out))
	}
}

func testServiceDelete(t *testing.T, k kn, serviceName string) {
	out, err := k.RunWithOpts([]string{"service", "delete", serviceName}, runOpts{NoNamespace: false})
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error executing 'kn service delete' command. Error: %s", err.Error()))
	}

	if !strings.Contains(out, fmt.Sprintf("Service '%s' successfully deleted in namespace '%s'.", serviceName, k.namespace)) {
		t.Fatalf(fmt.Sprintf("Expected to find: Service '%s' successfully deleted in namespace '%s'. Instead found:\n%s\n", serviceName, k.namespace, out))
	}
}
