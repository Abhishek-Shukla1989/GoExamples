// ##########################
// # CHECK IF PRIME         #
// ##########################
func IsPrime(n int) bool {
	if n < 0 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ##########################
// # FIND GCD               #
// ##########################
func gcd(a, b int) int {
	for b != 0 {
		a, b = a, a%b
	}
	return a
}

// ##########################
// # CALCULATE FACTORIAL    #
// ##########################
func factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// ##########################
// # SUM OF DIGITS          #
// ##########################
func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

// ##########################
// # REVERSE INTEGER        #
// ##########################
func reverseInt(n int) int {
	sum := 0
	sign := 1
	if n < 0 {
		n = -n
		sign = -1
	}
	for n > 0 {
		t := n % 10
		sum = sum*10 + t
		n = n / 10
	}
	return sum * sign
}

// ##########################
// # CHECK PALINDROME       #
// ##########################
func palindromInt(n int) bool {
	sum := 0
	if n < 0 {
		n = -n
	}
	originalNum := n
	for n > 0 {
		t := n % 10
		sum = sum*10 + t
		n = n / 10
	}
	return sum == originalNum
}

// ##########################
// # INTEGER TO BINARY/HEX  #
// ##########################
func toBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func toHexadecimal(n int) string {
	return strconv.FormatInt(int64(n), 16)
}

// ##########################
// # INTEGER TO STRING      #
// ##########################
func intToString(n int) string {
	return strconv.Itoa(n)
}

// ##########################
// # STRING TO INTEGER      #
// ##########################
func stringToInt(n string) int {
	if i, err := strconv.Atoi(n); err == nil {
		return i
	}
	return 0
}

// ##########################
// # GENERATE FIBONACCI     #
// ##########################
func fibonacciSeries(n int) []int {
	fibSeries := make([]int, n)
	if n > 0 {
		fibSeries[0] = 0
	}
	if n > 1 {
		fibSeries[1] = 1
	}
	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}
	return fibSeries
}

// ##########################
// # CALCULATE POWER        #
// ##########################
func power(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
