package bench

type Counter struct {
	A int
	B int
}

func increment(counter *Counter) {
	counter.A++
	counter.B++
}
