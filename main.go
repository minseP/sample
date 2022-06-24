package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

//Member -
type Member struct {
	phoneNo    string
	userName   string
	userNumber int
	userAddr   string
	birthDay   string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("DB Connect NOW...")
		db, err := sql.Open("mysql", "root:geoliner23@@tcp(localhost:3307)/academy")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// 하나의 Row를 갖는 SQL 쿼리 : QueryRow()
		var phoneNo, userName, userAddr, birthDay string
		var userNumber int
		err = db.QueryRow("select HDP_NO, CUST_NM, CUST_NO, CUST_ADDR, BIRTH_DY from cust_info where cust_no = '1000000101'").Scan(&phoneNo, &userName, &userNumber, &userAddr, &birthDay)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(phoneNo, userName, userNumber, userAddr, birthDay)
		// Go 데이타
		mem := Member{phoneNo, userName, userNumber, userAddr, birthDay}
		fmt.Println("Member:")
		fmt.Println("mem - Member:", mem)
		//fmt.Println(json.Marshal(mem.phoneNo))
		jsonBytes, err := json.Marshal(mem)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonBytes)
		fmt.Println(err)
		// JSON 바이트를 문자열로 변경
		jsonString := string(jsonBytes)
		// 콘솔 로그
		fmt.Println(jsonString)
		// 출력
		w.Write([]byte(jsonString))
	})
	http.ListenAndServe(":8888", nil)
}
