package sortfile

import (
	"reflect"
	"sort"
	"testing"
)

func TestDeleteDuplicate(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"bbb bbb", "bbb aaa", "bbb aaa", "ccc ccc"},
			[]string{"bbb bbb", "bbb aaa", "ccc ccc"},
		},
	}
	for _, test := range tableTest {
		unique := deleteDuplicateLines(test.input)
		sort.Strings(unique)
		sort.Strings(test.output)
		if !reflect.DeepEqual(unique, test.output) {
			t.Errorf("Error with deleteDuplicate: %v %v", unique, test.output)
		}
	}
}
