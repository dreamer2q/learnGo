package main

import(
	"io/ioutil"
	"log"
//	"io"
	"net/http"
//	"html/template"
	"encoding/json"
	//"runtime/debug"
	"strings"
	"strconv"
)

type Location struct{
	province string 
	city string 
	longitude string 
	latitude string
}

const(
	VER = "1.0"
)

func getJsonData(longitude string, latitude string)[]byte{

	url := "https://api.caiyunapp.com/v2/xqPIGAu6p6zfF9Bv/"+longitude+","+latitude+"/forecast.json"
	resp, err := http.Get(url)
	checkErr(err)

	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	return ret
}

func getLocation(ip string)(Location,error){
	URL := "https://restapi.amap.com/v3/ip?ip="+ip+"&output=json&key=00b37778b04147ae2188029f3d69c38e"
	resp, err := http.Get(URL)
	checkErr(err)

	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var loc Location
	var r interface{}
	err = json.Unmarshal(ret, &r)
	checkErr(err)
	
	res, ok := r.(map[string]interface{})
	if ok {
		for k,v := range res{
			switch v2 := v.(type){
				case string:
					switch k {
						case "province":
							loc.province = v2
						case "city":
							loc.city = v2
						case "rectangle":
							pos := strings.Split(v2, ";")
							start := strings.Split(pos[0], ",")
							end := strings.Split(pos[1], ",")
							longitude,_ := strconv.ParseFloat(start[0], 32) 
							tmp,_ := strconv.ParseFloat(end[0], 32)
							longitude += tmp
							longitude /= 2.0
							latitude,_ := strconv.ParseFloat(start[1], 32) 
							tmp,_ = strconv.ParseFloat(end[1], 32)
							latitude += tmp
							latitude /= 2.0
							loc.longitude = strconv.FormatFloat(longitude,'f',6,32)
							loc.latitude = strconv.FormatFloat(latitude,'f',6,32)
					}
				default:
					log.Println("Meet unrecognized Json data")
			}
		}	
	}else{
		return loc,err
	}
	log.Println("getLocation:",loc)
	return loc,nil
}
func checkOk(ok bool){
	if ok {
		//log.Println("checkOk ",ok)
	}else{
		log.Panicln("checkOk ",ok)	
	}
}

func parseWeatherJsonData(data []byte)map[string]interface{}{
	var r interface{}
	err := json.Unmarshal(data, &r)
	checkErr(err)	

	var ret map[string]interface{}
	ret = make(map[string]interface{})
	res, ok := r.(map[string]interface{})
	if ok {
		if res["status"] == "ok" {

			log.Println("status is ok")

			//log.Println("location:",res["location"])
			location,ok := res["location"].([]interface{})
			checkOk(ok)
			ret["location"] = location[:]
			//log.Println("location assigned success")

			result,ok := res["result"].(map[string]interface{})
			checkOk(ok)
			ret["keypoint"] = result["forecast_keypoint"]
			minutely,ok := result["minutely"].(map[string]interface{})
			checkOk(ok)
			ret["mkeypoint"] = minutely["description"]
			hourly,ok := result["hourly"].(map[string]interface{})
			checkOk(ok)
			temperature,ok := hourly["temperature"].([]interface{})
			checkOk(ok)
			ret["temperature"] = temperature[0]
			pm25,ok := hourly["pm25"].([]interface{})
			checkOk(ok)
			ret["pm25"] = pm25[0]
		
		}else{
			ret["status"] = res["status"]
			ret["error"] = res["error"]
		}
	}
	return ret
}

/*
GET /weather?longitude=
			latitude=
			
121.6544,25.1552

test URL: https://api.caiyunapp.com/v2/xqPIGAu6p6zfF9Bv/121.6544,25.1552/realtime.json
*/

func weatherHandler(w http.ResponseWriter,r *http.Request){ //返回一个经过简化的API json数据
	log.Println("Get request:",r.RequestURI," Host:",r.RemoteAddr)
	//log.Println("   Host:",r.RemoteAddr)

	addr := strings.Split(r.RemoteAddr, ":")	
	if len(addr) != 2 {
		log.Panicln("Addr error:",addr)
	}
	remoteIP := addr[0];
	loc,err := getLocation(remoteIP)

	jsonData := getJsonData(loc.longitude, loc.latitude)
	//log.Println(string(jsonData))
	retData := parseWeatherJsonData(jsonData)
	retData["ip"] = remoteIP
	retData["version"] = VER
	retData["loc"] = make(map[string]interface{})
	location := retData["loc"].(map[string]interface{})
	location["province"] = loc.province
	location["city"] = loc.city

	retBytes,err := json.Marshal(retData)
	checkErr(err)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(retBytes)
	log.Println(string(retBytes))
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func main(){

	log.Println("Caiyun started...")
	//parseJsonData(getJsonData("121.6544", "25.1552"))

	http.HandleFunc("/weather", weatherHandler)

	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		log.Fatal("ListenAndServe:",err.Error())
	}


}
