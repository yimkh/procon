package model

//Environment is all variable
type Environment struct {
	In     int
	Out    int
	Buffer [5]int
	Mutex  int
	Empty  int
	Full   int
}

//Wait is to wait
func Wait(x *int) bool {
	if *x == 0 {
		return false
	}

	*x--

	return true
}

//Signal is to signal
func Signal(x *int) {
	*x++
}
