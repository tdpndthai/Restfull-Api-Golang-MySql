package proc_api

import (
	"connectdb"
	"encoding/json"
	"entities"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		procs, err2 := procModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJSON(response, http.StatusOK, procs)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		procs, err2 := procModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJSON(response, http.StatusOK, procs)
		}
	}
}

func SearchPrices(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	smin := vars["min"]
	smax := vars["max"]
	min, _ := strconv.ParseFloat(smin, 64) //hàm chuyển đổi string -> float
	max, _ := strconv.ParseFloat(smax, 64)

	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		procs, err2 := procModel.SearchPrices(min, max)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err.Error())
		} else {
			respondWithJSON(response, http.StatusOK, procs)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var proc entities.Proc
	err := json.NewDecoder(request.Body).Decode(&proc)

	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		err2 := procModel.Create(&proc) 
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, proc)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var proc entities.Proc
	err := json.NewDecoder(request.Body).Decode(&proc)

	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		_,err2 := procModel.Update(&proc) 
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, proc)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	procId := vars["id"]
	id,_ := strconv.ParseInt(procId,10,64)
	db, err := connectdb.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		procModel := models.ProcModel{
			Db: db,
		}
		RowsAffected,err2 := procModel.Delete(id) 
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, map[string]int64{
				"Số dòng thay đổi":RowsAffected,
			})
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg}) //map có key và value là string,phần tử đầu key là error,value là msg
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
