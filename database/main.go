package main

import (
	"encoding/json"

	"log"
	"strconv"

	"github.com/Cswapi/server/database/database"
	"github.com/Cswapi/server/database/swapi"
)

var dbName = "./database/Cswapi.db"

func main() {

	reSetDB := false
	if reSetDB {
		// initialize the database for the first time
		// if it is already exist , do not initialize it again
		/*database.Init(dbName)*/
		database.Start(dbName)
		reSetDatabase()
	} else {
		database.Start(dbName)
		//database.Update([]byte("users") , []byte("vip") , []byte("123456") )
		//example
		//1-7
		//log.Println( database.GetValue([]byte("films"), []byte("1")) )
		//87 some is invalid
		//log.Println( database.GetValue([]byte("people") , []byte("1")) )
		//61
		//log.Println( database.GetValue([]byte("planets") , []byte("1")) )
		//37
		//log.Println( database.GetValue([]byte("species") , []byte("1")) )
		//39 some is invalid
		//log.Println(database.GetValue([]byte("starships") , []byte("10")) )
		//37 some is invalid
		//log.Println(database.GetValue([]byte("vehicles") , []byte("10")) )
		//log.Println(database.CheckKeyExist([]byte("vehicles") , []byte("8")) )
		//log.Println( database.GetValue([]byte("users") , []byte("vip") ) )
		database.Stop()
	}
}

func reSetDatabase() {
	findFilm()
	findPerson()
	findPlanet()
	findSpecies()
	findVehicle()
	findStarship()
}

func findFilm() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Film(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("films"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("films"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("films/%d is already exit", index)
		}
	}
}

func findPerson() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Person(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("people"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("people"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("person/%d is already exit", index)
		}
	}
}

func findPlanet() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Planet(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("planets"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("planets"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("planets/%d is already exit", index)
		}
	}
}

func findSpecies() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Species(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("species"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("species"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("species/%d is already exit", index)
		}
	}
}

func findStarship() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Starship(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("starships"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("starships"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("starships/%d is already exit", index)
		}
	}
}

func findVehicle() {
	c := swapi.DefaultClient
	invalidTime := 0
	for index := 1; ; index++ {
		jsonStr := dump(c.Vehicle(index))
		indexStr := strconv.Itoa(index)
		if len(database.GetValue([]byte("vehicles"), []byte(indexStr))) == 0 {
			if !putIntoDb([]byte("vehicles"), []byte(indexStr), jsonStr) {
				invalidTime++
				if invalidTime == 10 {
					break
				}
			} else {
				invalidTime = 0
			}
		} else {
			log.Printf("vehicles/%d is already exit", index)
		}
	}
}

func dump(data interface{}, err error) []byte {
	jsonStr, err := json.MarshalIndent(data, "", "  ")
	return jsonStr
}

func putIntoDb(bucketName []byte, index []byte, jsonStr []byte) bool {
	stb := &swapi.Film{}
	err := json.Unmarshal(jsonStr, &stb)

	if err != nil {
		log.Fatal(err)
		return false
	} else if len(stb.URL) == 0 {
		log.Printf("%s/%s is invalid", bucketName, index)
		return false
	}
	log.Printf("solve %s/%s", bucketName, index)
	database.Update(bucketName, index, jsonStr)
	return true

}
