// ##########################
// #    SUBSTRING SEARCH     #
// ##########################
func substringSearch() {
	substr := "K"
	str := strings.Contains(strings.ToLower("Abhishek"), strings.ToLower(substr))
	fmt.Println("Substring Search:", str)
}

// ##########################
// # FIND ALL OCCURRENCES OF #
// #    A SUBSTRING          #
// ##########################

func findAllOccurrences() {
	substr := "K"
	count := strings.Count(strings.ToLower("Abhishekaa"), strings.ToLower(substr))
	fmt.Println("Occurrences of Substring:", count)
}


// ##########################
// #   COUNT OCCURRENCES     #
// #  OF A CHARACTER 'a'     #
// ##########################
func countOccurrencesOfChar() {
	charCount := strings.Count(strings.ToLower("Abhishekaa"), "a")
	fmt.Println("Count of Character 'a':", charCount)
}

// ##########################
// #  LONGEST COMMON PREFIX  #
// ##########################
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // takes the first string

	for i := 1; i < len(strs); i++ { // loop over the strings
		// Trim the prefix until the current string starts with it
		for len(prefix) > 0 && !startsWith(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1] // Trim last character from prefix
			fmt.Printf("Prefix is %s\n", prefix)
		}

		// If the prefix becomes empty, there's no common prefix
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}


func startsWith(str, prefix string) bool {
	if len(prefix) > len(str) {
		return false
	}
	return str[:len(prefix)] == prefix
}

// ##########################
// # LONGEST REPEATED STRING #
// ##########################
func longestRepeatedSubString(str string) string {
	if len(str) == 0 {
		return ""
	}

	countN := 0
	subStr := ""

	for i := 0; i < len(str); i++ {
		str1 := string(str[:i])
		count := strings.Count(str, string(str1))
		if subStr == "" {
			countN = count
			subStr = string(str1)
			continue
		}

		if count >= countN {
			countN = count
			subStr = str
		}
	}
	return subStr
}

// ##########################
// #   STRING COMPRESSION    #
// ##########################
func stringCompression(text string) {
	var output = make([]string, 1)
	strMap := make(map[byte]int)

	for i := 0; i < len(text); i++ {
		if strMap[text[i]] != 0 || string(text[i]) == " " {
			continue
		} else {
			output = append(output, string(text[i]))
			output = append(output, strconv.Itoa(strings.Count(text, string(text[i]))))
			strMap[text[i]] = strings.Count(text, string(text[i]))
		}
	}
	fmt.Println("Compressed String:", strings.Join(output, ""))
}

// ##########################
// #   STRING DECOMPRESSION  #
// ##########################
func stringDeCompression(text string) {
	arr := make([]string, 1)
	for i := 0; i < len(text); i = i + 2 {
		arr = append(arr, string(text[i]))
		t, _ := strconv.Atoi(string(text[i+1]))
		for j := 0; j < t-1; j++ {
			arr = append(arr, string(text[i]))
		}
	}
	fmt.Println("Decompressed String:", strings.Join(arr, ""))
}

// ##########################
// #    REPLACE SUBSTRING    #
// ##########################
func replaceStr(text, old, new string) string {
	return strings.ReplaceAll(text, old, new)
}
