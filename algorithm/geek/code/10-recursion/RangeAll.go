package recursion

import (
	"fmt"
)

// RangeType 实现一组数据集合全排列
type RangeType struct {
	value []interface{}
}

// NewRangeType create a new RangeType object
func NewRangeType(n int) *RangeType {
	return &RangeType{
		make([]interface{}, n),
	}
}

// 没理解这个算法……
// 迭代实现全排列
//
func (slice *RangeType) RangeAll(start int) {
	len := len(slice.value)
	if start == len-1 {
		fmt.Println(slice.value)
	}

	for i:=start; i<len; i++ {
		if i==start || slice.value[i] != slice.value[start] {
			slice.value[i], slice.value[start] = slice.value[start], slice.value[i]
			slice.RangeAll(start+1)
			slice.value[i], slice.value[start] = slice.value[start], slice.value[i]
		}
	}
}