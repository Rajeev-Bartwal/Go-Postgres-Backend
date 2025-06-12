package middleware

import "net/http"




type Response struct{
    Id int64 `json:"id,omitempty"`
    Meesage string `json:"message,omitempty"`
}



func GetStocks(w http.ResponseWriter ,r *http.Request){
    w.Header().Set("Content-type" ,"application/json")
	 
}

func GetAllStocks(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

}

func GetById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

}

func UpdateStock(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

}

func DeleteStock(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

}