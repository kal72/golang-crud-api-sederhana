package usercontroller

import (
	"net/http"
	"io/ioutil"
	"golang-crud-api-sederhana/model"
	"encoding/json"
	"golang-crud-api-sederhana/service/userservice"
	"strconv"
	"fmt"
)

/** example
*(POST)(body: user) http://localhost:8080/users
*(GET) http://localhost:8080/users
*(GET) http://localhost:8080/users?id=1
*(DELETE) http://localhost:8080/users?id=1
*
 */
func Handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		b, e := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}

		var user model.User
		json.Unmarshal(b, &user)

		data := userservice.Save(&user)

		var result, err = json.Marshal(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	} else if r.Method == "GET" {
		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			data := userservice.Find()
			var result, err = json.Marshal(&data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			return
		}

		i,errInt := strconv.Atoi(keys[0])
		if errInt != nil {
			fmt.Print(errInt)
			http.Error(w,"", http.StatusInternalServerError)
			return
		}

		data := userservice.FindById(i)
		var result, err = json.Marshal(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return

	} else if r.Method == "DELETE" {
		keys, ok := r.URL.Query()["id"]

		if !ok || len(keys[0]) < 1 {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		i,errInt := strconv.Atoi(keys[0])
		if errInt != nil {
			fmt.Print(errInt)
			http.Error(w,"", http.StatusInternalServerError)
			return
		}

		message := userservice.Delete(i)
		var data = map[string]string{"message": message}
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
