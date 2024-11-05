

func main() {
	// ##########################
	// #       MAP BASICS       #
	// ##########################

	mp := make(map[int]string)

	// Insert values
	mp[2] = "Hello"
	mp[3] = "Hello"
	mp[4] = "Hello"

	// Check if key exists
	if val, found := mp[2]; found {
		fmt.Println("Found value:", val)
	}

	// Remove key from map
	delete(mp, 2)
	fmt.Println("Map after deletion:", mp)

	// Update value in map
	mp[4] = "Four"

	// Get value for a key
	if val, found := mp[3]; found {
		fmt.Println("Value at key 3:", val)
	}

	// Get all keys from map
	fmt.Println("Keys in map:")
	for key := range mp {
		fmt.Println(key)
	}

	// Get all values from map
	fmt.Println("Values in map:")
	for _, val := range mp {
		fmt.Println(val)
	}

	// Length of map
	fmt.Println("Length of map:", len(mp))

	// Clear a map, remove all keys
	for key := range mp {
		delete(mp, key)
	}

	// ##############################
	// # FIND KEY(S) BY VALUE IN MAP #
	// ##############################

	mp[2] = "Hello"
	mp[3] = "World"
	mp[4] = "Hello"

	// Find key by value
	for key, value := range mp {
		if value == "Hello" {
			fmt.Println("Key for 'Hello':", key)
		}
	}

	// Find all keys by value
	var keysForHello []int
	for key, value := range mp {
		if value == "Hello" {
			keysForHello = append(keysForHello, key)
		}
	}
	fmt.Println("Keys for 'Hello':", keysForHello)

	// ##############################
	// #   REVERSE MAP KEYS/VALUES  #
	// ##############################

	mp1 := make(map[string]int)
	for key, value := range mp {
		mp1[value] = key
	}
	fmt.Println("Reversed map:", mp1)

	// ##########################
	// #   MAP COMPARISON       #
	// ##########################

	var m1, m2 map[string]int
	m1 = map[string]int{"hello": 1}
	m2 = map[string]int{"hello": 1}
	fmt.Println("Maps equal:", reflect.DeepEqual(m1, m2))

	// ##########################################
	// # GROUP ELEMENTS BY FREQUENCY IN A MAP   #
	// ##########################################

	slc1 := []int{23, 32, 34, 45, 32, 34, 54, 54, 42, 24, 42}
	frequencyMap := make(map[int]int)
	for _, num := range slc1 {
		frequencyMap[num]++
	}
	fmt.Println("Element frequencies:", frequencyMap)

	// ############################
	// # ELEMENT FREQUENCY COUNT  #
	// ############################

	freq := make(map[int][]int)
	for key, value := range frequencyMap {
		freq[value] = append(freq[value], key)
	}
	fmt.Println("Grouped by frequency:", freq)

	// ##########################
	// #       MAP MERGE        #
	// ##########################

	s1 := map[string]int{"get": 1, "set": 2, "go": 2}
	s2 := map[string]int{"get": 1, "set": 4, "go": 2}
	mergedMap := map[string]int{}

	for key, value := range s1 {
		mergedMap[key] = value
	}

	for key, value := range s2 {
		if _, found := mergedMap[key]; found {
			mergedMap[key] += value
		} else {
			mergedMap[key] = value
		}
	}
	fmt.Println("Merged map:", mergedMap)

	// ###################################
	// #  MAP WITH SLICE AND NESTED MAP  #
	// ###################################

	// Nested maps
	nestedMap := map[int]map[int]int{}
	nestedMap[1] = map[int]int{1: 2, 2: 2}
	nestedMap[2] = map[int]int{1: 2, 10: 2}
	fmt.Println("Nested map:", nestedMap)

	// Map with Slice values
	mapWithSlice := map[int][]int{}
	mapWithSlice[1] = append(mapWithSlice[1], 2)
	mapWithSlice[2] = append(mapWithSlice[2], 2)
	fmt.Println("Map with slices:", mapWithSlice)

	// ###########################
	// #      MAP OF STRUCTS     #
	// ###########################

	type Person struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"lte=100"`
	}

	personMap := map[int]Person{
		1: {Name: "John", Age: 30},
		2: {Name: "Doe", Age: 25},
	}
	fmt.Println("Map of structs:", personMap)

	// ##########################
	// # MOST FREQUENT ELEMENT  #
	// ##########################

	slc2 := []int{23, 32, 34, 45, 32, 32, 32, 54, 42, 24, 42}
	frequencyCount := make(map[int]int)
	for _, num := range slc2 {
		frequencyCount[num]++
	}
	maxElem, maxCount := 0, 0
	for key, val := range frequencyCount {
		if val > maxCount {
			maxCount = val
			maxElem = key
		}
	}
	fmt.Println("Most frequent element:", maxElem, "with count:", maxCount)

	// ##########################
	// # SORT ELEMENTS BY KEYS  #
	// ##########################

	type kv struct {
		Key string
		Val int
	}
	sortSlc := []kv{}
	for key, val := range mergedMap {
		sortSlc = append(sortSlc, kv{Key: key, Val: val})
	}

	sort.Slice(sortSlc, func(i, j int) bool {
		return sortSlc[i].Val < sortSlc[j].Val
	})
	fmt.Println("Sorted elements by values:", sortSlc)
}

