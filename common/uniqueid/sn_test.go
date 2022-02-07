package uniqueid

import "testing"

func TestGenSn(t *testing.T) {
	s := GenSn("O_")
	println(s)
}
