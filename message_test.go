package quasar_fire_core

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
		{
			Messages: [][]string{
				{"", "", "", "es", "", "mensaje", "", "un", "amigo"},
				{"este", "", "un", "mensaje", "", "", "amigo"},
				{"", "este", "", "un", "mensaje", "de", "", "amigo"},
			},
			Expected: "este es un mensaje de un amigo",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("tc: %+v", tc)
	}
}

func TestMergeMessages(t *testing.T) {
	type TestCase struct{
		Messages [][]string
		Expected []string
	}

	testCases := []TestCase{
		{
			Messages: [][]string{
				{"", "este", "es", "un", "mensaje"},
				{"este", "", "un", "mensaje"},
			},
			Expected: []string{"este", "es", "un", "mensaje"},
		},
		{
			Messages: [][]string{
				{"este", "es", "un", "mensaje"},
				{"", "", "es", "", "mensaje"},
			},
			Expected: []string{"este", "es", "un", "mensaje"},
		},
		{
			Messages: [][]string{
				{"", "", "", "es", "", "mensaje", "", "un", "amigo"},
				{"este", "", "un", "mensaje", "", "", "amigo"},
			},
			Expected: []string{"este", "es", "un", "mensaje", "", "un", "amigo"},
		},
		{
			Messages: [][]string{
				{"este", "es", "un", "mensaje", "", "un", "amigo"},
				{"", "este", "", "un", "mensaje", "de", "", "amigo"},

			},
			Expected: []string{"este", "es", "un", "mensaje", "de", "un", "amigo"},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("tc: %+v", tc)
	}
}
