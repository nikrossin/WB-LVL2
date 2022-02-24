package sortfile

import (
	"reflect"
	"testing"
)

func TestSortFile(t *testing.T) {

	tableTest := []struct {
		input  []string
		output []string
		flags  *Flags
	}{
		{
			[]string{"33G january 123", "   6n       September 8.7", "8n July     87", "kek noMonth     ", "kek noMonth     "},
			[]string{"33G january 123", "8n July     87", "   6n       September 8.7", "kek noMonth     "},
			&Flags{K: 1, R: true, U: true, B: true, H: true},
		},
		{
			[]string{"33G january 123", "   6n       September 8.7", "8n July     87", "kek noMonth     ", "kek noMonth     "},
			[]string{"   6n       September 8.7", "8n July     87", "33G january 123", "kek noMonth     "},
			&Flags{K: 2, R: true, U: true, B: true, M: true},
		},
		{
			[]string{"33G january 123", "   6n       September 8.7", "8n July     87", "kek noMonth     ", "kek noMonth     "},
			[]string{"kek noMonth     ", "   6n       September 8.7", "8n July     87", "33G january 123"},
			&Flags{K: 3, R: false, U: true, B: true, N: true},
		},
		{
			[]string{"33G january 123", "   6n       september 8.7", "8n july     87", "kek noMonth     ", "kek noMonth     "},
			[]string{"   6n       september 8.7", "kek noMonth     ", "8n july     87", "33G january 123"},
			&Flags{K: 2, R: true, U: true, B: true},
		},
		{
			[]string{"33G january 123", "   6n       september 8.7", "8n july     87", "kek noMonth     ", "kek noMonth     "},
			[]string{"   6n       september 8.7", "kek noMonth     ", "8n july     87", "33G january 123"},
			&Flags{K: 2, R: true, U: true, B: true},
		},
	}

	for _, test := range tableTest {
		//fmt.Println("a")
		file := NewTextFile()
		file.Flags = test.flags
		file.dataStrings = test.input
		SortFile(file)
		if !reflect.DeepEqual(file.dataStrings, test.output) {
			t.Errorf("Error with sorting: %v %v", file.dataStrings, test.output)
		}
	}

	tableTestC := []struct {
		input []string
		flags *Flags
	}{

		{
			[]string{"aaa", "bbb", "ccc"},
			&Flags{C: true},
		},
		{
			[]string{"aaa 5", "bbb 4", "ccc 2"},
			&Flags{C: true, K: 2, N: true, R: true},
		},
		{
			[]string{"  aaa  January", "bbb  MaY", "ccc    October"},
			&Flags{C: true, K: 2, M: true, B: true},
		},
	}

	for _, test := range tableTestC {
		//fmt.Println("a")
		file := NewTextFile()
		file.Flags = test.flags
		file.dataStrings = test.input
		SortFile(file)
		if !file.isSorted {
			t.Errorf("Error with sorting: %v", file.dataStrings)
		}
	}
}
