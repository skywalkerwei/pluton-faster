package test

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"testing"
)

func TestMapReduceVoid(t *testing.T) {
	//场景 批量处理数据
	homestayIds := []int64{1, 2, 3, 4, 5, 6}
	var newList []int64
	//generate GenerateFunc, mapper MapperFunc, reducer VoidReducerFunc
	err := mr.MapReduceVoid(func(source chan<- interface{}) {
		for _, id := range homestayIds {
			source <- id
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		id := item.(int64)
		fmt.Println("处理对象 ===", id)
		homestay := map[string]interface{}{
			"id": id,
		}
		writer.Write(homestay)
	}, func(pipe <-chan interface{}, cancel func(error)) {
		for item := range pipe {
			fmt.Println("pipe item ===", item)
			aa := item.(map[string]interface{})
			newList = append(newList, aa["id"].(int64))
		}
	})

	if err != nil {
		fmt.Println("mr err ===", err)
	} else {
		fmt.Println("mr newList ===", newList)
	}

}

func TestFinish(t *testing.T) {
	homestay := map[string]interface{}{
		"id": 1,
	}
	//场景 例如 商品详情依赖了多个服务获取数据，做并发的依赖处理
	err := mr.Finish(func() (err error) {
		homestay["User"] = "UserInfo"
		//rpc 服务
		return
	}, func() (err error) {
		homestay["Order"] = "OrderInfo"
		return
	}, func() (err error) {
		homestay["Carts"] = "CartsInfo"
		return
	})

	if err != nil {
		fmt.Println("mr err ===", err)
	} else {
		fmt.Println("mr homestay ===", homestay)
	}

}
