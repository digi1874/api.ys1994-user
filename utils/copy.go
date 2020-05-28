/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-22 10:18:28
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-24 21:51:02
 */

package utils


import (
	"encoding/json"
)

// Copy  b, _ := json.Marshal(v1) -> json.Unmarshal(b, &v2)
func Copy(v1 interface{}, v2 interface{}) error {
	b, err := json.Marshal(v1)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v2)
}
