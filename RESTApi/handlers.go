package main

import (
	"document"
	"encoding/json"
	"handler"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/message", handlerGetDocByID).Methods("GET")
	muxRouter.HandleFunc("/message", handlerPostDoc).Methods("POST")
	return muxRouter
}

//Retrieve only document matching query
func handlerGetDocByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	doc, err := product.FindByValue(query.Get("message"))
	if err != nil {
		handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondWithJSON(w, http.StatusOK, doc)
}

//Post document to database
func handlerPostDoc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var doc document.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doc); err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	doc.ID = bson.NewObjectId()
	err := product.Insert(doc)
	switch {
	case mgo.IsDup(err):
		handler.RespondWithError(w, http.StatusConflict, err.Error())
		return
	case err != nil:
		handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	handler.RespondWithJSON(w, http.StatusCreated, doc)
}
