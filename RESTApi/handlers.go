package main

import (
	"crypto/sha256"
	"document"
	"encoding/hex"
	"encoding/json"
	"handler"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/message/{digest}", handlerGetDocByID).Methods("GET")
	muxRouter.HandleFunc("/message", handlerPostDoc).Methods("POST")
	return muxRouter
}

//Retrieve only document matching query
func handlerGetDocByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	doc, err := product.FindByValue(params["digest"])
	if err != nil {
		if err == mgo.ErrNotFound {
			handler.RespondWithError(w, http.StatusNotFound, "Message not found")
			return
		}
		handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handler.RespondWithJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: doc.Message})
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

	doc.Digest = encrypt(doc.Message)
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
	handler.RespondWithJSON(w, http.StatusCreated, struct {
		Digest string `json:"digest"`
	}{Digest: doc.Digest})
}

func encrypt(message string) string {
	hash := sha256.New()
	hash.Write([]byte(message))
	digest := hex.EncodeToString(hash.Sum(nil))
	return digest
}
