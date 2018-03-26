package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetVersion returns latest app version for selected platform
func GetVersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := Database{}
	db.Initialize()
	v, ev := db.GetAppVersion(params["applicationType"])
	if ev != nil {
		json.NewEncoder(w).Encode(ev)
		return
	}

	f, ef := db.GetForceUpdate(params["applicationType"])
	if ef != nil {
		json.NewEncoder(w).Encode(ef)
		return
	}

	u, eu := db.GetURL(params["applicationType"])
	if eu != nil {
		json.NewEncoder(w).Encode(eu)
		return
	}

	version := Version{
		v,
		f,
		u,
	}
	json.NewEncoder(w).Encode(version)
}
