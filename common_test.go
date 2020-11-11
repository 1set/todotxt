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

func checkTaskListOrder(t *testing.T, gotList TaskList, expStrList []string) {
	if len(gotList) < len(expStrList) {
		t.Errorf("Got less elements (%d) than expected (%d)", len(gotList), len(expStrList))
	}

	for i, expected := range expStrList {
		if got := gotList[i].Task(); got != expected {
			t.Errorf("Expected Task[%d] after Sort() to be [%s], but got [%s]", i, expected, got)
		}
	}
}
