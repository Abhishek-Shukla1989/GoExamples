// ##########################
// #    PALINDROME CHECK     #
// ##########################

import (
    "fmt"
    "slices"
    "strings"
    "unicode"
)

func CheckPalindrom() {
    // Palindrome check by reversing a string
    str := "heleh"
    reverseStr := strings.Split(str, "")
    slices.Reverse(reverseStr)
    newStr := strings.Join(reverseStr, "")

    if strings.EqualFold(str, newStr) {
        fmt.Println("String is palindrome")
    } else {
        fmt.Println("String is not palindrome")
    }

    strSlice := []byte(str) // Creates a deep copy
    slices.Reverse(strSlice)
    str1 := string(strSlice)

    if strings.EqualFold(str, str1) {
        fmt.Println("String is palindrome")
    } else {
        fmt.Println("String is not palindrome")
    }
}

// ##########################
// #    STRING LENGTH WITHOUT BUILTIN LEN FUNCTION     #
// ##########################

func StringLength() {
    count := 0
    newStr := "abhishekShukla"
    for range newStr {
        count++
    }
    fmt.Printf("String length is %d\n", count)
}

// ##########################
// #  STRING CONCATENATION   #
// ##########################

func StringConcat() {
    str_one := "abhishek"
    str_two := "shukla"
    s := str_one + " " + str_two
    fmt.Println(s)
}

// ##########################
// #   VOWELS COUNT          #
// ##########################

func CountVowels(str_one string) {
    f := func(r rune) bool {
        r = unicode.ToLower(r)
        return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
    }

    countN := 0
    for _, r := range str_one {
        if f(r) {
            countN++
        }
    }
    fmt.Printf("Count of vowels is %d\n", countN)
}

// ##########################
// # REMOVE VOWELS FROM STR  #
// ##########################

func RemoveVowels(str_one string) {
    var result []rune
    f := func(r rune) bool {
        r = unicode.ToLower(r)
        return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
    }

    for _, r := range str_one {
        if !f(r) {
            result = append(result, r)
        }
    }
    finalStr := string(result)
    fmt.Printf("String after removing vowels: %q\n", finalStr)
}

// ##########################
// #  REMOVE WHITESPACES     #
// ##########################

func RemoveWhitespaces() {
    str_one1 := "  abhishek   "
    str_one1 = strings.TrimSpace(str_one1)
    fmt.Printf("Whitespace removed: %q\n", str_one1)
}

// ##########################
// #   SUBSTRING SEARCH      #
// ##########################

func SubstringSearch() {
    sunstr := "K"
    str := strings.Contains(strings.ToLower("Abhishek"), strings.ToLower(sunstr))
    fmt.Println(str) // Outputs true if substring is found, false otherwise
}
