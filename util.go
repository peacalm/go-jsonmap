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
	"bytes"
	"encoding/json"
	"fmt"
)

func DeepCopyMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = DeepCopyMap(vm)
		} else {
			cp[k] = v
		}
	}
	return cp
}

// overwrite new value in `src` to `dst` if some key conflicts
func DeepMergeMap(dst, src map[string]interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst should not be nil")
	}
	if src == nil {
		return nil
	}
	deepMergeMap(dst, src)
	return nil
}

func deepMergeMap(dst, src map[string]interface{}) {
	for srcKey, srcValue := range src {
		srcValueMap, srcValueIsMap := srcValue.(map[string]interface{})
		if srcValueIsMap {
			if _, ok := dst[srcKey].(map[string]interface{}); !ok {
				dst[srcKey] = make(map[string]interface{})
			}
			deepMergeMap(dst[srcKey].(map[string]interface{}), srcValueMap)
		} else {
			dst[srcKey] = srcValue
		}
	}
}

func JsonUnmarshalUseNumber(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buf)
	decoder.UseNumber()
	return decoder.Decode(v)
}

func Unmarshal(data []byte, useNumber bool) (jm JsonMap, err error) {
	m := make(map[string]interface{})
	if useNumber {
		err = JsonUnmarshalUseNumber(data, &m)
	} else {
		err = json.Unmarshal(data, &m)
	}
	return m, err
}
