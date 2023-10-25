package tictactoe

type Field struct {
	Row, Col int
}

// Annahme: die Felder sind nicht doppelt in der Liste vorhanden
func IstGewonnen(fields []Field, size int) bool {

	// Speicherverbrauch O(n)
	countRows := make([]int, size)
	countCols := make([]int, size)
	countTopLeftBottomRight := 0
	countTopRightBottomLeft := 0

	// nur eine Schleife, Laufzeit O(n*n)
	for _, field := range fields {
		countRows[field.Row]++
		countCols[field.Col]++
		if field.Row == field.Col {
			countTopLeftBottomRight++
		}
		if field.Row+field.Col == size-1 {
			countTopRightBottomLeft++
		}

		if countRows[field.Row] == size ||
			countCols[field.Col] == size ||
			countTopLeftBottomRight == size ||
			countTopRightBottomLeft == size {
			return true
		}
	}
	return false
}
