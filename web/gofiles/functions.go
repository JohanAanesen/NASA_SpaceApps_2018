package gofiles

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleMain(w http.ResponseWriter, r *http.Request){
	temp, err := template.ParseFiles("www/index.html")
	if err != nil {
		log.Fatal(err)
	}

	var data Launches

	data = getAllFromDB()

	temp.Execute(w, data)
}

func HandleAPI(w http.ResponseWriter, r *http.Request){

	//data := getApi()

	data := getAllFromDB()

	http.Header.Add(w.Header(), "content-type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func getApi()Launches{
	resp, err := http.Get("https://launchlibrary.net/1.4/launch/next/50")
	if err != nil {
		log.Fatal(err)
	}

	var data Launches

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Error with json decoder: %s", err)
	}

	return data
}

func HandleUpload(w http.ResponseWriter, r *http.Request){
	/*launches := getApi()
	for i := range launches.Launch {
		SaveData(launches.Launch[i])
	}*/

	//data := getLaunchFromDB(1069)

	//http.Header.Add(w.Header(), "content-type", "application/json")

	//json.NewEncoder(w).Encode(data)

	UpdateDatabase()
}

func UpdateDatabase(){
	var freshData Launches

	freshData = getApi()

	if len(freshData.Launch) > 1 {
		DeleteLaunchesFromDB()

		for i := range freshData.Launch {
			SaveData(freshData.Launch[i])
			fmt.Printf("Saved Launch with ID: %d\n", freshData.Launch[i].Id)
		}
	}
}