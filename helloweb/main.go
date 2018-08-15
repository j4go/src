package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
	"html/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST则解析form data，但是解析不了application/json
	// GET or POST 192.168.197.46:9090?p1=aa&p2=bb&p1=cc
	/*
		输出
		get r.Form: map[p1:[aa cc] p2:[bb]]
		get path: /
		get param p1: [aa cc]
	*/
	//fmt.Printf("get r.Form: %s\n", r.Form)

	fmt.Printf("get path: %s\n", r.URL.Path)
	//fmt.Printf("get param p1: %s\n", r.Form["p1"])
	//fmt.Printf("get param p1: %s\n", r.FormValue("p1"))
	fmt.Printf("get param p1: %s\n", r.Form["p1"])
	// 通过 r.Form.Get() 来获取值，因为如果字段不存在，通过该方式获取的是空值。
	// 但是通过 r.Form.Get() 只能获取单个的值，如果是map的数组值，必须通过上面的方式来获取。
	fmt.Printf("get param p1: %s\n", r.Form.Get("p1"))

	/*
		Request本身也提供了FormValue()函数来获取用户提交的参数。
		如r.Form["username"]也可写成
		r.FormValue("username")。调用r.FormValue时会自动调用r.ParseForm，
		所以不必提前调用。r.FormValue只会返
		回同名参数中的第一个，若参数不存在则返回空字符串。
	*/

	// 获取json body
	/*
			请求body：
			{
				"p1": "dd",
				"p3": "hehe"
			}
			输出：
			get r.Form: map[]
			get path: /
			get param p1: []
			get json body: {
					"p1": "dd",
					"p3": "hehe"
		}

	*/
	//var data map[string][]string
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Printf("get json body: %s\n", string(body))
	//json.Unmarshal(body, &data)
	//for k, v := range data {
	//	fmt.Printf("%s : %s", k, strings.Join(v, " "))
	//}

	fmt.Fprintf(w, "Hello Yianny.") // 写入到w 返回

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("method:%s\n", r.Method)
	//r.ParseForm() // 更好的写法 使用FormValue
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Printf("username:%s\n", r.FormValue("username"))
		fmt.Printf("password:%s\n", r.FormValue("password"))

	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
