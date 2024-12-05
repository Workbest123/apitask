package main

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
)

type ValCurs struct {
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}
type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func fetchRates(date string) (*ValCurs, error) {
	url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily_eng.asp?date_req=%s", date)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("error unmarshalling XML: %v", err)
	}

	return &valCurs, nil
}
func main() {
	date := "02/01/2006"

	fmt.Println(fetchRates(date))

}
