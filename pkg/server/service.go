package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *api) checkMutant(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var cdn Sequence
	json.Unmarshal(body, &cdn)

	mutantHandler(cdn.Dna, w, r)
}

func mutantHandler(sample []string, w http.ResponseWriter, r *http.Request) {

	if isMutant(sample) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "HTTP 200-OK")
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "HTTP 403-FORBIDDEN")
	}
}
