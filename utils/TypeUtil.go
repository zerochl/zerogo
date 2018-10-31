package utils

import "reflect"

//类型检测 要检测的变量  期望变量类型
func TypeCheck(params interface{}, t string) bool {
	//数据初始化
	var (
		return_val bool = false
	)
	v := reflect.ValueOf(params)
	//获取传递参数类型
	v_t := v.Type()

	//类型名称对比
	if v_t.String() == t {
		return_val = true
	}
	return return_val
}
