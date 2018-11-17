package removeduplicate

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	tcs := []struct {
		input  []string
		output []string
	}{
		{[]string{"a", "a", "c"}, []string{"a", "c"}},
		{[]string{"a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, tc := range tcs {
		r := RemoveDuplicate(tc.input)
		if !ssEqual(r, tc.output) {
			t.Errorf("Expect %v but got %v", tc.output, r)
		}
	}
}

func BenchmarkRemoveDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicate([]string{"a", "a", "c"})
	}
}

func ssEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
