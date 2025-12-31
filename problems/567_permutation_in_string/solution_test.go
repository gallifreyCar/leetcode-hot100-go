package permutation_in_string

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name   string
		s1     string
		s2     string
		expect bool
	}{
		{"example1", "ab", "eidbaooo", true},
		{"example2", "ab", "eidboaoo", false},
		{"single", "a", "a", true},
		{"not_found", "abc", "bbbd", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckInclusion(tt.s1, tt.s2); got != tt.expect {
				t.Errorf("CheckInclusion() = %v, want %v", got, tt.expect)
			}
		})
	}
}
