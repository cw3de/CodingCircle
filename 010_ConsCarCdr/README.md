Gegeben ist:

type PairFunc func(int,int) int

func Cons(a,b int) func(PairFunc) int {
  return func(f PairFunc) int {
    return f(a,b)
  }
}

Gesucht ist:

Funktion Car und Cdr, so dass gilt:
- Car(Cons(3,4)) => 3
- Cdr(Cons(3,4)) => 4

