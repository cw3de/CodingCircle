package maximum

// langsam (O(n*k)) aber garantiert korrekt
func SlowRunningMaximum(values []int, k int) []int {
	if k > len(values) {
		return []int{}
	}

	result := make([]int, 0, len(values)-k+1)

	for i := k; i <= len(values); i++ {
		max := values[i-k]
		for _, v := range values[i-k+1 : i] {
			if v > max {
				max = v
			}
		}
		result = append(result, max)
	}
	return result
}

// Pair aus Wert und Anzahl
type maximumCount struct {
	Value int
	Count int
}

type maximumQueue []maximumCount

func addValue(q maximumQueue, value int) maximumQueue {
	newCount := 1
	queueLen := len(q)

	// solange der neue Wert größer ist als der letzte Wert in der Queue
	// wird der letzte Wert aus der Queue entfernt und die Anzahl der
	// Wiederholungen addiert
	for queueLen > 0 && value >= q[queueLen-1].Value {
		newCount += q[queueLen-1].Count
		queueLen--
	}

	// der neue Wert wird an die Queue angehängt, ggfs. mit der neuen Anzahl
	return append(q[:queueLen], maximumCount{value, newCount})
}

func getNextMaximum(q maximumQueue) (maximumQueue, int) {
	if len(q) == 0 {
		panic("queue is empty")
	}

	maximum := q[0].Value
	// solange der Zähler des ersten Elements größer als 1 ist, wird der
	// Zähler dekrementiert, ansonsten wird das erste Element entfernt
	if q[0].Count > 1 {
		q[0].Count--
	} else {
		q = q[1:]
	}
	return q, maximum
}

func FastRunningMaximum(values []int, k int) []int {
	if k <= 0 || k > len(values) {
		return []int{}
	}

	result := make([]int, 0, len(values)-k+1)
	var queue maximumQueue

	for i, v := range values {
		queue = addValue(queue, v)

		// für die ersten k-1 Elemente gibt es noch keinen Maximum
		if i >= k-1 {
			newqueue, maximum := getNextMaximum(queue)
			result = append(result, maximum)
			queue = newqueue
		}
	}
	return result
}

func GoloRunningMaximum(values []int, k int) []int {
	if k <= 0 || k > len(values) {
		return []int{}
	}

	result := make([]int, 0, len(values)-k+1)
	var window []int

	for i, v := range values {
		// Werte entfernen, die außerhalb des Fensters liegen
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		// Werte aus dem Fenster entfernen, die kleiner als der neue Wert sind
		for len(window) > 0 && values[window[len(window)-1]] <= v {
			window = window[:len(window)-1]
		}

		// Einfügen des neuen Werts
		window = append(window, i)

		if i >= k-1 {
			result = append(result, values[window[0]])
		}
	}
	return result
}
