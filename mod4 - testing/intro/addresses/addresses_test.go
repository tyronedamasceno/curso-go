package addresses_test

import (
	"intro/addresses"
	"testing"
)

type testScenario struct {
	inputAddress   string
	expectedOutput string
}

func Test_AddressType(t *testing.T) {
	t.Parallel()

	scenarios := []testScenario{
		{"Street 123", "Street"},
		{"STREET 567", "Street"},
		{"Avenue Jackson", "Avenue"},
		{"Road Black", "Road"},
		{"My House", "Invalid"},
	}

	for _, scenario := range scenarios {
		result := addresses.AddressType(scenario.inputAddress)
		if result != scenario.expectedOutput {
			t.Errorf("Result %s is different from expected %s", result, scenario.expectedOutput)
		}
	}
}

func Test_2(t *testing.T) {
	t.Parallel()
}
