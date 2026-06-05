package main

import (
	"fmt"
    "time"
	"net/http"
    "html/template"
    "net/url"
    "encoding/json"
    "os"
)
// Stałe globalne
const (
    author = "Antoni Kędra"
    port = "8080"
)
// Struktury dla odpowiedzi z Api geocoding i open-meteo
type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

type MeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Time        string  `json:"time"`
	} `json:"current_weather"`
}
// Struktura danych przekazywana do HTML
type TemplateData struct {
	Name        string
    Temperature float64
    Date        string
}
// Szablon HTML
var tmpl = template.Must(template.ParseFiles("index.html"))
// FUnkcja która obsługuje żądania http
func handler(w http.ResponseWriter, req *http.Request){
    data:= TemplateData{}
    if req.Method =="POST"{
        city:= req.FormValue("city") // Pobranie nazwy miasta z formularza
        fmt.Println(city) // Wydrukowanie nazwy do logów
        data.Name= city // Nazwa do szablonu
        plCity := url.QueryEscape(city) // Kodowanie do formatu URL
        //URL do API
        geoURL := fmt.Sprintf(
            "https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=pl&format=json", plCity)
        resG, err := http.Get(geoURL) // Żądanie GET
        if err != nil {
            fmt.Printf("error making http request: %s\n", err)
            return
        }
        defer resG.Body.Close()

        fmt.Printf("client: got response!\n")
        fmt.Printf("client: status code: %d\n", resG.StatusCode)
        
        var geoStruct GeoResponse // Dekodowanie do GeoResponse
        err = json.NewDecoder(resG.Body).Decode(&geoStruct)
        if err != nil {
            fmt.Printf("error %s\n", err)
            return
        }
        // Współrzędne miasta
        var lat, lon float64
        if len(geoStruct.Results)>0 {
            lat = geoStruct.Results[0].Latitude
            lon = geoStruct.Results[0].Longitude
        }else {
			tmpl.Execute(w, data)
			return
		}
        // URL do API
        meteoURL := fmt.Sprintf(
            "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)
        resM, err := http.Get(meteoURL) // Żądanie GET
        if err != nil {
            fmt.Printf("error making http request: %s\n", err)
            return
        }
        defer resM.Body.Close()

        fmt.Printf("client: got response!\n")
        fmt.Printf("client: status code: %d\n", resM.StatusCode)

        var meteoStruct MeteoResponse // Dekodowanie do MeteoResponse
        err = json.NewDecoder(resM.Body).Decode(&meteoStruct)
        if err != nil {
            fmt.Printf("error %s\n", err)
            return
        }
        // Zmienne do szablonu
        data.Temperature = meteoStruct.CurrentWeather.Temperature 
        data.Date = meteoStruct.CurrentWeather.Time
    }
    tmpl.Execute(w, data) //Renderowanie szablonu
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {

    // Obsługa healthchecka Dockera
    if len(os.Args) > 1 && os.Args[1] == "healthcheck" {
        resp, err := http.Get("http://localhost:8080/health")
        if err != nil || resp.StatusCode != http.StatusOK {
            os.Exit(1)
        }
        os.Exit(0)
    }

    // Logi
    fmt.Println("Start:", time.Now())
	fmt.Println("Autor:", author)
	fmt.Println("Port:", port)
    
    http.HandleFunc("/",handler)
    http.HandleFunc("/health", healthHandler)
    http.ListenAndServe(":"+port, nil)
}