package geo

type Node struct {
	Name            string
	NodesToTheSouth []*Node
	NodesToTheWest  []*Node
}

func (a *Node) IsNorthOf(b *Node) bool {
	for _, n := range a.NodesToTheSouth {
		if n == b {
			return true
		}
	}
	return false
}

func (a *Node) IsEastOf(b *Node) bool {
	for _, e := range a.NodesToTheWest {
		if e == b {
			return true
		}
	}
	return false
}

func (a *Node) NorthOf(b *Node) {
	a.NodesToTheSouth = append(a.NodesToTheSouth, b)
}

func (a *Node) EastOf(b *Node) {
	a.NodesToTheWest = append(a.NodesToTheWest, b)
}
