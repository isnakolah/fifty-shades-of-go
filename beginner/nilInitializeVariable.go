package beginner

func ReturnsNil() interface{} {
	var x interface{} = nil
	_ = x
	return x
}
