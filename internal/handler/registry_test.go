package handler

import (
	"fmt"
	"testing"

	"github.com/kryptn/http-client/language"
)

type testBlockResult []byte

func (tb testBlockResult) Value() []byte {
	return []byte(tb)
}

func (tb testBlockResult) HasValue() bool {
	return len(tb) > 0
}

func (tb testBlockResult) Error() error {
	return nil
}

func testBlockHandlerFunc(block language.Block) Result {
	return nil
}

func testFunctionHandlerFunc(prev Result, fn language.FunctionInvocation) Result {
	return nil
}

var registryNameChecks = []struct {
	names       []string
	expectError bool
}{
	{names: []string{"duplicationName"}, expectError: false},
	{names: []string{"duplicationName", "duplicationName"}, expectError: true},
	{names: []string{"caseInsensitiveDuplicates"}, expectError: false},
	{names: []string{"caseInsensitiveDuplicates", "Caseinsensitiveduplicates"}, expectError: true},
}

func runBlockDupeTest(fnNames []string, willError bool) func(*testing.T) {

	return func(t *testing.T) {
		var err error
		for _, name := range fnNames {
			err = registerBlockHandler(name, testBlockHandlerFunc)
		}

		if willError && err == nil {
			t.Errorf("case failed, expected error and got none")
		} else if !willError && err != nil {
			t.Errorf("case failed, did not expect error ")
		}
	}
}

func Test_register_block_handler(t *testing.T) {
	for i, testCase := range registryNameChecks {
		name := fmt.Sprintf("block name duplication test %d", i)
		t.Run(name, runBlockDupeTest(testCase.names, testCase.expectError))

	}
}

func runFunctionDupeTest(fnNames []string, willError bool) func(*testing.T) {

	return func(t *testing.T) {
		var err error
		for _, name := range fnNames {
			err = registerFunctionHandler(name, testFunctionHandlerFunc)
		}

		if willError && err == nil {
			t.Errorf("case failed, expected error and got none")
		} else if !willError && err != nil {
			t.Errorf("case failed, did not expect error ")
		}
	}
}

func Test_register_function_handler(t *testing.T) {
	for i, testCase := range registryNameChecks {
		name := fmt.Sprintf("function name duplication test %d", i)
		t.Run(name, runFunctionDupeTest(testCase.names, testCase.expectError))

	}
}

var getRegistryCheck = []struct {
	registerNames []string
	lookupName    string
	expectError   bool
}{
	{registerNames: []string{}, lookupName: "unregisteredName", expectError: true},
	{registerNames: []string{"registeredName"}, lookupName: "registeredName", expectError: false},
	{registerNames: []string{"caseInsName"}, lookupName: "caseinsname", expectError: false},
}

func runGetBlockRegistry(registryNames []string, lookupName string, willError bool) func(*testing.T) {
	for _, name := range registryNames {
		_ = registerBlockHandler(name, testBlockHandlerFunc)
	}

	return func(t *testing.T) {
		_, err := GetBlockHandler(lookupName)
		if willError && err == nil {
			t.Errorf("case failed, expected error and got none")

		} else if !willError && err != nil {
			t.Errorf("case failed, did not expect error ")
		}
	}
}

func Test_GetBlockHandler(t *testing.T) {
	for i, testCase := range getRegistryCheck {
		name := fmt.Sprintf("function name duplication test %d", i)
		t.Run(name, runGetBlockRegistry(testCase.registerNames, testCase.lookupName, testCase.expectError))
	}
}
