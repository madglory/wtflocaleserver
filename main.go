package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type geoIP struct {
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	City          string `json:"city"`
	ContinentCode string `json:"continentCode"`
	CountryCode   string `json:"countryCode"`
	CountryCode3  string `json:"countryCode3"`
	CountryName   string `json:"countryName"`
	PostalCode    string `json:"postalCode"`
	Region        string `json:"region"`
	AreaCode      string `json:"areaCode"`
	MetroCode     string `json:"metroCode"`
}

func geoIPFromRequest(r *http.Request) (g geoIP) {
	return geoIP{
		r.Header.Get("X-GEO-LATITUDE"),
		r.Header.Get("X-GEO-LONGITUDE"),
		r.Header.Get("X-GEO-CITY"),
		r.Header.Get("X-GEO-CONTINENT-CODE"),
		r.Header.Get("X-GEO-COUNTRY-CODE"),
		r.Header.Get("X-GEO-COUNTRY-CODE-3"),
		r.Header.Get("X-GEO-COUNTRY-NAME"),
		r.Header.Get("X-GEO-POSTAL-CODE"),
		r.Header.Get("X-GEO-REGION"),
		r.Header.Get("X-GEO-AREA-CODE"),
		r.Header.Get("X-GEO-METRO-CODE"),
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		g := geoIPFromRequest(r)
		j, err := json.Marshal(g)
		if err != nil {
			log.Println(err)
			return
		}
		log.Print(string(j))

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1728000")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "public, max-age=900, stale-while-revalidate=604800, stale-if-error=604800")

		fmt.Fprint(w, string(j))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
