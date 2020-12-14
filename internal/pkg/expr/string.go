package expr

type Str struct {
	Eq []string
}

func StrEq(val string) Str {
	return Str{Eq: []string{val}}
}

func StrIn(val []string) Str {
	return Str{Eq: val}
}
