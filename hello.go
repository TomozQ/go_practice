// JSONデータ処理

package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	
)

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
	
	// 型を指定した変数を用意
	var data []interface{}
	
	er = json.Unmarshal(s, &data) // 第二引数にdataのポインタを渡す
	// -> 変数sに入れてあるJSONデータがGoの[]interface{}型の値に変換されdataに代入される。

	fmt.Println(data)
	// [
	// 		map[mail:zero@zero name:taro tel:000-000] 
	// 		map[mail:syoda@tuyano.com name:tuyano tel:999-999] 
	// 		map[mail:hanako@flower name:hanako tel:888-888] 
	// 		map[mail:sachiko@happy name:sachiko tel:777-777]
	// ]

	if er != nil {
		panic(er)
	}

	// dataにループを回す
	for i, im := range(data){ // dataから各データの値が変数imに取り出される。
		// 型アサーション ... 空のインターフェイス型の値を特定の方の値として取り出し直すのに使う機能
		// 変数 := 値.( 型 ) -> map[string]interface{} => キーをstring型、値をinterface型で指定。
		m := im.(	map[string]interface{} ) // 指定した型として値が変数に取り出される。
		fmt.Println(i, m["name"].(string), m["mail"].(string), m["tel"].(string)) // 出力も 値.( 型 ) としてstring型に変換している。
	}

}