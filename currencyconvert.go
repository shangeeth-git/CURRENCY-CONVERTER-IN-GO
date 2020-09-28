package main
import	(
		"fmt"
		"net/http"
		"encoding/json"
		"html/template"
)
type jasonRes struct {
	Result             string `json:"result"`
	Documentation      string `json:"documentation"`
	TermsOfUse         string `json:"terms_of_use"`
	TimeLastUpdateUnix int    `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string `json:"time_last_update_utc"`
	TimeNextUpdateUnix int    `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string `json:"time_next_update_utc"`
	BaseCode           string `json:"base_code"`
	Summa			string `json:"summa"`
	ConversionRates    struct {
		USD float64     `json:"USD"`
		AED float64 `json:"AED"`
		ARS float64 `json:"ARS"`
		AUD float64 `json:"AUD"`
		BGN float64 `json:"BGN"`
		BRL float64 `json:"BRL"`
		BSD float64 `json:"BSD"`
		CAD float64 `json:"CAD"`
		CHF float64 `json:"CHF"`
		CLP float64 `json:"CLP"`
		CNY float64 `json:"CNY"`
		COP float64 `json:"COP"`
		CZK float64 `json:"CZK"`
		DKK float64 `json:"DKK"`
		DOP float64 `json:"DOP"`
		EGP float64 `json:"EGP"`
		EUR float64 `json:"EUR"`
		FJD float64 `json:"FJD"`
		GBP float64 `json:"GBP"`
		GTQ float64 `json:"GTQ"`
		HKD float64 `json:"HKD"`
		HRK float64 `json:"HRK"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		INR float64 `json:"INR"`
		ISK float64 `json:"ISK"`
		JPY float64 `json:"JPY"`
		KRW float64 `json:"KRW"`
		KZT float64 `json:"KZT"`
		MVR float64 `json:"MVR"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		NOK float64 `json:"NOK"`
		NZD float64 `json:"NZD"`
		PAB float64 `json:"PAB"`
		PEN float64 `json:"PEN"`
		PHP float64 `json:"PHP"`
		PKR float64 `json:"PKR"`
		PLN float64 `json:"PLN"`
		PYG float64 `json:"PYG"`
		RON float64 `json:"RON"`
		RUB float64 `json:"RUB"`
		SAR float64 `json:"SAR"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		THB float64 `json:"THB"`
		TRY float64 `json:"TRY"`
		TWD float64 `json:"TWD"`
		UAH float64 `json:"UAH"`
		UYU float64 `json:"UYU"`
		ZAR float64 `json:"ZAR"`
	} `json:"conversion_rates"`
}
func main(){
	disp:=func(w http.ResponseWriter, r *http.Request){

	var Data jasonRes

	endpoint := "https://v6.exchangerate-api.com/v6/7d65be461f7fee8d268703c2/latest/"
			
		location := r.FormValue("city")

		res := endpoint + location

		req,err:=http.Get(res)

		defer req.Body.Close()

	fmt.Println(location)

	if(err!=nil){

		fmt.Println(err)
	}
	decoder:=json.NewDecoder(req.Body)

	er1 := decoder.Decode(&Data)

	Data.Summa=r.FormValue("city2")

	if(er1!=nil){

		fmt.Println(er1)
	}
		wr,er := template.ParseFiles("result.html")

		if er != nil{
			print("Something went wrong")
		}else{
			wr.Execute(w,Data)
		}
	
}

	input:=func(w http.ResponseWriter, r *http.Request){

		wr,er := template.ParseFiles("input.html")

		if er != nil{

			fmt.Print("Something went wrong")

		}else{
			
			wr.Execute(w,nil)

			fmt.Print("redirected")
		}
	}
	http.HandleFunc("/input",input)
	http.HandleFunc("/temp",disp)
	http.ListenAndServe(":9005",nil)
}
