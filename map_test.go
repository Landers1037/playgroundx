/*
Name: playgroundx
File: map_test.go
Author: Landers
*/

package playgroundx

import "testing"

// 哈希表的使用
func TestCreateHash(t *testing.T) {
	var m map[string]string
	t.Logf("%+v", m)
	// map的初始化必须使用显式的make
	m = make(map[string]string)
	// add key
	m["key1"] = "value1"
	t.Logf("%+v", m)

	// 键不存在时会有第二个返回值
	if v, ok := m["key1"]; !ok {
		t.Fatal("key not exist")
	}else {
		t.Logf("key is %s value: %s", "key1", v)
	}
	if _, ok := m["key2"]; !ok {
		t.Fatal("key not exist")
	}
}
