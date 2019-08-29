// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"reflect"

	"github.com/MDGSF/utils/log"
)

// CutMsgWithTemplate 根据 tpl 从 msg 中取出数据，放到 result 中
func CutMsgWithTemplate(tpl interface{}, msg interface{}) interface{} {

	switch tpl.(type) {
	case map[string]interface{}:
		mtpl := tpl.(map[string]interface{})

		mmsg, ok := msg.(map[string]interface{})
		if !ok {
			return nil
		}

		mresult := make(map[string]interface{})

		for strKey, ItValue := range mtpl {
			valuetype := reflect.TypeOf(ItValue)
			strValueType := valuetype.String()
			valuetypeName := valuetype.Name()
			log.Verbose("valuetype = %v, %v, %v", valuetype, strValueType, valuetypeName)

			if _, ok := mmsg[strKey]; !ok {
				log.Verbose("收到的数据中 key=[%v] 的数据不存在，跳过该字段", strKey)
				continue
			}

			switch ItValue.(type) {
			case float64:
				mresult[strKey] = mmsg[strKey]
			case map[string]interface{}:
				out := CutMsgWithTemplate(mtpl[strKey], mmsg[strKey])
				mresult[strKey] = out
				// log.Info("mresult = %v", mresult)
			case []interface{}:
				arrTpl := ItValue.([]interface{})
				if len(arrTpl) != 1 {
					log.Error("模板中的数组只能填写一个元素模板，当前数量是 %v 个。", len(arrTpl))
					continue
				}
				arrItemTpl := arrTpl[0]

				arrMsg, ok := mmsg[strKey].([]interface{})
				if !ok {
					log.Error("收到的数据中 key=[%v] 的类型不是数组", strKey)
					continue
				}

				// 这个注释掉，则会发送空的数组
				// if len(arrMsg) == 0 {
				// 	continue
				// }

				arrResult := make([]interface{}, len(arrMsg))
				for k := range arrMsg {
					arrResult[k] = CutMsgWithTemplate(arrItemTpl, arrMsg[k])
				}
				mresult[strKey] = arrResult

			default:
				log.Info("Unknown ItValue type")
			}
		}

		return mresult
	default:
		log.Info("Unknown tpl type")
	}

	return nil
}
