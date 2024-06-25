package fraction

import (
	"fmt"
)

type Fraction struct {
	Num   int // Zähler
	Denom int // Nenner
}

func New(z, n int) Fraction {
	return Fraction{z, n}
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Num, f.Denom)
}

func (f Fraction) ToFloat() float64 {
	return float64(f.Num) / float64(f.Denom)
}

func (f Fraction) MakeSimple() []Fraction {
	rest := f
	result := []Fraction{}
	for denominator := 2; rest.Num > 0 && denominator < 9999; denominator++ {
		// der Zähler der Differenz
		diffNum := rest.Num*denominator - rest.Denom
		// muss größer oder gleich 0 sein
		if diffNum >= 0 {
			// jetzt erst den Nenner der Differenz ermitteln
			rest = New(diffNum, rest.Denom*denominator)
			result = append(result, New(1, denominator))
		}
	}
	if rest.Num > 0 {
		fmt.Printf("Warning: %v could not be made simple\n", f)
	}
	return result
}
