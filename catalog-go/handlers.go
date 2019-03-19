package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
)

func HomePage(w http.ResponseWriter, r *http.Request){

	template := template.Must(template.ParseFiles("templates/homepage.html"))
	  
    err := template.Execute(w, nil) //execute the template
    if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
		http.Error(w, err.Error(), http.StatusInternalServerError)
  	}
}

func GetProducts(w http.ResponseWriter, r *http.Request){

	products := Products{
		Product{ ItemId: "", Name: "Goku", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Vegeta", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Gohan", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Piccolo", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Kame-Sennin", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Bulma", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Freeza", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Cell", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Broly", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "MajinBoo", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Kuririn", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Tenshinhan", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Gohan", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Trunks", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Android18", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Goten", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Beerus", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Whis", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "RedRibbonArmy", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "GeneralBlue", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "DoctorGero", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Taopaipai", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Raditz", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Nappa", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Bobbidi", Description: "Description", Price: 12.2},
		Product{ ItemId: "", Name: "Dabra", Description: "Description", Price: 12.2},
    }
 
    if err := json.NewEncoder(w).Encode(products); err != nil {
        panic(err)
    }
}