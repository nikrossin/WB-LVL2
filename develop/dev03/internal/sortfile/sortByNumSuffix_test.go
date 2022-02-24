package sortfile

import (
	"reflect"
	"testing"
)

func TestSortByNumSuffix(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"33G aaa", "8n ads", "6n ddd"},
			[]string{"6n ddd", "8n ads", "33G aaa"},
			&Flags{K: 1},
		},
		{
			[]string{" aaa 33G ", "ads        8n ", "ddd 6 "},
			[]string{" aaa 33G ", "ddd 6 ", "ads        8n "},
			&Flags{K: 2, B: true, R: true},
		},
		{
			[]string{"8.5n ddd", "33GGG aaa", "8n ads"},
			[]string{"33GGG aaa", "8n ads", "8.5n ddd"},
			&Flags{K: 1},
		},
		{
			[]string{"8.5n ddd", "lG aaa", "8n ads"},
			[]string{"lG aaa", "8n ads", "8.5n ddd"},
			&Flags{K: 1},
		},
	}

	for _, test := range tableTest {
		//fmt.Println("a")
		m := SortBySuffix(test.input, test.flags)
		if !reflect.DeepEqual(m, test.output) {
			t.Errorf("Error with testByMonth: %v %v", m, test.output)
		}
	}
}
