package slc

import "fmt"

// RemoveItemWithSelf returns a slice after removing the specifed item
func RemoveItemWithSelf(slc []string, item string) []string {
	for i := 0; i < len(slc); i++ {
		if slc[i] == item {
			slc = append(slc[:i], slc[i+1:]...)
			i--
		}
	}
	return slc
}

// RemoveItemByCopy returns a slice after removing the specifed item
func RemoveItemByCopy(slc []string, item string) []string {
	results := make([]string, 0, len(slc))
	for _, v := range slc {
		if v != item {
			results = append(results, v)
		}
	}
	return results
}

// RemoveItemBySub returns a slice after removing the specified item
func RemoveItemBySub(slc []string, item string) []string {
	i := 0
	for _, v := range slc {
		if v != item {
			slc[i] = v
			i++
		}
	}
	fmt.Println(slc)
	return slc[:i]
}

// RemoveItemWithShare returns a slice after removing the specified
func RemoveItemWithShare(slc []string, item string) []string {
	results := slc[:0]
	for _, v := range slc {
		if v != item {
			results = append(results, v)
		}
	}
	return results
}
