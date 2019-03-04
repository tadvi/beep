// +build !windows

package beep

func Beep(freq, duration int) (err error) {
	return nil
}

func Alert()   {}
func Error()   {}
func Stop()    {}
func Warning() {}

func Play(typ int) (err error) {
	return nil
}
