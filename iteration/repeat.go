package iteration

func Repeat(character string, repeatCount int) string {
	var repeated string
	var i int

	for i < repeatCount {
		repeated += character
		i++
	}

	return repeated
}
