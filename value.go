package gotools

func ValueWithDefault[T any](val *T, def T) T {
	if val == nil {
		return def
	}
	return *val
}
