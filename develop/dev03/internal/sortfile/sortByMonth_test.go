package sortfile

import (
	"reflect"
	"testing"
)

func TestSortByMonth(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"february bbb", "September aaa", "JuNe aaa", "may ccc"},
			[]string{"february bbb", "may ccc", "JuNe aaa", "September aaa"},
			&Flags{K: 1},
		},
		{
			[]string{" bbb  february", "aaa September", "aaa JuNe", "ccc may"},
			[]string{" bbb  february", "ccc may", "aaa JuNe", "aaa September"},
			&Flags{K: 2, B: true},
		},
		{
			[]string{" bbb  february", "aaa September", "aaa JuNe", "ccc may"},
			[]string{"aaa September", "aaa JuNe", "ccc may", " bbb  february"},
			&Flags{K: 2, B: true, R: true},
		},
	}
	for _, test := range tableTest {
		//fmt.Println("a")
		m := SortByMonth(test.input, test.flags)
		if !reflect.DeepEqual(m, test.output) {
			t.Errorf("Error with testByMonth: %v %v", m, test.output)
		}
	}
}
