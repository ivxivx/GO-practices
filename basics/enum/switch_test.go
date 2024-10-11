package enum

import (
	"testing"
)

func Test_Enum_Switch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		nameEnum       Name
		expectedResult string
	}{
		{
			name:           "zhang san",
			nameEnum:       ZhangSan,
			expectedResult: "San Zhang",
		},
		{
			name:           "li shi",
			nameEnum:       LiShi,
			expectedResult: "Shi Li",
		},
		{
			name:           "other",
			nameEnum:       Other,
			expectedResult: "invalid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualResult := convert(tc.nameEnum)

			if actualResult != tc.expectedResult {
				t.Errorf("expected body %v, but got %v", tc.expectedResult, actualResult)
			}
		})
	}
}
