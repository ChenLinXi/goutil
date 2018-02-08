package json

import (
	"os"
	"fmt"

	"io/ioutil"
	"encoding/json"
)

//将json config文件转换成json对象
func Parse(filename string, structure interface{}) error {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		return fmt.Errorf("加载 %s 配置文件出错", file)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("加载 %s 配置文件出错", filename)
	}
	return json.Unmarshal(data, &structure)
}

//将json 字符串转换成json对象
func ParseJsonString(jsonstr string, structure interface{}) error {
	return json.Unmarshal([]byte(jsonstr), &structure)
}

//将json byte转换成json对象
func ParseJsonByte(jsonstr []byte, structure interface{}) error {
	return json.Unmarshal(jsonstr, &structure)
}
