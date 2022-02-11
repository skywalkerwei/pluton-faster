package test

import (
	"fmt"
	"github.com/didi/gendry/builder"
	"testing"
)

func TestDidi(t *testing.T) {
	where := map[string]interface{}{
		"city": []string{"beijing", "shanghai"},
		// 默认可以省略 in 操作符，等同于:
		// "city in": []string{"beijing", "shanghai"},
		"score":    5,
		"age >":    35,
		"address":  builder.IsNotNull,
		"_orderby": "bonus desc",
		"_groupby": "department",
	}
	table := "some_table"
	selectFields := []string{"name", "age", "sex"}
	cond, values, err := builder.BuildSelect(table, where, selectFields)
	fmt.Println(cond, values, err)

}
