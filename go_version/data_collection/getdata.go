package data_collection

import (
    "io/ioutil"
    "log"
    "net/http"
    "time"
    "strconv"
)




func GetData(contractID string, days string) []byte {
    const baseURL ="https://www.predictit.org/Home/GetChartPriceData?contractId=" 
    const suffix = "&days="
    res, err := http.Get(baseURL + contractID + suffix + days)
    if err != nil {
        log.Fatal(err)
    }
    data, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    return data
}

type PriceData struct{
    Date time.Time
    Price float64
}

type DataList []PriceData

func ParseData(data []byte) DataList {
    e := string(data)
    a := e[1:]
    var datalist DataList
    const frmt = "01/02/2006"
    for i := 0; i < len(a) - 1; i += 43 {
        d := a[i +9:i+19]
        dte, _ := time.Parse(frmt,d)
        price, _ := strconv.ParseFloat(string(a[i+38:i+41]), 2)
        datalist = append(datalist,PriceData{Date: dte, Price: price,})
}
    return datalist
}
