package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UnmarshalPerson lala
func UnmarshalPerson(data []byte) (Person, error) {
	var r Person
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal lala
func (p *Person) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// Person lal lallalal
type Person struct {
	NameSurname string    `json:"name_surname"`
	Phone       string    `json:"phone"`
	Date        string    `json:"date"`
	Money       string    `json:"money"`
	Currency    string    `json:"currency"`
	Account     string    `json:"account"`
	Type        PhoneType `json:"phone_type"`
	RusMessage  string    `json:"rus_message"`
	KgMessage   string    `json:"kg_message"`
	PersonID    string    `json:"person_id"`
}

// FormatData lal lallalal
func (p *Person) FormatData() {
	if len(p.Phone) > 3 {
		runes := []rune(p.Phone)
		if string(runes[0]) == "0" {
			p.Phone = p.Phone[1:len(p.Phone)]
		}
		p.Phone = strings.Replace(p.Phone, " ", "", -1)
		p.Phone = strings.Replace(p.Phone, "+", "", -1)
		p.Phone = strings.Replace(p.Phone, "-", "", -1)

		runes1 := []rune(p.Phone)
		if string(runes1[0:3]) != "996" {
			p.Phone = "996" + p.Phone
		}
	}
	p.Money = strings.Replace(p.Money, " ", "", -1)
}

// NurtelecomFormat lal lallalal
func (p *Person) NurtelecomFormat() {
	if len(p.Phone) > 3 {
		runes := []rune(p.Phone)
		if string(runes[0:3]) == "996" {
			p.Phone = "0" + p.Phone[3:len(p.Phone)]
		}
	}
}

// CreateMessage lal lallalal
func (p *Person) CreateMessage() {

	//fakeText := fmt.Sprintf("Уважаемый/ая %s! Дос-Кредобанк напоминает о погашении кредита до", p.NameSurname)
	//p.RusMessage = fakeText
	//p.KgMessage = ""

	rusText := fmt.Sprintf("Уважаемый/ая %s! Дос-Кредобанк напоминает о погашении кредита до %v, к оплате %.2f сом. Лицевой счет № %s. Погасить можно во всех отделениях Банка, терминалах, мобильных кошельках. Тел.: 8686.", p.NameSurname, p.DateValue().Format("2006-01-02"), p.FloatValue(), p.Account)
	p.RusMessage = rusText

	kgText := fmt.Sprintf("Урматтуу %s! Дос-Кредобанк эскертет: кредитти %v чейин төлөөңүз. Төлөө суммасы %.2f сом. Сиздин жеке эсеп № %s. Банктын бөлүмдөрү, терминалдары, мобилдик капчыктары аркылуу төлөө мүмкүн. Тел.: 8686", p.NameSurname, p.DateValue().Format("2006-01-02"), p.FloatValue(), p.Account)
	p.KgMessage = kgText
}

//FloatValue see format
func (p Person) FloatValue() float64 {
	if len(p.Money) > 0 {
		finalValue := strings.Replace(p.Money, ",", ".", -1)
		f, err := strconv.ParseFloat(finalValue, 64)
		if err != nil {
			return float64(0)
		}
		return f
	}
	return float64(0)
}

//DateValue see format https://yourbasic.org/golang/format-parse-string-time-date-example/
func (p Person) DateValue() time.Time {
	if len(p.Date) > 0 {
		layout := "02.01.2006 3:04:05" //"20.09.2018 0:00:00"
		t, err := time.Parse(layout, p.Date)
		if err != nil {
			fmt.Println(err)
		}
		return t
	}
	return time.Now()
}

//MakeUniqID l
func (p *Person) MakeUniqID() {
	uniqKey := time.Now().UnixNano() / int64(time.Nanosecond)
	uniqKeyStr := strconv.FormatInt(int64(uniqKey), 10)
	p.PersonID = uniqKeyStr
}
