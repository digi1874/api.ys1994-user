/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 11:38:08
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-03 19:11:21
 */

package controllers

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var ordersRE = regexp.MustCompile(`^\[\["[0-9a-zA-Z_]+","(DESC|ASC)"\](,\["[0-9a-zA-Z_]+","(DESC|ASC)"\])*\]$`)

func listHandle(c *gin.Context, filter1, filter2, orders interface{}) (page, size int, err error) {
	page, _ = strconv.Atoi(c.Query("page"))
	size, _ = strconv.Atoi(c.Query("size"))
	ft      := c.Query("filter")
	ods     := c.Query("orders")

	if ft != "" {
		err = json.Unmarshal([]byte(ft), &filter1)

		if err == nil {
			err = json.Unmarshal([]byte(ft), &filter2)
		}
	}

	if err == nil && ods != "" {
		odsByte := []byte(ods)
		err = listValidatorOrder(filter1, odsByte)

		if err == nil {
			err = json.Unmarshal(odsByte, &orders)
		}
	}

	return page, size, err
}

func listValidatorOrder(table interface{}, ordersByte []byte) (err error) {
	if ordersRE.Match(ordersByte) == false {
		return errors.New("orders err")
	}

	var orders [][2]string
	err = json.Unmarshal(ordersByte, &orders)
	if err == nil {
		ftOf := reflect.TypeOf(table).Elem()
		for _, v := range orders {
			name := strings.ReplaceAll(v[0], "_", " ")
			name = strings.Title(name)
			name = strings.ReplaceAll(name, " ", "")
			if _, has := ftOf.FieldByName(name); has == false {
				err = errors.New("orders 没有 " + v[0])
				break
			}
		}
	}
	return err
}
