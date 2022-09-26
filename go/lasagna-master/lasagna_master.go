package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(ingredients []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(ingredients) * averagePreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (int, float64) {
	noodleGrams, sauceLiters := 0, 0.0
	for _, v := range ingredients {
		if v == "noodles" {
			noodleGrams += 50
		} else if v == "sauce" {
			sauceLiters += 0.2
		}
	}
	return noodleGrams, sauceLiters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends []string, my []string) []string {
	my_last_index := len(my) - 1
	friends_last_index := len(friends) - 1
	my[my_last_index] = friends[friends_last_index]
	return my
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	scaledQuantities := []float64{}
	scale := float64(portion) / 2
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, v*scale)
	}
	return scaledQuantities
}
