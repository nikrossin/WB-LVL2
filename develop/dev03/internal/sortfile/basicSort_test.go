package sortfile

import (
	"reflect"
	"testing"
)

func TestBasicSort(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"bbb bbb", "bbb aaa", "aaa aaa", "ccc ccc"},
			[]string{"aaa aaa", "bbb aaa", "bbb bbb", "ccc ccc"},
			&Flags{R: false},
		},
		{
			[]string{"bbb bbb", "bbb aaa", "aaa aaa", "ccc ccc"},
			[]string{"ccc ccc", "bbb bbb", "bbb aaa", "aaa aaa"},
			&Flags{R: true},
		},
	}
	for _, test := range tableTest {
		b := basicSort(test.input, test.flags)
		if !reflect.DeepEqual(b, test.output) {
			t.Errorf("Error with testBasic: %v %v", b, test.output)
		}
	}
}
