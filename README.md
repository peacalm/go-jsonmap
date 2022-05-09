# jsonmap for Golang

Get value with specific type(string/bool/float64/float32/int64/uint64/int32/uint32/int/uint etc) from json.

Returns a 3-item-tuple (value, whether found, error).

## Example
```go
package main

import (
	"fmt"
	"github.com/peacalm/go-jsonmap"
)

func main() {
	data := `{"b":true, "i":1, "f":1.3, "s":"str", "long":7095620078347567873, "sub":{"i":2}, "arr":[1,2]}`
	jm, err := jsonmap.Unmarshal([]byte(data), false)
	if err != nil {
		fmt.Printf("jsonmap.Numarshal failed: %v. useNumber = false, string = %v", err, data)
	}
	{
		val, found, err := jm.GetBool("b", false)
		fmt.Println("GetBool: ", val, found, err)
	}
	{
		val, found, err := jm.GetInt("i", 0)
		fmt.Println("GetInt: ", val, found, err)
	}
	{
		val, found, err := jm.GetFloat64("f", 0)
		fmt.Println("GetFloat64: ", val, found, err)
	}
	{
		val, found, err := jm.GetString("s", 0)
		fmt.Println("GetString: ", val, found, err)
	}
	{
		val, found, err := jm.RGetInt([]string{"sub", "i"}, 0)
		fmt.Println("RGetInt: ", val, found, err)
	}
	{
		val, found, err := jm.GetIntSlice("arr", []int{})
		fmt.Println("GetIntSlice: ", val, found, err)

	}
	{
		long1, found, err := jm.GetInt64("long", 0) // NOTICE: precision lost
		fmt.Println("useNumber=false, GetInt64: ", long1, found, err)
		jm2, _ := jsonmap.Unmarshal([]byte(data), true) // NOTICE: no precision lost
		long2, found, err := jm2.GetInt64("long", 0)
		fmt.Println("useNumber=true,  GetInt64: ", long2, found, err)
		fmt.Println("long1 == long2 ? ", long1 == long2) // false
                
		// NOTICE: get int by float number
		// If useNumber=false: convert by int(${float number}), err will be nil
		val1, found, err := jm.GetInt("f", 0) // 1 true <nil>
		fmt.Println("Test GetInt by float, useNumber=False: ", val1, found, err)
		// If useNumber=true: convert by strconv.ParseInt("${string form of the float number}", 10, 64), err != nil
		val2, found, err := jm2.GetInt("f", 0) // 0 true strconv.ParseInt: parsing "1.3": invalid syntax
		fmt.Println("Test GetInt by float, useNumber=true:  ", val2, found, err)
	}
}
```
