package sortfile

import "sort"

func basicSort(input []string, flags *Flags) []string {

	if flags.R {
		sort.Sort(sort.Reverse(sort.StringSlice(input)))
	} else {
		sort.Strings(input)
	}
	return input
}
