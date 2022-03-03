package test

import (
	cut "lvl2/develop/dev06/internal"
	"reflect"
	"testing"
)

func TestParseFlagF(t *testing.T) {
	c := &cut.Config{}

	Tests := []struct {
		flag    string
		Fvalues cut.FValues
	}{
		{
			"2",
			cut.FValues{1: false},
		},
		{
			"-3",
			cut.FValues{0: false, 1: false, 2: false},
		},
		{
			"2-",
			cut.FValues{1: true},
		},
		{
			"3-5",
			cut.FValues{2: false, 3: false, 4: false},
		},
		{
			"2,4-6,7-8",
			cut.FValues{1: false, 3: false, 4: false, 5: false, 6: false, 7: false},
		},
		{
			"-2,3-4,6-",
			cut.FValues{0: false, 1: false, 2: false, 3: false, 5: true},
		},
		{
			"-3,4,5-",
			cut.FValues{0: false, 1: false, 2: false, 3: false, 4: true},
		},
		{
			"-2,4,5,6-8,100-",
			cut.FValues{0: false, 1: false, 3: false, 4: false, 5: false, 6: false, 7: false, 99: true},
		},
		{
			"-5,7-9,12,14-15,6",
			cut.FValues{0: false, 1: false, 2: false, 3: false, 4: false, 6: false, 7: false, 8: false, 11: false,
				13: false, 14: false, 5: false},
		},
	}

	for num, val := range Tests {
		c.F = val.flag
		c.FValues = make(cut.FValues)
		if err := c.ParseFlagF(); err != nil {
			t.Errorf("Error from test with correct flag F")
		}
		if !reflect.DeepEqual(c.FValues, val.Fvalues) {
			t.Errorf("Error with ranges of flag F %v %v %v", num, c.FValues, val.Fvalues)
		}
	}

	ErrorTests := []string{"-", "--", "--6", "6--", "-6-", "1-2-3", "-1-2", "2-3-", "34a", "4-5d5", "5a-", "-5a"}
	for _, val := range ErrorTests {
		c.F = val
		c.FValues = make(cut.FValues)
		if err := c.ParseFlagF(); err == nil {
			t.Errorf("Incorrect work errors: %v", val)
		}
	}
}
