package plzsort

type plzList []int

func PlzSort(unsorted []int) []int {
	plzhash := make(map[int]plzList, 99999)

	for _, v := range unsorted {
		plzhash[v] = append(plzhash[v], v)
	}

	sorted := make([]int, 0, len(unsorted))

	// Konstante Laufzeit, da die Schleife immer 99999 mal durchlaufen wird.
	for plz := 0; plz < 100000; plz++ {
		list, exists := plzhash[plz]
		if exists {
			sorted = append(sorted, list...)
		}
	}

	return sorted
}
