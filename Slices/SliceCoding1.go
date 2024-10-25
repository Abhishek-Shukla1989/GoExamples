// ##########################
// # APPEND TO A SLICE      #
// ##########################
func AppendSlice() {
    oldSlice := []string{"one", "two", "three"}
    // oldSliceWithMake := make([]int, 5) // created with capacity and initialized with 0
    oldSlice = append(oldSlice, "four", "five", "six")
    fmt.Println(oldSlice)
}

// ##########################
// # INSERT ELEMENT IN SLICE #
// ##########################
func InsertElement(m string, position int) {
    // First way
    s1 := []string{"one", "two", "three"}
    if position < len(s1) {
        s1 = slices.Insert(s1, position, m)
    }
    fmt.Println(s1)

    // Second way
    s1 = append(s1, "")
    copy(s1[position+1:], s1[position:])
    s1[position] = m
    fmt.Println(s1)
}

// ##########################
// # CONCATENATE TWO SLICES  #
// ##########################
func ConcatSlice(s1, s2 []string) {
    // First method
    s3 := slices.Concat(s1, s2)
    fmt.Println(s3)

    // Append slice to the first one
    s1 = append(s1, s2...)
    fmt.Println(s1)
}

// ##########################
// # DELETE ELEMENT AT INDEX #
// ##########################
func DeleteElementAtIndex(position int) {
    s1 := []string{"one", "two", "three"}

    // First method using slice package
    if len(s1) >= position-1 {
        s1 = slices.Delete(s1, position, position+1)
    }
    fmt.Println(s1)

    // Second method
    s1 = append(s1[:position], s1[position+1:]...)
    fmt.Println(s1)
}

// ##########################
// # COPY SLICE              #
// ##########################
func CopySlice() {
    s1 := []string{"one", "two", "three"}
    var s2 = make([]string, 5)
    copy(s2, s1)
    fmt.Println(s2)
}

// ##########################
// # REVERSE SLICE           #
// ##########################
func ReverseSlice() {
    s1 := []int{23, 34, 17, 23, 8}
    slices.Reverse(s1)
    fmt.Println(s1)
}

// ##########################
// # CHECK ELEMENT CONTAINED #
// ##########################
func ContainElement(element int) {
    s1 := []int{23, 34, 17, 23, 8}
    fmt.Println(slices.Contains(s1, element))
}

// ##########################
// # MAX ELEMENT IN SLICE    #
// ##########################
func MaxElement() {
    s1 := []int{23, 34, 17, 23, 8}
    s2 := []string{"zen", "ewp", "ers", "abc"}

    fmt.Println(slices.Max(s1)) // Get max element
    fmt.Println(slices.Max(s2)) // Get max element

    fmt.Println(slices.Min(s1)) // Get min element
    fmt.Println(slices.Min(s2)) // Get min element
}

// ##########################
// # SORT SLICE              #
// ##########################
func SortSlice() {
    s1 := []int{23, 34, 17, 23, 8}
    s2 := []string{"zen", "ewp", "ers", "abc"}

    // Method one - Sort in descending order using custom function
    slices.SortFunc(s2, func(a, b string) int {
        return strings.Compare(strings.ToLower(b), strings.ToLower(a))
    })

    // Ascending order sort
    sort.Ints(s1)   // Integer sort, ascending
    sort.Strings(s2) // String sort, ascending

    // Method two - Sort using custom comparison function
    sort.Slice(s2, func(i, j int) bool {
        return s2[i] > s2[j]
    })
    sort.Slice(s1, func(i, j int) bool {
        return s1[i] > s1[j]
    })
    
    fmt.Println(s1)
    fmt.Println(s2)
}
