package test

import (
	"testing"
)

func TestT(t *testing.T) {
	
	//sed -i "" 's/,omitempty//g' *.pb.go

	//t.Error("除法函数测试没通过")
	//t.Logf("test add succ")
	//uniqueid.GenSn("22")
	//ch := make(chan int, 1000)
	//ch := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//}()
	//go func() {
	//	for {
	//		a, ok := <-ch
	//		if !ok {
	//			fmt.Println("close")
	//			return
	//		}
	//		fmt.Println("a: ", a)
	//	}
	//}()
	//defer close(ch)
	//fmt.Println("ok")
	//time.Sleep(time.Second * 10)
}

func try() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
}
