package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/text/language"
)

// This is a list of all languages that are supported by the service.
var matcher = language.NewMatcher([]language.Tag{
	language.AmericanEnglish,
	language.Afrikaans,
	language.Amharic,
	language.Arabic,
	language.ModernStandardArabic,
	language.Azerbaijani,
	language.Bulgarian,
	language.Bengali,
	language.Catalan,
	language.Czech,
	language.Danish,
	language.German,
	language.Greek,
	language.English,
	language.BritishEnglish,
	language.Spanish,
	language.EuropeanSpanish,
	language.LatinAmericanSpanish,
	language.Estonian,
	language.Persian,
	language.Finnish,
	language.Filipino,
	language.French,
	language.CanadianFrench,
	language.Gujarati,
	language.Hebrew,
	language.Hindi,
	language.Croatian,
	language.Hungarian,
	language.Armenian,
	language.Indonesian,
	language.Icelandic,
	language.Italian,
	language.Japanese,
	language.Georgian,
	language.Kazakh,
	language.Khmer,
	language.Kannada,
	language.Korean,
	language.Kirghiz,
	language.Lao,
	language.Lithuanian,
	language.Latvian,
	language.Macedonian,
	language.Malayalam,
	language.Mongolian,
	language.Marathi,
	language.Malay,
	language.Burmese,
	language.Nepali,
	language.Dutch,
	language.Norwegian,
	language.Punjabi,
	language.Polish,
	language.Portuguese,
	language.BrazilianPortuguese,
	language.EuropeanPortuguese,
	language.Romanian,
	language.Russian,
	language.Sinhala,
	language.Slovak,
	language.Slovenian,
	language.Albanian,
	language.Serbian,
	language.SerbianLatin,
	language.Swedish,
	language.Swahili,
	language.Tamil,
	language.Telugu,
	language.Thai,
	language.Turkish,
	language.Ukrainian,
	language.Urdu,
	language.Uzbek,
	language.Vietnamese,
	language.Chinese,
	language.SimplifiedChinese,
	language.TraditionalChinese,
	language.Zulu,
})

type geoIP struct {
	Language      string `json:"language"`
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
	t, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	// We ignore the error: the default language will be selected for t == nil.
	tag, _, _ := matcher.Match(t...)

	return geoIP{
		tag.String(),
		r.Header.Get("X-GEO-LATITUDE"),
		r.Header.Get("X-GEO-LONGITUDE"),
		r.Header.Get("X-GEO-CITY"),
		r.Header.Get("X-GEO-CONTINENT-CODE"),
		r.Header.Get("X-GEO-COUNTRY-CODE"),
		r.Header.Get("X-GEO-COUNTRY-CODE3"),
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
		for k, v := range r.Header {
			log.Println("key:", k, "value:", v)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1728000")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "public, max-age=0")
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprint(w, string(j))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
