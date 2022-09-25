package slc

// RemoveItem returns a slice after removing the item
func RemoveItem(slc []string, item string) []string {
	i := 0
	for _, v := range slc {
		if v != item {
			slc[i] = v
			i++
		}
	}
	return slc
}
