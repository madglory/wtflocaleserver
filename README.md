# WTF Locale Server

This is a small server, written in GO, that returns GEO location for use in
client-side apps.

At the moment, it relies on Fastly's GEO IP looking, but could someday
include something like MaxMind internally.

## Development

There are no external dependencies, so checkout the project into your $GOPATH and:

    go build cmd && PORT=8000 $GOPATH/bin/wtflocaleserver

## Production

I deploy this to Heroku on a Hobby dyno.  It can handle a few thousand requests
per second with a ~20ms response time.

In order to get the GEO information, we look for a series of headers set by Fastly.  You can ask
Fastly to set them using it's Headers feature or a custom VCL.

Here's a screenshot of setting up the headers feature:
![Fastly](/docs/fastly.png)

If you'd prefer a custom VCL, add this to the `vcl_recv` section.:
````
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
````
