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

func TestMap(t *testing.T) {
	mp := map[string]interface{}{
		"country": "China",
		"role":    "driver",
		"age >":   45,
		"_or": []map[string]interface{}{
			{
				"x1":    11,
				"x2 >=": 45,
			},
			{
				"x3":    "234",
				"x4 <>": "tx2",
			},
		},
		"_groupby": "name",
		"_having": map[string]interface{}{
			"total >":  1000,
			"total <=": 50000,
		},
		"_orderby": "age desc",
	}
	cond, vals, err := builder.BuildSelect("tableName", mp, []string{"name", "count(price) as total", "age"})
	fmt.Println(cond, vals, err)

	//SELECT name,count(price) as total,age FROM tableName
	//WHERE (((x1=? AND x2>=?) OR (x3=? AND x4!=?)) AND country=? AND role=? AND age>?)
	//GROUP BY name
	//HAVING (total>? AND total<=?)
	//ORDER BY age desc

	//vals: []interface{}{11, 45, "234", "tx2", "China", "driver", 45, 1000, 50000}

	//"_limit": []uint{a,b} => LIMIT a,b
	//"_limit": []uint{a} => LIMIT 0,a

}
