package raindrops

import "strconv"

func Convert(number int) string {
	var raindrop string

	if number%3 == 0 {
		raindrop += "Pling"
	}
	if number%5 == 0 {
		raindrop += "Plang"
	}
	if number%7 == 0 {
		raindrop += "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		raindrop = strconv.Itoa(number)
	}
	return raindrop
}
