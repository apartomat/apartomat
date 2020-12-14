package expr

type Int struct {
	Eq  []int
	Gt  *int
	Gte *int
}

func IntEq(val int) Int {
	return Int{Eq: []int{val}}
}
