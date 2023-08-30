package types

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDurationUnmarshal(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		expectedResult *Duration
		expectedErr    error
	}

	testCases := []testCase{
		{
			name:           "valid duration I",
			input:          "60s",
			expectedResult: &Duration{Duration: time.Minute},
		},
		{
			name:           "valid duration II",
			input:          "1m0s",
			expectedResult: &Duration{Duration: time.Minute},
		},
		{
			name:        "int value",
			input:       "60",
			expectedErr: fmt.Errorf("time: missing unit in duration \"60\""),
		},
		{
			name:        "no duration value",
			input:       "abc",
			expectedErr: fmt.Errorf("time: invalid duration \"abc\""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var d Duration
			input, err := json.Marshal(testCase.input)
			require.NoError(t, err)
			err = json.Unmarshal(input, &d)

			if testCase.expectedResult != nil {
				require.Equal(t, (*testCase.expectedResult).Nanoseconds(), d.Nanoseconds())
			}

			if err != nil {
				require.Equal(t, testCase.expectedErr.Error(), err.Error())
			}
		})
	}
}

func TestDurationMarshal(t *testing.T) {
	type testCase struct {
		name           string
		input          *Duration
		expectedResult string
	}

	testCases := []testCase{
		{
			name:           "valid duration",
			input:          &Duration{Duration: time.Minute},
			expectedResult: `"1m0s"`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			byteDuration, err := json.Marshal(testCase.input)
			require.NoError(t, err)
			require.Equal(t, string(byteDuration), testCase.expectedResult)
		})
	}
}
