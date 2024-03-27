package gox

func StringMatcher(test string) func(v string) bool {
	return func(v string) bool {
		return v == test
	}
}

func NumberMatcher[T Number](test T) func(v T) bool {
	return func(v T) bool {
		return v == test
	}
}
