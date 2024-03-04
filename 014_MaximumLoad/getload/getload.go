package getload

type Event struct {
	Time  int64
	Type  string
	Count int
}

type Load struct {
	MaximumCount int
	From         int64
	To           int64
}

func GetMaximumLoad(events []Event) Load {
	var maxload Load
	currentLoad := 0

	for _, event := range events {
		if event.Type == "joined" {
			currentLoad += event.Count
		} else {
			currentLoad -= event.Count
		}
		// fmt.Printf("event: %v, currentLoad: %v\n", event, currentLoad)

		if currentLoad > maxload.MaximumCount {
			maxload.MaximumCount = currentLoad
			maxload.From = event.Time
			maxload.To = 0
		} else if maxload.To == 0 {
			maxload.To = event.Time
		}
	}

	return maxload
}
