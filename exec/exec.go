package exec

func Loop(code func()) {
	While(func() bool {
		code()
		return true
	})
}

func While(code func() bool) {
	for {
		ok := while_(code)
		if ok == false {
			break
		}
	}
}

func while_(code func() bool) (b bool) {
	defer func() {
		err := recover()
		if err != nil {
			b = true
			return
		}
	}()

	return code()
}
