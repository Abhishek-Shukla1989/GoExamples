// ##########################
// #    CASE CONVERSION      #
// ##########################
func toUpperCase(text string) string {
	return strings.ToUpper(text)
}

func toLowerCase(text string) string {
	return strings.ToLower(text)
}

// ##########################
// #  TITLE CASE CONVERSION  #
// ##########################
func toTitle(text string) string {
	words := strings.Fields(text)
	for i, txt := range words {
		if len(txt) == 1 {
			words[i] = strings.ToUpper(txt)
		} else {
			words[i] = string(unicode.ToUpper(rune(txt[0]))) + strings.ToLower(txt[1:])
		}
	}
	return strings.Join(words, " ")
}

// ##########################
// #    ROTATE STRING        #
// ##########################
func rotateStringLeft(text string, n int) string {
	if n >= len(text) {
		return ""
	}
	return text[n:] + text[:n]
}

// ##########################
// #  REMOVE DUPLICATES      #
// ##########################
func removeDuplicates(text string) string {
	if len(text) <= 1 {
		return text
	}
	seen := make(map[rune]bool)
	result := ""
	for _, char := range text {
		if !seen[unicode.ToLower(char)] {
			result += string(char)
			seen[unicode.ToLower(char)] = true
		}
	}
	return result
}

// ##########################
// # REMOVE SPECIFIC CHARS   #
// ##########################
func removeSpecificChars(text string) string {
	f := func(r rune) rune {
		if unicode.ToLower(r) == 'a' || unicode.ToLower(r) == 'b' {
			return -1
		}
		return r
	}
	return strings.Map(f, text)
}

// ##########################
// #   STRING TO INTEGER     #
// ##########################
func atoi(s string) (int, error) {
	sign := 1
	num := 0
	start := 0

	if len(s) == 0 {
		return 0, errors.New("empty string")
	}
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, errors.New("invalid input")
		}
		num = num*10 + int(s[i]-'0')
	}
	return sign * num, nil
}

// ##########################
// #   INTEGER TO STRING     #
// ##########################
func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	result := ""
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	return sign + result
}

// ##########################
// #    EVALUATE EXPRESSION  #
      	Parse and evaluate a simple arithmetic expression given as a string (e.g., "3+5-2").
// ##########################
func evaluateExpression(expr string) int {
	result := 0
	currentNum := 0
	sign := 1

	for i := 0; i < len(expr); i++ {
		char := expr[i]

		if unicode.IsDigit(rune(char)) {
			currentNum = currentNum*10 + int(char-'0')
		}

		if char == '+' || char == '-' || i == len(expr)-1 {
			result += sign * currentNum
			currentNum = 0

			if char == '+' {
				sign = 1
			} else if char == '-' {
				sign = -1
			}
		}
	}

	return result
}

// ##########################
// #  SPLIT STRING BY DELIM  #
// ##########################
func splitString(str, delimiter string) []string {
	return strings.Split(str, delimiter)
}

// ##########################
// #  JOIN SLICE OF STRINGS  #
// ##########################
func joinStrings(sliceArr []string, separator string) string {
	return strings.Join(sliceArr, separator)
}

// ##########################
// #    REPLACE SUBSTRING    #
// ##########################
func replaceStr(text, old, new string) string {
	return strings.ReplaceAll(text, old, new)
}
