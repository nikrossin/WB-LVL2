package main

import (
	"fmt"
	"testing"
)

func TestHandleJson(t *testing.T) {
	tableTest := []struct {
		testString     string
		expectedString string
		err            error
	}{
		{
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"abcd",
			"abcd",
			nil,
		},
		{
			"45",
			"",
			fmt.Errorf("Incorrect string"),
		},
		{
			"qwe\\4\\5",
			"qwe45",
			nil,
		},
		{
			"qwe\\\\5",
			"qwe\\\\\\\\\\",
			nil,
		},
		{
			"qwe\\45",
			"qwe44444",
			nil,
		},
		{
			"",
			"",
			nil,
		},
		{
			"b4b5\\",
			"",
			fmt.Errorf("Incorrect string"),
		},
	}

	for _, testItem := range tableTest {
		s, err := unpack(testItem.testString)
		if s != testItem.expectedString {
			t.Errorf("Fail test with testString: %v", testItem.testString)
		}
		if !((err != nil && testItem.err != nil) || (err == nil && testItem.err == nil)) {
			t.Errorf("Fail test with testString: %v err: %v", testItem.testString, testItem.err)
		}

	}
}
