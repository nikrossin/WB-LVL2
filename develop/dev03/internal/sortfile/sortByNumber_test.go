package sortfile

import (
	"reflect"
	"testing"
)

func TestSortByNumber(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"123 aaa", "-13 ads", "-22.5 ddd"},
			[]string{"-22.5 ddd", "-13 ads", "123 aaa"},
			&Flags{K: 1},
		},
		{
			[]string{" aaa  123 ", "ddd -22.5", "    ads -13"},
			[]string{" aaa  123 ", "    ads -13", "ddd -22.5"},
			&Flags{K: 2, B: true, R: true},
		},
		{
			[]string{"    ads -13", " aaa  123 ", "ddd -22.5", "    ads -13"},
			[]string{" aaa  123 ", "    ads -13", "    ads -13", "ddd -22.5"},
			&Flags{K: 2, B: true, R: true},
		},
	}

	for _, test := range tableTest {
		//fmt.Println("a")
		m := sortByNumbers(test.input, test.flags)
		if !reflect.DeepEqual(m, test.output) {
			t.Errorf("Error with testByMonth: %v %v", m, test.output)
		}
	}
}
