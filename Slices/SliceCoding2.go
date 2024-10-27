
// ##########################
// # REMOVE DUPLICATES FROM SLICE #
// ##########################
func RemoveDuplicatesFromSlice() {
	s1 := []int{23, 34, 17, 23, 8, 23, 23, 34, 32, 34}
	for i := 0; i < len(s1); {
		sn := s1[i]
		index := slices.Index(s1, sn)
		if index != i {
			s1 = slices.Delete(s1, i, i+1)
		} else {
			i++
		}
	}
	fmt.Println("After removing duplicates:", s1)
}

// ##########################
// # ROTATE SLICE N POSITIONS #
// ##########################
func RotateSliceNPosition(n int) []string {
	slc := []string{"test0", "test1", "test2", "test3", "test4"}
	if n > len(slc)-1 {
		return nil
	}
	slc1 := slc[:n]
	slices.Reverse(slc1)
	slc1 = append(slc1, slc[n:]...)
	return slc1
}

// ##########################
// # MERGE TWO SORTED SLICES #
// ##########################
func MergeTwoSlices() {
	slc1 := []int{1, 10, 20, 30, 40}
	slc2 := []int{5, 15, 25, 35, 45, 55, 65}
	finalSlice := make([]int, 0, len(slc1)+len(slc2))
	i, j := 0, 0

	for i < len(slc1) && j < len(slc2) {
		if slc1[i] < slc2[j] {
			finalSlice = append(finalSlice, slc1[i])
			i++
		} else if slc1[i] > slc2[j] {
			finalSlice = append(finalSlice, slc2[j])
			j++
		} else {
			finalSlice = append(finalSlice, slc1[i], slc2[j])
			i++
			j++
		}
	}

	for i < len(slc1) {
		finalSlice = append(finalSlice, slc1[i])
		i++
	}

	for j < len(slc2) {
		finalSlice = append(finalSlice, slc2[j])
		j++
	}

	fmt.Println("Merged slice:", finalSlice)
}

// ##########################
// # SPLIT SLICE INTO CHUNKS #
// ##########################
func SplitSlice(size int) {
	slc := make([]string, 0)
	for i := 0; i < 50; i++ {
		slc = append(slc, fmt.Sprintf("Test %d", i))
	}

	i, j := 0, size
	for j < len(slc) {
		fmt.Printf("Chunk: %q\n\n", slc[i:j])
		i += size
		j += size
	}
}

// ##########################
// # REMOVE ALL OCCURRENCES #
// ##########################
func RemoveAllOccurrences() {
	slc := []string{"test0", "test1", "test2", "test3", "test4"}
	strToRemove := "Test1"
	slc = slices.DeleteFunc(slc, func(str string) bool {
		return strings.EqualFold(str, strToRemove)
	})
	fmt.Println("After removing all occurrences:", slc)
}

// ##########################
// # COUNT OCCURRENCES OF ELEMENTS #
// ##########################
func CountOccurrences() {
	slc := []int{12, 23, 12, 23, 43, 32, 65, 54, 45, 43, 65, 90, 12, 90, 54}
	mp := make(map[int]int)
	for _, item := range slc {
		mp[item]++
	}
	fmt.Println("Occurrences:", mp)
}

// ##########################
// # SUM OF SLICE ELEMENTS #
// ##########################
func SumOfSliceElements() {
	slc := []int{12, 23, 12, 23, 43, 32, 65, 54, 45, 43, 65, 90, 12, 90, 54}
	sum := 0
	for _, item := range slc {
		sum += item
	}
	fmt.Println("Sum of elements:", sum)
}

// ##########################
// # SLIDING WINDOW MAXIMUM #
// ##########################
func SlidingWindowMaxInSlice() {
	window := 3
	slc := []int{12, 23, 12, 23, 43, 32, 65, 54, 45, 43, 65, 90, 12, 90, 54}
	maxCount := 0
	for i := 0; i <= len(slc)-window; i++ {
		localCount := 0
		for j := i; j < i+window; j++ {
			localCount += slc[j]
		}
		if localCount > maxCount {
			maxCount = localCount
		}
	}
	fmt.Println("Sliding window max count:", maxCount)
}

// ##########################
// # SLIDING WINDOW MINIMUM #
// ##########################
func SlidingWindowMinInSlice() {
	window := 3
	slc := []int{12, 23, 12, 23, 43, 32, 65, 54, 45, 43, 65, 90, 12, 90, 54}
	minCount := math.MaxInt
	for i := 0; i <= len(slc)-window; i++ {
		localCount := 0
		for j := i; j < i+window; j++ {
			localCount += slc[j]
		}
		if localCount < minCount {
			minCount = localCount
		}
	}
	fmt.Println("Sliding window min count:", minCount)
}

// ##########################
// # ROTATE SLICE BY K ELEMENTS #
// ##########################
func RotateKSlice(k int) {
	slc := []int{12, 23, -12, -23, 43, -32, 65, -54, 45, -43, -65, -90, 12, 90}
	if k > len(slc) {
		return
	}
	i, j := 0, k-1
	for j > i {
		slc[i], slc[j] = slc[j], slc[i]
		i++
		j--
	}
	fmt.Println("Rotated slice:", slc)
}

// ##########################
// # CHECK ELEMENT CONTAINED #
// ##########################
func ContainElement(element int) {
	s1 := []int{23, 34, 17, 23, 8}
	fmt.Println("Contains element:", slices.Contains(s1, element))
}
