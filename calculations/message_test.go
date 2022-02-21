package calculations

import (
	"fmt"
	"testing"
)

func TestGetMessage(t *testing.T) {
	type TestCase struct{
		Messages [][]string
		Expected string
	}

	// First: Message brute force solution
	// Then: Location calculation
	// Then: API controller
	// Then GCP Deployment
	// Then Brute force optimisation
	testCases := []TestCase{
		{
			Messages: [][]string{
				{"", "este", "es", "un", "mensaje"},
				{"este", "", "un", "mensaje"},
				{"", "", "es", "", "mensaje"},
			},
			Expected: "este es un mensaje",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("tc: %+v", tc)
	}
}
