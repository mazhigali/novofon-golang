![GitHub last commit](https://img.shields.io/github/last-commit/mazhigali/novofon-golang?logo=github)
![Github Repository Size](https://img.shields.io/github/repo-size/mazhigali/novofon-golang?logo=github)
[![Github forks](https://img.shields.io/github/forks/mazhigali/novofon-golang?logo=github)](https://github.com/mazhigali/novofon-golang/network/members)
![Lines of code](https://img.shields.io/tokei/lines/github.com/mazhigali/novofon-golang?logo=github)
[![GitHub open issues](https://img.shields.io/github/issues/mazhigali/novofon-golang?logo=github)](https://github.com/mazhigali/novofon-golang/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/mazhigali/novofon-golang?logo=github)](https://github.com/mazhigali/novofon-golang/issues)

[![made-with-Go](https://img.shields.io/badge/novofon-Go-00aed8?logo=go)](https://pkg.go.dev/github.com/mazhigali/novofon-golang/novofon)
[![Golang go.mod Go version](https://img.shields.io/github/go-mod/go-version/mazhigali/novofon-golang?label=mod&logo=go)](https://pkg.go.dev/search?q=novofon)
[![GoReportCard](https://goreportcard.com/badge/github.com/mazhigali/novofon-golang)](https://goreportcard.com/report/github.com/mazhigali/novofon-golang)

[![GitHub Repo stars](https://img.shields.io/github/stars/mazhigali/novofon-golang?label=novofon-golang&logo=github&color=505050&logoColor=fff)](https://github.com/mazhigali/novofon-golang)
[![GitHub User's stars](https://img.shields.io/github/stars/mazhigali?label=mazhigali&logo=github&color=505050&logoColor=fff)](https://github.com/mazhigali)

![novofon Golang](https://raw.githubusercontent.com/mazhigali/novofon-golang/master/novofon_golang.jpg)

# novofon-golang
Library which help you work with API novofon (v1)


## Main file for next examples
``` go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	zApi "github.com/mazhigali/novofon-golang/v1"
)

func main() {

	if err := infoBalance(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := infoPrice(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := statistics(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := smsSend(); err != nil {
		fmt.Printf("%v\n", err)
	}

}

func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s\n", string(dataJSON))
}
```

## Example get balance

<a href="https://novofon.com/instructions/api/#api_info_balance_nvf" target="_parent">
https://novofon.com/instructions/api/#api_info_balance_nvf</a>

``` go
func infoBalance() error {

	balance := zApi.New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
	}

	data := []byte{}

	if err := balance.Request(&data); err != nil {
		return err
	}

	catchDataToStruct := struct {
		Status   string  `json:"status"`
		Balance  float32 `json:"balance"`
		Currency string  `json:"currency"`
		Message  string  `json:"message"`//if status == error
	}{}
	//or catchDataToStruct := zApi.CatchInfoBalance{},  if zApi.(struct) implement
    //not all ctructs implement !!!

	if err := json.Unmarshal(data, &catchDataToStruct); err != nil {
		return err
	}

	prettyPrint(catchDataToStruct)

	return nil
}
```

## Example Price

[https://novofon.com/instructions/api/#api_info_price_nvf](https://novofon.com/instructions/api/#api_info_price_nvf)
``` go
func infoPrice() error {

	priceKazakhstan := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=77270000000",
	}

	dataKaz := []byte{}

	if err := priceKazakhstan.Request(&dataKaz); err != nil {
		return err
	}

	kaz := zApi.CatchInfoPrice{}

	if err := json.Unmarshal(dataKaz, &kaz); err != nil {
		return err
	}
	prettyPrint(kaz)
	//---------------------------------------------------------------

	priceUnitedKingdom := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=380440000000", ///Ukraine Kyiv
		ParamsMap: map[string]string{ //higher priority above the ParamsString
			"number": "442040000000", //United Kingdom London
		},
	}

	dataKingdom := []byte{}

	if err := priceUnitedKingdom.Request(&dataKingdom); err != nil {
		return err
	}

	king := zApi.CatchInfoPrice{} //if zApi.(struct) implement
	if err := json.Unmarshal(dataKingdom, &king); err != nil {
		return err
	}
	prettyPrint(king)
	//---------------------------------------------------------------

	//Params priority
	//higher priority ParamsUrlValues
	//medium priority ParamsMap
	//low priority ParamsString

	//Kyiv or RusFed or London
	priceRusFed := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=380440000000", ///Ukraine Kyiv
		ParamsUrlValues: url.Values{ //higher priority above the ParamsMap
			"number": []string{"74952409106"}, //Moscow
		},
		ParamsMap: map[string]string{ //higher priority above the ParamsString
			"number": "442040000000", //United Kingdom London
		},
	}
	dataRusFed := []byte{}

	if err := priceRusFed.Request(&dataRusFed); err != nil {
		return err
	}

	resp := zApi.CatchInfoPrice{}
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	priceRusFed.ParamsUrlValues = url.Values{ //change number
		"number": []string{"73952407198"}, //Irkutsk
	}
	if err := priceRusFed.Request(&dataRusFed); err != nil { //try again
		return err
	}
	resp = zApi.CatchInfoPrice{}
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	priceRusFed.ParamsUrlValues = url.Values{ //change number
		"number": []string{"78612179165"}, ///Krasnodar
	}
	if err := priceRusFed.Request(&dataRusFed); err != nil { //try again
		return err
	}
	resp = zApi.CatchInfoPrice{} //if zApi.(struct) implement
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	return nil
}
```

## Get statistics

[https://novofon.com/instructions/api/#api_statistic_nvf](https://novofon.com/instructions/api/#api_statistic_nvf)
``` go
func statistics() error {

	type CatchStatistics struct {
		Status string `json:"status"`
		Start  string `json:"start"`
		End    string `json:"end"`
		Stats  []struct {
			Id          string  `json:"id"`
			Sip         string  `json:"sip"`
			Callstart   string  `json:"callstart"`
			From        int     `json:"from"`
			To          int     `json:"to"`
			Description string  `json:"description"`
			Disposition string  `json:"disposition"`
			Billseconds int     `json:"billseconds"`
			Cost        float32 `json:"cost"`
			Billcost    float32 `json:"billcost"`
			Currency    string  `json:"currency"`
			Seconds     int     `json:"seconds"`
		}
		Message string `json:"message"` //if status == error
	}

	statistics := zApi.New{
		APIMethod:    "/v1/statistics/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsUrlValues: url.Values{
			"start": []string{"2018-09-01 08:00:00"},
			"end":   []string{"2018-10-04 08:00:00"},
		},
		Timeout: 10000, //Miliseconds
	}

	data := []byte{}

	if err := statistics.Request(&data); err != nil {
		return err
	}

	catchDataToStruct := zApi.CatchStatistics{}
	//or
    catchDataToStruct := CatchStatistics{}, //if zApi.(struct) not implement

	if err := json.Unmarshal(data, &catchDataToStruct); err != nil {
		return err
	}

	prettyPrint(catchDataToStruct)

	return nil
}
```

## Get timezone

[https://novofon.com/instructions/api/#api_info_timezone](https://novofon.com/instructions/api/#api_info_timezone)
``` go
func infoTimezone() error {

    key   := "e30e16c201343883f77e"
	secret := "dbf5606ea4c1f2234201"
	
    timeZone := zApi.New{
		APIMethod:    "/v1/info/timezone/",
		APIUserKey:   key,
		APISecretKey: secret,
	}

	data := []byte{}

	if err := timeZone.Request(&data); err != nil {
		return err
	}
	fmt.Printf("%+v", string(data))
	return nil
}
```

```
