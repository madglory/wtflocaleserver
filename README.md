
go build cmd && $GOPATH/bin/wtflocalserver

VCL
````
sub vcl_recv {

  set req.http.X-GEO-LATITUDE = geoip.latitude
  set req.http.X-GEO-LONGITUDE = geoip.longitude
  set req.http.X-GEO-CITY = geoip.city
  set req.http.X-GEO-CONTINENT-CODE = geoip.continent_code
  set req.http.X-GEO-COUNTRY-CODE =   geoip.country_code
  set req.http.X-GEO-COUNTRY-CODE3 = geoip.country_code3
  set req.http.X-GEO-COUNTRY-NAME = geoip.country_name
  set req.http.X-GEO-POSTAL-CODE = geoip.postal_code
  set req.http.X-GEO-REGION = geoip.region
  set req.http.X-GEO-AREA-CODE = geoip.area_code
  set req.http.X-GEO-METRO-CODE = geoip.metro_code
}
````
