package panic

func noop() int {
	defer func() {
		recover()
	}()
	panic(-1)
}
