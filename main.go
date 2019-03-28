
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)
var importStr = "import UIKit \nimport ObjectMapper\nclass  DeviceResponse: BaseResponse {\n"
var initStr   = "\n override func mapping(map: Map){\n \n super.mapping(map: map)\n"
func main() {



	server := echo.New()
	server.Static("/","20181207mobile")
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//server.POST("/swift", func(context echo.Context) error {
	//	return context.JSON(200,"ok")
	//})
	server.POST("/swift",handleSwift)
	server.Start(":8088")

}

func handleSwift(c echo.Context ) error{


	body, _ := ioutil.ReadAll(c.Request().Body)
	//println("laile=======",string(body))
	var dic map[string]interface{}
	json.Unmarshal(body,&dic)

	//println("dic============",dic["jsonString"])

	var dicmap = dic["jsonString"]
	var dataMap map[string]interface{}
	json.Unmarshal([]byte(dicmap.(string)),&dataMap)

	println("开始打印=========",dataMap)

	//var dicnew = dicmap.(map[string]interface{})
	v,m := print_json(dataMap)
	fmt.Println("v===",v)
	fmt.Println("m===",m)
	saveFile(v,m)
	return c.JSON(200,"ok")
}
func print_json(m map[string]interface{}) (string,string){

	var valias = ""
	var mapings = ""

	//类型为字符串
	var strArry  =  []string{} //make([]string)
	//类型为float
	var floatArry  = make([]string,0)
	//类型为Int
	var intArry  = make([]string,0)
	//类型为bool
	var boolArry  = make([]string,0)
	//类型为数组对象
	var objArry  = make([]string,0)
	//类型为数组对象
	var objMapArry  = make([]string,0)

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
			strArry = append(strArry, k)

		case float64:
			fmt.Println(k, "is float", int64(vv))
			floatArry = append(floatArry, k)

		case int:
			fmt.Println(k, "is int", vv)
			intArry = append(intArry, k)

		case bool:
			fmt.Println(k, "is bool", vv)
			boolArry = append(boolArry, k)

		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
			objArry = append(objArry, k)
		case nil:
			fmt.Println(k, "is nil", "null")

		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			//print_json(vv)
			objMapArry = append(objMapArry, k)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
	fmt.Println("sss==",strArry)

	//saveFile(appendStringVar(strArry,"string"))
	valias = valias + appendStringVar(strArry,"String")
	fmt.Println("strArry===",strArry)
	mapings = mapings + appendStringMapping(strArry)

	valias = valias + appendStringVar(floatArry,"Double")
	mapings = mapings + appendStringMapping(floatArry)
	fmt.Println("floatArry===",floatArry)

	valias = valias + appendStringVar(intArry,"Int")
	mapings = mapings + appendStringMapping(intArry)
	fmt.Println("intArry===",intArry)

	valias = valias + appendStringVar(boolArry,"Bool")
	mapings = mapings + appendStringMapping(boolArry)
	fmt.Println("boolArry===",boolArry)

	fmt.Println("objArry===",objArry)
	fmt.Println("objMapArry===",objMapArry)
	return valias,mapings
}
func saveFile(st string ,st2 string)  {

	fmt.Println("time===",time.Now().Unix())
	//base 10进制
	times := strconv.FormatInt(time.Now().Unix(),10)
	var file = "./" + times + ".swift"
	f, _ := os.Create(file) //创建文件

	//defer f.Close()
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, _ := w.WriteString(importStr)
	fmt.Printf("写入 %d 个字节n", n4)

	w.WriteString(st)
	w.WriteString(initStr)
	w.WriteString(st2)

	//f.Close()
	var going = func(){
		w.WriteString("\n  }\n }")
		w.Flush()
		f.Close()
	}
	defer  going()

}

func appendStringVar(arry []string ,typeS string) string {
	var s = ""
	for str := range arry{
		s = s + "\n var " + arry[str] + ":" + typeS+ "?\n"
	}
	return s
}
func appendStringMapping(arry []string ) string {
	var s = ""
	for str := range arry{
		s = s + "\n self. " + arry[str] + " <- " +  "map[" + arry[str] + "]\n"
	}
	return s
}