# phone-go

查询 phone 信息 go 版本

### 效率
100W 次/s
```
./example  0.29s user 0.28s system 68% cpu 0.847 total 20w次时间
```


### 示例

`go get github.com/M2shad0w/phone-go`

```
package main

import (
	"fmt"
	"github.com/M2shad0w/phone-go"
	"strconv"
)

func init() {
	if err := m2phone.Init("../phone/phone.dat"); err != nil {
		panic(err)
	}
}

func main() {
	for i := 1329900; i < 1529900; i++ {
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