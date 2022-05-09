package jsonmap_test

import (
	"fmt"
	"testing"

	"github.com/peacalm/go-jsonmap"
)

func TestGet(t *testing.T) {
	data := `{"b1":true, "b2":false, "i":1, "f":1.3, "s":"some-string", "longint":7095620078347567873}`
	jm1, err := jsonmap.Unmarshal([]byte(data), false)
	if err != nil {
		t.Fatalf("jsonmap.Numarshal failed: %v. useNumber = false, string = %v", err, data)
	}
	jm2, err := jsonmap.Unmarshal([]byte(data), true)
	if err != nil {
		t.Fatalf("jsonmap.Numarshal failed: %v. useNumber = true, string = %v", err, data)
	}

	// bool
	testGet(t, jm1, jm2, "b1", false, true, true, false)
	testGet(t, jm1, jm2, "b1", true, true, true, false)
	testGet(t, jm1, jm2, "b2", false, false, true, false)
	testGet(t, jm1, jm2, "b2", true, false, true, false)
	testGet(t, jm1, jm2, "b3", false, false, false, false)
	testGet(t, jm1, jm2, "b3", true, true, false, false)
	testGet(t, jm1, jm2, "i", false, false, true, true)
	testGet(t, jm1, jm2, "i", true, true, true, true)

	// string
	testGet(t, jm1, jm2, "s", "", "some-string", true, false)
	testGet(t, jm1, jm2, "s", "def", "some-string", true, false)
	testGet(t, jm1, jm2, "none", "", "", false, false)
	testGet(t, jm1, jm2, "none", "def", "def", false, false)
	testGet(t, jm1, jm2, "i", "", "", true, true)
	testGet(t, jm1, jm2, "i", "def", "def", true, true)
	testGet(t, jm1, jm2, "f", "", "", true, true)
	testGet(t, jm1, jm2, "f", "def", "def", true, true)
	testGet(t, jm1, jm2, "b1", "", "", true, true)
	testGet(t, jm1, jm2, "b1", "def", "def", true, true)

	// int
	testGet(t, jm1, jm2, "i", 0, 1, true, false)
	testGet(t, jm1, jm2, "i", 1, 1, true, false)
	testGet(t, jm1, jm2, "none", 0, 0, false, false)
	testGet(t, jm1, jm2, "none", -1, -1, false, false)
	f, _, _ := jm1.GetFloat64("f", 0.0)
	{
		v, f, e := jm1.GetInt("f", 0)
		fmt.Println("GetInt from float number 1.3, with useNumber=false:", v, f, e)
	}
	{
		v, f, e := jm2.GetInt("f", 0)
		fmt.Println("GetInt from float number 1.3, with useNumber=true: ", v, f, e)
	}
	testGetInt(t, jm1, "f", 0, int(f), true, false) // NOTICE: no error, but precision lost
	testGetInt(t, jm2, "f", 0, 0, true, true)       // NOTICE: error, return 0
	testGet(t, jm1, jm2, "b1", -1, -1, true, true)

	// 64bit long int
	longint1, _, _ := jm1.GetInt64("longint", 0)
	longint2, _, _ := jm2.GetInt64("longint", 0)
	fmt.Println("GetInt64 with useNumber=false:", longint1)
	fmt.Println("GetInt64 with useNumber=true: ", longint2, ". eq? ", longint1 == longint2)

}

func testGet(t *testing.T, jm1, jm2 jsonmap.JsonMap, key string, def, expectedVal interface{}, keyExists, hasErr bool) {
	if _, ok := def.(bool); ok {
		testGetBool(t, jm1, key, def.(bool), expectedVal.(bool), keyExists, hasErr)
		testGetBool(t, jm2, key, def.(bool), expectedVal.(bool), keyExists, hasErr)
	}
	if _, ok := def.(string); ok {
		testGetString(t, jm1, key, def.(string), expectedVal.(string), keyExists, hasErr)
		testGetString(t, jm2, key, def.(string), expectedVal.(string), keyExists, hasErr)
	}
	if _, ok := def.(int); ok {
		testGetInt(t, jm1, key, def.(int), expectedVal.(int), keyExists, hasErr)
		testGetInt(t, jm2, key, def.(int), expectedVal.(int), keyExists, hasErr)
	}
	if _, ok := def.(float64); ok {
		testGetFloat64(t, jm1, key, def.(float64), expectedVal.(float64), keyExists, hasErr)
		testGetFloat64(t, jm2, key, def.(float64), expectedVal.(float64), keyExists, hasErr)
	}
}

func testGetBool(t *testing.T, jm jsonmap.JsonMap, key string, def, expectedVal bool, keyExists, hasErr bool) {
	if v, f, e := jm.GetBool(key, def); v != expectedVal || f != keyExists || bool(e != nil) != hasErr {
		t.Fatalf("GetBool failed: got (%v, %v, %v), expect (%v, %v, hasErr:%v), key = %v, jm = %v", v, f, e, expectedVal, keyExists, hasErr, key, jm)
	}
}
func testGetString(t *testing.T, jm jsonmap.JsonMap, key string, def, expectedVal string, keyExists, hasErr bool) {
	if v, f, e := jm.GetString(key, def); v != expectedVal || f != keyExists || bool(e != nil) != hasErr {
		t.Fatalf("GetString failed: got (%v, %v, %v), expect = (%v, %v, hasErr:%v), key = %v, jm = %v", v, f, e, expectedVal, keyExists, hasErr, key, jm)
	}
}
func testGetInt(t *testing.T, jm jsonmap.JsonMap, key string, def, expectedVal int, keyExists, hasErr bool) {
	if v, f, e := jm.GetInt(key, def); v != expectedVal || f != keyExists || bool(e != nil) != hasErr {
		t.Fatalf("GetInt failed: got (%v, %v, %v), expect = (%v, %v, hasErr:%v), key = %v, jm = %v", v, f, e, expectedVal, keyExists, hasErr, key, jm)
	}
}
func testGetFloat64(t *testing.T, jm jsonmap.JsonMap, key string, def, expectedVal float64, keyExists, hasErr bool) {
	if v, f, e := jm.GetFloat64(key, def); v != expectedVal || f != keyExists || bool(e != nil) != hasErr {
		t.Fatalf("GetFloat64 failed: got (%v, %v, %v), expect = (%v, %v, hasErr:%v), key = %v, jm = %v", v, f, e, expectedVal, keyExists, hasErr, key, jm)
	}
}

func TestRGet(t *testing.T) {
	data := `{"a":{"b":{ "i":1,"b":true,"s":"str"}}}`
	jm1, err := jsonmap.Unmarshal([]byte(data), false)
	if err != nil {
		t.Fatalf("jsonmap.Numarshal failed: %v. useNumber = false, string = %v", err, data)
	}
	jm2, err := jsonmap.Unmarshal([]byte(data), true)
	if err != nil {
		t.Fatalf("jsonmap.Numarshal failed: %v. useNumber = false, string = %v", err, data)
	}
	i1, _, _ := jm1.RGetInt([]string{"a", "b", "i"}, 0)
	i2, _, _ := jm2.RGetInt([]string{"a", "b", "i"}, 0)
	b1, _, _ := jm1.RGetBool([]string{"a", "b", "b"}, false)
	b2, _, _ := jm2.RGetBool([]string{"a", "b", "b"}, false)
	s1, _, _ := jm1.RGetString([]string{"a", "b", "s"}, "")
	s2, _, _ := jm2.RGetString([]string{"a", "b", "s"}, "")
	if i1 != 1 || i2 != 1 || !b1 || !b2 || s1 != "str" || s2 != "str" {
		t.Fatal("RGet failed")
	}
}

const rep = 1000000
const jsonStrForPerfTest = `{"a":{"b":{"i":1234567890,"b":true,"s":"str"}},"i":1234567890,"b":true,"s":"str"}`

func Test_Performance_RGetInt_NotUseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), false)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.RGetInt([]string{"a", "b", "i"}, 0)
	}
}
func Test_Performance_RGetInt_UseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), true)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.RGetInt([]string{"a", "b", "i"}, 0)
	}
}
func Test_Performance_GetInt_NotUseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), false)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.GetInt("i", 0)
	}
}
func Test_Performance_GetInt_UseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), true)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.GetInt("i", 0)
	}
}

func Test_Performance_RGetString_NotUseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), false)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.RGetString([]string{"a", "b", "s"}, "")
	}
}
func Test_Performance_RGetString_UseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), true)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.RGetString([]string{"a", "b", "s"}, "")
	}
}
func Test_Performance_GetString_NotUseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), false)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.GetString("s", "")
	}
}
func Test_Performance_GetString_UseNumber(t *testing.T) {
	jm, _ := jsonmap.Unmarshal([]byte(jsonStrForPerfTest), true)
	for i := 0; i < rep; i++ {
		_, _, _ = jm.GetString("s", "")
	}
}
