
// ##########################
// # COMPARE FLOAT       #
// ###########################
func comparefloat(f1, f2, tolerance float64) bool {

	fmt.Println("Abs value ", math.Abs(f1-f2))
	return math.Abs(f1-f2) > tolerance
}

// ##########################
// # ROUND FLOAT      #
// ##########################
func roundFloat(f1, decimalplaces float64) float64 {

	shift := math.Pow(10, float64(decimalplaces))
	return math.Round(f1*shift) / shift
}

// ##########################
// # FLOAT TO STRING        #
// ##########################
func floatToString(num float64, precision int) string {
	return strconv.FormatFloat(num, 'f', precision, 64)
}

// ##########################
// # STRING TO FLOAT        #
// ##########################
func stringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// ##########################
// # SQUARE ROOT            #
// ##########################
func squareRoot(num float64) float64 {
	return math.Sqrt(num)
}

// ##########################
// # NATURAL LOGARITHM      #
// ##########################
func naturalLog(num float64) float64 {
	return math.Log(num)
}

// ##########################
// # EXPONENTIAL FUNCTION   #
// ##########################
func exponential(num float64) float64 {
	return math.Exp(num)
}

// ##########################
// # TRIGONOMETRIC FUNCTIONS #
// ##########################
func sine(angle float64) float64 {
	return math.Sin(angle)
}

func cosine(angle float64) float64 {
	return math.Cos(angle)
}

func tangent(angle float64) float64 {
	return math.Tan(angle)
}

// ##########################
// # AREA OF A CIRCLE       #
// ##########################
func areaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

// ##########################
// # VOLUME OF A SPHERE     #
// ##########################
func volumeOfSphere(radius float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
}

// ##########################
// # MAINTAIN FLOAT PRECISION #
// ##########################
func preciseSum(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// ##########################
// # CHANGE FLOAT TO INT #
// ##########################
func floattoInt(f1 float64) int {

	return int(math.Round(f1))
}
