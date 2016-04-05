
go build cmd && $GOPATH/bin/wtflocalserver

VCL
````
sub vcl_recv {

  set req.http.X-GEO-LATITUDE = geoip.latitude
  set req.http.X-GEO-LONGITUDE = geoip.longitude
  set req.http.X-GEO-CITY = geoip.city
  set req.http.X-GEO-CONTINENT-CODE = geoip.continent-code
  set req.http.X-GEO-COUNTRY-CODE = geoip.country-code
  set req.http.X-GEO-COUNTRY-CODE-3 = geoip.country-code-3
  set req.http.X-GEO-COUNTRY-NAME = geoip.country-name
  set req.http.X-GEO-POSTAL-CODE = geoip.postal-code
  set req.http.X-GEO-REGION = geoip.region
  set req.http.X-GEO-AREA-CODE = geoip.area-code
  set req.http.X-GEO-METRO-CODE = geoip.metro-code
}
````
