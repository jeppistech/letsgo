package datastore

import (
	"github.com/boltdb/bolt"

	"../api/weather"

	"encoding/json"

	"log"
)


type WeatherDB struct {
	boltdb	*bolt.DB
}

func (db WeatherDB) GetAll()([]weather.WeatherPoint) {
	var points []weather.WeatherPoint
	err := db.boltdb.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte(weatherBucket))

	    b.ForEach(func(k, v []byte) error {
	    	var point weather.WeatherPoint
	        err := json.Unmarshal(v, &point)
	        if err != nil {
	        	log.Fatalf("Couldn't convert from json to a go object: ", err)
	        }

	        return nil
	    })

	    return nil
	})
}