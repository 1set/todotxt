package todotxt

import (
	"fmt"
	"strings"
	"testing"
)

func compareSlices(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func compareMaps(map1 map[string]string, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	compare := func(map1 map[string]string, map2 map[string]string) bool {
		for key, value := range map1 {
			if value2, found := map2[key]; !found {
				return false
			} else if value != value2 {
				return false
			}
		}
		return true
	}

	return compare(map1, map2) && compare(map2, map1)
}

func isSameTaskSegmentList(s1, s2 []*TaskSegment) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		a, b := s1[i], s2[i]
		if a.Type != b.Type {
			return false
		}
		if a.Display != b.Display {
			return false
		}
		if !compareSlices(a.Originals, b.Originals) {
			return false
		}
	}
	return true
}

func strTaskSegmentList(l []*TaskSegment) string {
	var parts []string
	for _, s := range l {
		parts = append(parts, fmt.Sprintf("%v", *s))
	}
	return strings.Join(parts, ", ")
}

func Test_lessStrings(t *testing.T) {
	tests := []struct {
		a    []string
		b    []string
		want bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"b"}, true},
		{[]string{"a", "b", "c"}, []string{""}, false},
		{[]string{"a", "b", "c"}, []string{}, false},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b"}, []string{"a", "a"}, false},
		{[]string{"a", "b"}, []string{"a", "c"}, true},
		{[]string{"a", "b"}, []string{"b"}, true},
		{[]string{"a", "a"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "a"}, []string{"a", "b"}, true},
		{[]string{"a", "a"}, []string{"a", "a"}, false},
		{[]string{"a", "a"}, []string{"a", "c"}, true},
		{[]string{"a", "a"}, []string{"b"}, true},
		{[]string{"a", "a"}, []string{"b", "a"}, true},
		{[]string{"a", "a"}, []string{"c"}, true},
		{[]string{"a", "c"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "c"}, []string{"a", "c"}, false},
		{[]string{"a", "c"}, []string{"b"}, true},
		{[]string{"a", "c"}, []string{"c"}, true},
		{[]string{"b"}, []string{"a", "b", "c"}, false},
		{[]string{"b"}, []string{"a", "c"}, false},
		{[]string{"b"}, []string{"b"}, false},
		{[]string{"b"}, []string{"b", "a"}, true},
		{[]string{"b"}, []string{"c"}, true},
		{[]string{"b"}, []string{""}, false},
		{[]string{"b"}, []string{}, false},
		{[]string{"b", "a"}, []string{"a", "b", "c"}, false},
		{[]string{"b", "a"}, []string{"b"}, false},
		{[]string{"b", "a"}, []string{"b", "a"}, false},
		{[]string{"b", "a"}, []string{"c"}, true},
		{[]string{"c"}, []string{"a", "b", "c"}, false},
		{[]string{"c"}, []string{"b"}, false},
		{[]string{""}, []string{"a", "b", "c"}, true},
		{[]string{""}, []string{"c"}, true},
		{[]string{""}, []string{""}, false},
		{[]string{""}, []string{}, false},
		{[]string{}, []string{"a", "b", "c"}, true},
		{[]string{}, []string{"c"}, true},
		{[]string{}, []string{""}, true},
		{[]string{}, []string{}, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case#%d", i+1), func(t *testing.T) {
			if got := lessStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("lessStrings() %v < %v got = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
