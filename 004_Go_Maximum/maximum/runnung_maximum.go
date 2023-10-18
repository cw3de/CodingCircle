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

type maximumCount struct {
	Value int
	Count int
}

type maximumQueue []maximumCount

func NewQueue() *maximumQueue {
	return &maximumQueue{}
}

func addValue(q maximumQueue, value int) maximumQueue {
	newCount := 1
	queueLen := len(q)

	// fmt.Printf("add %d to queue len = %d\n", value, queueLen)

	for queueLen > 0 && value >= q[queueLen-1].Value {
		// fmt.Printf("replace %d (%d) with %d\n", q[queueLen-1].Value, q[queueLen-1].Count, value)

		newCount += q[queueLen-1].Count
		queueLen--
	}
	// fmt.Printf("append %d (%d)\n", value, newCount)

	return append(q[:queueLen], maximumCount{value, newCount})
}

func getNextMaximum(q maximumQueue) (maximumQueue, int) {
	if len(q) == 0 {
		panic("queue is empty")
	}

	maximum := q[0].Value
	if q[0].Count > 1 {
		q[0].Count--
	} else {
		// remove first element
		q = q[1:]
	}
	return q, maximum
}

func FastRunningMaximum(values []int, k int) []int {
	if k > len(values) {
		return []int{}
	}

	result := make([]int, 0, len(values)-k+1)
	var queue maximumQueue

	// TODO: Implement
	for i, v := range values {
		queue = addValue(queue, v)

		if i >= k-1 {
			newqueue, maximum := getNextMaximum(queue)
			result = append(result, maximum)
			queue = newqueue
		}
	}
	return result
}
