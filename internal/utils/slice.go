package utils

// Slice util to deal with slice

// SliceHas checks given interface if in the interface slice.
func SliceHas(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceDiffUint64 returns diff slice of slice1 - slice2.
// for uint64
func SliceDiffUint64(slice1, slice2 []uint64) (diffslice []uint64) {
	for _, v := range slice1 {
		inSlice2 := false
		for _, vv := range slice2 {
			if vv == v {
				inSlice2 = true
			}
		}

		if inSlice2 == false {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceDiff returns diff slice of slice1 - slice2.
func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		// if v is not in the slice2
		if !SliceHas(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceIntersect returns the intersection of slice1 and slice2.
func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if SliceHas(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

// SliceMerge merges interface slices to one slice.
func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}
