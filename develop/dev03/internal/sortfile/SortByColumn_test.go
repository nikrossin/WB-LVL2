package sortfile

import (
	"reflect"
	"testing"
)

func TestSortByColumn(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"bbb bbb", "bbb aaa", "aaa aaa", "ccc ccc"},
			[]string{"bbb aaa", "aaa aaa", "bbb bbb", "ccc ccc"},
			&Flags{K: 2},
		},
		{
			[]string{" bbb  bbb", " bbb  bbb", "bbb aaa", "ccc ccc"},
			[]string{"bbb aaa", " bbb  bbb", " bbb  bbb", "ccc ccc"},
			&Flags{K: 2, B: true},
		},
		{
			[]string{" bbb  bbb", " bbb  bbb", "bbb aaa", "ccc ccc"},
			[]string{"ccc ccc", " bbb  bbb", " bbb  bbb", "bbb aaa"},
			&Flags{K: 2, B: true, R: true},
		},
	}
	for _, test := range tableTest {
		if !reflect.DeepEqual(sortByColumn(test.input, test.flags), test.output) {
			t.Errorf("Error with testByColumn: %v %v", test.output, sortByColumn(test.input, test.flags))
		}
	}
}
