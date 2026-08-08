package alx

import "sort"

// Sort sorts a slice of Language structs based on their Quality values in descending order.
// It uses a stable sort to preserve the original order of elements with the same Quality,
// which matches the order they appeared in the header.
func Sort(langs []Language) {
	sort.SliceStable(langs, func(i, j int) bool {
		return langs[i].Quality > langs[j].Quality
	})
}
