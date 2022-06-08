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
	//pp := strconv.FormatBool(isMutant(cdn.Dna))

	//fmt.Fprintf(w, pp)
	/*if !isMutant(cdn.Dna) {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println(pp)
			fmt.Fprintf(w, cdn.Dna[0])
		} else {
			//w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, cdn.Dna[2])
			//returnForbidden(w, r)
		}

	}*/
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

func returnForbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("403 HTTP status code returned!"))
}
