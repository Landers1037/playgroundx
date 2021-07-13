/*
Name: playgroundx
File: slice_test.go
Author: Landers
*/

package playgroundx

import "testing"

// 针对数组和切片的简单练习
func TestCreateArray(t *testing.T) {
	// 数组的创建方式： 使用字面量创建
	var array1 = [3]int{1, 2, 3}
	t.Log(array1)
	// 数组的值内存地址一定是连续的
	for _, v := range array1 {
		t.Logf("address of %d is %p", v, &v)
	}
}

// 当元素小于4时go的优化会直接在栈上创建而不是静态分配
func TestCreateBigArray(t *testing.T) {
	var arrayBig = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(arrayBig)
	t.Logf("%p", &arrayBig)
	// 数组的值内存地址一定是连续的
	for index, v := range arrayBig {
		t.Logf("address of %d is %p", v, &arrayBig[index])
	}
}

// 切片的创建
// 可以使用三种方式创建
func TestCreateSlice(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var s1 = a[:] // a copy of a
	var s2 = make([]int, 3)
	var s3 []int // empty slice

	t.Logf("%+v\n", s1)
	t.Logf("%+v\n", s2)
	t.Logf("%+v\n", s3)

	for i:=0;i<4;i++ {
		s3 = append(s3, i)
	}
	t.Logf("%+v\n", s3)
}

// 切换的长度和容量
// 当切片来自一个固定数组时 其长度变化 容量初始化 = 原数组大小
func TestLenSlice(t *testing.T) {
	var a = [3]int{1, 2, 3}
	t.Logf("len %d cap %d\n", len(a), cap(a))
	var s = a[:2]
	t.Log(s)
	t.Logf("len %d cap %d\n", len(s), cap(s))
	// 对切片的扩容会导致容量增加 内存重新分配
	for i:=0;i<4;i++ {
		s = append(s, i)
	}
	t.Log(s)
	t.Logf("len %d cap %d\n", len(s), cap(s))
}