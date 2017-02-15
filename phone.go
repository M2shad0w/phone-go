package m2phone

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)

var (
	std              *PhoneRecord
	buf              []byte
	errInvailidPhone = errors.New("invalid phone format")
)

type PhoneRecord struct {
	Phone     string
	Province  string
	City      string
	Zipcode   string
	Areacode  string
	Phonetype byte
}

func Init(dataFile string) (err error) {
	if std != nil {
		return
	}
	std, err = NewLocator(dataFile)
	return
}

func NewLocator(dataFile string) (loc *PhoneRecord, err error) {
	buf, err = ioutil.ReadFile(dataFile)
	if err != nil {
		return
	}
	loc = new(PhoneRecord)
	return
}

func (p *PhoneRecord) Humanphonetype() string {

	switch int(p.Phonetype) {
	case 1:
		return "移动"
	case 2:
		return "联通"
	case 3:
		return "电信"
	case 4:
		return "电信虚拟运营商"
	case 5:
		return "联通虚拟运营商"
	case 6:
		return "移动虚拟运营商"
	default:
		return "未知运营商"
	}
}

func (p *PhoneRecord) Formatresult(b []byte) {
	s := []byte{'|'}
	attr_arr := bytes.Split(b, s)
	p.Province = string(attr_arr[0])
	p.City = string(attr_arr[1])
	p.Zipcode = string(attr_arr[2])
	p.Areacode = string(attr_arr[3])
}

func (p *PhoneRecord) Humanphoneinfo() {
	fmt.Println("手机号码:", p.Phone)
	fmt.Println("城市:", p.Province, p.City)
	fmt.Println("邮编:", p.Zipcode)
	fmt.Println("区号:", p.Areacode)
	fmt.Println("卡类型:", p.Humanphonetype())
}

func Find(phonestr string) (*PhoneRecord, error) {
	return std.Find(phonestr)
}

func (p *PhoneRecord) Find(phonestr string) (*PhoneRecord, error) {
	if len(phonestr) < 7 {
		err := errInvailidPhone
		return p, err
	}

	phone, err := strconv.Atoi(phonestr[0:7])
	//	version := string(buf[0:4])
	first_index_offset := int(binary.LittleEndian.Uint32(buf[4:8]))

	//	fmt.Println("version:", version, "first_index_offset", first_index_offset)

	recode_length := 9
	phone_record_count := (len(buf) - first_index_offset) / recode_length
	//	fmt.Println("记录条数:", phone_record_count)
	left := 0
	right := phone_record_count
	var middle, cur_offset, cur_phone int
	for left <= right {
		middle = (left + right) / 2
		cur_offset = first_index_offset + middle*recode_length
		if cur_offset > len(buf) {
			return p, err
		}
		//		fmt.Println("right", right, "left", left)
		cur_phone = int(binary.LittleEndian.Uint32(buf[cur_offset : cur_offset+4]))
		if cur_phone > phone {
			right = middle - 1
		} else if cur_phone < phone {
			left = middle + 1
		} else {
			record_offset := int(binary.LittleEndian.Uint32(buf[cur_offset+4 : cur_offset+8]))
			m := []byte{0}
			end_offset := record_offset + bytes.Index(buf[record_offset:], m)
			p.Phonetype = buf[cur_offset+8]
			p.Formatresult(buf[record_offset:end_offset])
			break
		}

	}
	return p, err
}
