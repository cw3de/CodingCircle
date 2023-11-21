package factors

func GetFactors(list []int) []int {
	result := make([]int, len(list))

	// um eine verschachtelte Schleife zu vermeiden, müssen wir das Ergebnis
	// aus der Zeile davor, bzw. darüber wieder verwenden

	// i: 0 . 1 . 2 .  3 .  4 .  5
	// V: 7   8   9   10   11   12
	// 0:     8 * 9 * 10 * 11 * 12 ; before[0] = 1              ; after[0] =  8 * after[1]
	// 1: 7       9 * 10 * 11 * 12 ; before[1] = before[0] *  7 ; after[1] =  9 * after[2]
	// 2: 7 * 8       10 * 11 * 12 ; before[2] = before[1] *  8 ; after[2] = 10 * after[3]
	// 3: 7 * 8 * 9        11 * 12 ; before[3] = before[2] *  9 ; after[3] = 11 * after[4]
	// 4: 7 * 8 * 9 * 10        12 ; before[4] = before[3] * 10 ; after[4] = 12 * after[5]
	// 5: 7 * 8 * 9 * 10 * 11      ; before[5] = before[4] * 11 ; after[5] =  1

	// einmal von vorne nach hinten für den linken Teil ...
	before := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		if i == 0 {
			before[i] = 1
		} else {
			before[i] = before[i-1] * list[i-1]
		}
	}

	// ... und einmal von hinten nach vorne für den rechten Teil
	after := make([]int, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		if i == len(list)-1 {
			after[i] = 1
		} else {
			after[i] = after[i+1] * list[i+1]
		}

		// das Ergebnis ist das Produkt aus beiden Teilen
		result[i] = before[i] * after[i]
	}

	return result
}
