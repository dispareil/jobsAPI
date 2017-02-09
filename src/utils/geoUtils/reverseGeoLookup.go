package geoUtils

import (
	"config"
	"log"
	"logger"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var loggerInstance *log.Logger
var googleMapsClient *maps.Client

func GetLocationFromGeoCode(lat float64, lon float64, locationMap map[string]string) {

	reverseGeoCodeRequest := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lon,
		},
	}

	reverseGeoCodeResponse, requestErr := googleMapsClient.ReverseGeocode(context.Background(), reverseGeoCodeRequest)

	if requestErr != nil {
		loggerInstance.Println(requestErr.Error())
	} else {

		addressComponentsList := reverseGeoCodeResponse[0].AddressComponents

		for _, addressComponent := range addressComponentsList {
			for _, addressType := range addressComponent.Types {
				locationMap[addressType] = addressComponent.LongName
			}
		}
	}
}

func init() {
	var googleMapsClientErr error

	loggerInstance = logger.Logger

	googleMapsClient, googleMapsClientErr = maps.NewClient(maps.WithAPIKey(config.GetConfig("googleGeoAPIKey")))

	if googleMapsClientErr != nil {
		loggerInstance.Panicln(googleMapsClientErr.Error())
	}
}
