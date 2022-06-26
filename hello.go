// JSONデータ処理

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	
)

// Mydata is json structure
type Mydata struct {
	Name string
	Mail string
	Tel string
}

// Str get string value
func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + ", " + m.Tel + ">"
}

// FirebaseからJSONデータを取得する
func main(){
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	// reにResponseが格納される。
	re, er := http.Get(p)

	if er != nil {
		panic(er)
	}
	defer re.Body.Close()
	
	// re.Bodyから全てのコンテンツを変数に格納
	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}
	
	var items []Mydata
	er = json.Unmarshal(s, &items)
	if er != nil {
		panic(er)
	}

	for i, im := range items {
		println(i, im.Str())
	}

}