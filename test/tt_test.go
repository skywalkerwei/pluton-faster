package test

import (
	"github.com/skywalkerwei/pluton-faster/common/uniqueid"
	"testing"
)

func TestGenSn(t *testing.T) {
	uniqueid.GenSn("22")
}
