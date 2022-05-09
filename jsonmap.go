// Copyright (c) 2022 Shuangquan Li. All Rights Reserved.
//
// Licensed under the MIT License (the "License"); you may not use this file
// except in compliance with the License. You may obtain a copy of the License
// at
//
//   http://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package jsonmap

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type JsonMap map[string]interface{}

func (d JsonMap) String() string {
	return fmt.Sprintf("JsonMap(%v)", map[string]interface{}(d))
}

func (d JsonMap) ToMap() map[string]interface{} {
	return d
}

//// directly get, and specialization for string/bool/float64/float32/int64/uint64/int32/uint32/int/uint

// fetch origin value, no type assurance
func (d JsonMap) Get(key string, def interface{}) (val interface{}, found bool, err error) {
	if val, found = d[key]; found {
		return val, true, nil
	}
	return def, false, nil
}

// this method ensures val’s type is same as def
func (d JsonMap) GetAny(key string, def interface{}) (val interface{}, found bool, err error) {
	if raw, found := d[key]; found {
		val, err = toAny(raw, def)
		return val, true, err
	}
	return def, false, nil
}

func (d JsonMap) GetSubMap(key string, def JsonMap) (val JsonMap, found bool, err error) {
	raw, found := d[key]
	if !found {
		return def, false, nil
	}
	v, ok := raw.(map[string]interface{})
	if !ok {
		return def, true, fmt.Errorf("key %s type %T is not map", key, raw)
	}
	return JsonMap(v), true, nil
}

func (d JsonMap) GetString(key string, def string) (val string, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(string)
	return
}

func (d JsonMap) GetBool(key string, def bool) (val bool, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(bool)
	return
}

func (d JsonMap) GetFloat64(key string, def float64) (val float64, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(float64)
	return
}

func (d JsonMap) GetFloat32(key string, def float32) (val float32, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(float32)
	return
}

func (d JsonMap) GetInt64(key string, def int64) (val int64, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(int64)
	return
}

func (d JsonMap) GetUint64(key string, def uint64) (val uint64, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(uint64)
	return
}

func (d JsonMap) GetInt32(key string, def int32) (val int32, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(int32)
	return
}

func (d JsonMap) GetUint32(key string, def uint32) (val uint32, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(uint32)
	return
}

func (d JsonMap) GetInt(key string, def int) (val int, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(int)
	return
}

func (d JsonMap) GetUint(key string, def uint) (val uint, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.(uint)
	return
}

//// recursively get, and specialization for string/bool/float64/float32/int64/uint64/int32/uint32/int/uint

func (d JsonMap) rGet(keyPath []string, idx int, def interface{}) (val interface{}, found bool, err error) {
	k := keyPath[idx]
	v, found := d[k]
	if !found {
		return def, false, nil
	}
	if idx == len(keyPath)-1 {
		return v, true, nil
	}
	vmp, ok := v.(map[string]interface{})
	if !ok {
		return def, false, fmt.Errorf("key %s type %T is not map", keyPath[0:idx+1], v)
	}
	return JsonMap(vmp).rGet(keyPath, idx+1, def)
}

// fetch origin value, no type assurance
func (d JsonMap) RGet(keyPath []string, def interface{}) (val interface{}, found bool, err error) {
	if len(keyPath) == 0 {
		return def, false, fmt.Errorf("keyPath empty")
	}
	return d.rGet(keyPath, 0, def)
}

// this method ensures val’s type is same as def
func (d JsonMap) RGetAny(keyPath []string, def interface{}) (val interface{}, found bool, err error) {
	if len(keyPath) == 0 {
		return def, false, fmt.Errorf("keyPath empty")
	}
	val, found, err = d.rGet(keyPath, 0, def)
	if !found || err != nil {
		return
	}
	val, err = toAny(val, def)
	return
}

func (d JsonMap) RGetSubMap(keyPath []string, def JsonMap) (val JsonMap, found bool, err error) {
	raw, found, err := d.RGet(keyPath, def)
	if !found || err != nil {
		return def, found, err
	}
	v, ok := raw.(map[string]interface{})
	if !ok {
		return def, true, fmt.Errorf("key %s type %T is not map", keyPath, raw)
	}
	return JsonMap(v), true, nil
}

func (d JsonMap) RGetString(keyPath []string, def string) (val string, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(string)
	return
}

func (d JsonMap) RGetBool(keyPath []string, def bool) (val bool, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(bool)
	return
}

func (d JsonMap) RGetFloat64(keyPath []string, def float64) (val float64, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(float64)
	return
}

func (d JsonMap) RGetFloat32(keyPath []string, def float32) (val float32, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(float32)
	return
}

func (d JsonMap) RGetInt64(keyPath []string, def int64) (val int64, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(int64)
	return
}

func (d JsonMap) RGetUint64(keyPath []string, def uint64) (val uint64, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(uint64)
	return
}

func (d JsonMap) RGetInt32(keyPath []string, def int32) (val int32, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(int32)
	return
}

func (d JsonMap) RGetUint32(keyPath []string, def uint32) (val uint32, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(uint32)
	return
}

func (d JsonMap) RGetInt(keyPath []string, def int) (val int, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(int)
	return
}

func (d JsonMap) RGetUint(keyPath []string, def uint) (val uint, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.(uint)
	return
}

//// directly get slice, and specialization for string/bool/float64/float32/int64/uint64/int32/uint32/int/uint

func (d JsonMap) GetSlice(key string, def []interface{}) (val []interface{}, found bool, err error) {
	raw, found, err := d.GetAny(key, def)
	val = raw.([]interface{})
	return
}

func (d JsonMap) GetAnySlice(key string, def []interface{}, itemType interface{}) (
	val []interface{}, found bool, err error) {
	raw, found, err := d.GetSlice(key, def)
	if !found || err != nil {
		return def, found, err
	}
	if len(raw) == 0 {
		return raw, found, err
	}
	ret := make([]interface{}, 0)
	for idx, i := range raw {
		v, e := toAny(i, itemType)
		if e != nil {
			err = fmt.Errorf("key %s index %d: %v", key, idx, e)
			return def, true, err
		}
		ret = append(ret, v)
	}
	return ret, true, nil
}

func (d JsonMap) GetStringSlice(key string, def []string) (val []string, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, string(""))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]string, 0)
	for _, i := range raw {
		val = append(val, i.(string))
	}
	return
}

func (d JsonMap) GetBoolSlice(key string, def []bool) (val []bool, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, bool(false))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]bool, 0)
	for _, i := range raw {
		val = append(val, i.(bool))
	}
	return
}

func (d JsonMap) GetFloat64Slice(key string, def []float64) (val []float64, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, float64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]float64, 0)
	for _, i := range raw {
		val = append(val, i.(float64))
	}
	return
}

func (d JsonMap) GetFloat32Slice(key string, def []float32) (val []float32, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, float32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]float32, 0)
	for _, i := range raw {
		val = append(val, i.(float32))
	}
	return
}

func (d JsonMap) GetInt64Slice(key string, def []int64) (val []int64, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, int64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int64, 0)
	for _, i := range raw {
		val = append(val, i.(int64))
	}
	return
}

func (d JsonMap) GetUint64Slice(key string, def []uint64) (val []uint64, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, uint64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint64, 0)
	for _, i := range raw {
		val = append(val, i.(uint64))
	}
	return
}

func (d JsonMap) GetInt32Slice(key string, def []int32) (val []int32, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, int32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int32, 0)
	for _, i := range raw {
		val = append(val, i.(int32))
	}
	return
}

func (d JsonMap) GetUint32Slice(key string, def []uint32) (val []uint32, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, uint32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint32, 0)
	for _, i := range raw {
		val = append(val, i.(uint32))
	}
	return
}

func (d JsonMap) GetIntSlice(key string, def []int) (val []int, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, int(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int, 0)
	for _, i := range raw {
		val = append(val, i.(int))
	}
	return
}

func (d JsonMap) GetUintSlice(key string, def []uint) (val []uint, found bool, err error) {
	raw, found, err := d.GetAnySlice(key, []interface{}{}, uint(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint, 0)
	for _, i := range raw {
		val = append(val, i.(uint))
	}
	return
}

//// recursively get slice, and specialization for string/bool/float64/float32/int64/uint64/int32/uint32/int/uint

func (d JsonMap) RGetSlice(keyPath []string, def []interface{}) (val []interface{}, found bool, err error) {
	raw, found, err := d.RGetAny(keyPath, def)
	val = raw.([]interface{})
	return
}

func (d JsonMap) RGetAnySlice(keyPath []string, def []interface{}, itemType interface{}) (
	val []interface{}, found bool, err error) {
	raw, found, err := d.RGetSlice(keyPath, def)
	if !found || err != nil {
		return def, found, err
	}
	if len(raw) == 0 {
		return raw, found, err
	}
	ret := make([]interface{}, 0)
	for idx, i := range raw {
		v, e := toAny(i, itemType)
		if e != nil {
			err = fmt.Errorf("keyPath %s index %d: %v", keyPath, idx, e)
			return def, true, err
		}
		ret = append(ret, v)
	}
	return ret, true, nil
}

func (d JsonMap) RGetStringSlice(keyPath []string, def []string) (val []string, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, string(""))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]string, 0)
	for _, i := range raw {
		val = append(val, i.(string))
	}
	return
}

func (d JsonMap) RGetBoolSlice(keyPath []string, def []bool) (val []bool, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, bool(false))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]bool, 0)
	for _, i := range raw {
		val = append(val, i.(bool))
	}
	return
}

func (d JsonMap) RGetFloat64Slice(keyPath []string, def []float64) (val []float64, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, float64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]float64, 0)
	for _, i := range raw {
		val = append(val, i.(float64))
	}
	return
}

func (d JsonMap) RGetFloat32Slice(keyPath []string, def []float32) (val []float32, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, float32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]float32, 0)
	for _, i := range raw {
		val = append(val, i.(float32))
	}
	return
}

func (d JsonMap) RGetInt64Slice(keyPath []string, def []int64) (val []int64, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, int64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int64, 0)
	for _, i := range raw {
		val = append(val, i.(int64))
	}
	return
}

func (d JsonMap) RGetUint64Slice(keyPath []string, def []uint64) (val []uint64, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, uint64(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint64, 0)
	for _, i := range raw {
		val = append(val, i.(uint64))
	}
	return
}

func (d JsonMap) RGetInt32Slice(keyPath []string, def []int32) (val []int32, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, int32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int32, 0)
	for _, i := range raw {
		val = append(val, i.(int32))
	}
	return
}

func (d JsonMap) RGetUint32Slice(keyPath []string, def []uint32) (val []uint32, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, uint32(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint32, 0)
	for _, i := range raw {
		val = append(val, i.(uint32))
	}
	return
}

func (d JsonMap) RGetIntSlice(keyPath []string, def []int) (val []int, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, int(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]int, 0)
	for _, i := range raw {
		val = append(val, i.(int))
	}
	return
}

func (d JsonMap) RGetUintSlice(keyPath []string, def []uint) (val []uint, found bool, err error) {
	raw, found, err := d.RGetAnySlice(keyPath, []interface{}{}, uint(0))
	if !found || err != nil {
		return def, found, err
	}
	val = make([]uint, 0)
	for _, i := range raw {
		val = append(val, i.(uint))
	}
	return
}

//// type conversion

// convert raw to same type as def, will return def if failed
// if def is number, only float64, float32, int64, uint64, int32, uint32, int, uint are supported
func toAny(raw, def interface{}) (val interface{}, err error) {
	if def == nil && raw == nil {
		return raw, nil
	}
	if raw != nil && def != nil {
		rtk := reflect.TypeOf(raw).Kind()
		dtk := reflect.TypeOf(def).Kind()
		if f, ok := raw.(float64); ok {
			// NOTICE: json unmarshal all numbers to float64 as default, maybe precision lost
			if dtk == reflect.Float64 {
				return raw, nil
			} else if dtk == reflect.Float32 {
				return float32(f), nil
			} else if dtk == reflect.Int64 {
				return int64(f), nil
			} else if dtk == reflect.Uint64 {
				return uint64(f), nil
			} else if dtk == reflect.Int32 {
				return int32(f), nil
			} else if dtk == reflect.Uint32 {
				return uint32(f), nil
			} else if dtk == reflect.Int {
				return int(f), nil
			} else if dtk == reflect.Uint {
				return uint(f), nil
			}
		} else if n, ok := raw.(json.Number); ok {
			// NOTICE: if err != nil, val is not def, use strconv.Parsexxx returned
			if dtk == reflect.Float64 {
				return n.Float64()
			} else if dtk == reflect.Float32 {
				v, e := strconv.ParseFloat(string(n), 32)
				return float32(v), e
			} else if dtk == reflect.Int64 {
				return n.Int64()
			} else if dtk == reflect.Uint64 {
				return strconv.ParseUint(string(n), 10, 64)
			} else if dtk == reflect.Int32 {
				v, e := strconv.ParseInt(string(n), 10, 32)
				return int32(v), e
			} else if dtk == reflect.Uint32 {
				v, e := strconv.ParseUint(string(n), 10, 32)
				return uint32(v), e
			} else if dtk == reflect.Int {
				v, e := strconv.ParseInt(string(n), 10, int(strconv.IntSize))
				return int(v), e
			} else if dtk == reflect.Uint {
				v, e := strconv.ParseUint(string(n), 10, int(strconv.IntSize))
				return uint(v), e
			}
		} else 	if rtk == dtk {
			return raw, nil
		}
	}
	return def, fmt.Errorf("type error: got %T but expected %T", raw, def)
}
