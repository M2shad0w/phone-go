# phone-go

查询 phone 信息 go 版本

### 示例
```
package main

import (
	"../phone"
	"fmt"
	"strconv"
)

func init() {
	if err := m2phone.Init("../phone/phone.dat"); err != nil {
		panic(err)
	}
}

func main() {
	for i := 1329900; i < 1529999; i++ {
		phonestr := strconv.Itoa(i)
		fmt.Println(phonestr)
		ph, err := m2phone.Find(phonestr)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(ph.Province, ph.Phone)
	}
}

```
### 返回信息结构

```
type PhoneRecord struct {
	Phone     string
	Province  string
	City      string
	Zipcode   string
	Areacode  string
	Phonetype byte
}
```