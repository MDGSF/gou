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
