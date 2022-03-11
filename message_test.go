package quasar_fire_core

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	type TestCase struct {
		Messages      [][]string
		ExpectedValue []string
		ExpectedError error
	}

	testCases := []TestCase{
		{
			Messages: [][]string{
				{"", "este", "es", "un", "mensaje", "", "un", "amigo"},
				{"", "esta", "", "", "mensajero", "de", "", ""},

			},
			ExpectedValue: nil,
			ExpectedError: errors.New("A matching pair of words was not found. A matching pair is necessary to calculate the offset"),
		},
		{
			Messages: [][]string{
				{"", "este", "es", "un", "mensaje"},
				{"este", "", "un", "mensaje"},
			},
			ExpectedValue: []string{"este", "es", "un", "mensaje"},
			ExpectedError: nil,
		},
		{
			Messages: [][]string{
				{"este", "es", "un", "mensaje"},
				{"", "", "es", "", "mensaje"},
			},
			ExpectedValue: []string{"este", "es", "un", "mensaje"},
			ExpectedError: nil,
		},
		{
			Messages: [][]string{
				{"", "", "", "es", "", "mensaje", "", "un", "amigo"},
				{"este", "", "un", "mensaje", "", "", "amigo"},
			},
			ExpectedValue: []string{"este", "es", "un", "mensaje", "", "un", "amigo"},
			ExpectedError: nil,
		},
		{
			Messages: [][]string{
				{"este", "es", "un", "mensaje", "", "un", "amigo"},
				{"", "este", "", "un", "mensaje", "de", "", "amigo"},

			},
			ExpectedValue: []string{"este", "es", "un", "mensaje", "de", "un", "amigo"},
			ExpectedError: nil,
		},
		{
			Messages: [][]string{
				{"", "este", "es", "un", "mensaje", "", "un", "amigo"},
				{"", "este", "", "", "mensaje", "de", "", ""},

			},
			ExpectedValue: []string{"este", "es", "un", "mensaje", "de", "un", "amigo"},
			ExpectedError: nil,
		},
	}

	for _, tc := range testCases {
		fmt.Printf("tc.Messages[0]: %+v\n", tc.Messages[0])
		fmt.Printf("tc.Messages[1]: %+v\n", tc.Messages[1])
		message, err := MergeMessages(tc.Messages[0], tc.Messages[1])

		assert.Equal(t, tc.ExpectedError, err)
		assert.Equal(t, tc.ExpectedValue, message)
	}
}
