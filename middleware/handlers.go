package middleware

import (
	"encoding/json"
	"go-postgres/connection"
	"go-postgres/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Response struct{
    Stock models.Stock `json:"stock,omitempty"`
    Message string `json:"message,omitempty"`
}

 var db = connection.Connection{}



func GetAllStocks(w http.ResponseWriter ,r *http.Request){
    
    con := db.CreateConection()
    defer con.Close()
    
    rows , err := con.Query("SELECt * FROM stocks")
    if err != nil{
        http.Error(w , err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var stocks []models.Stock
    
    for rows.Next(){
        var s models.Stock
        
        if err := rows.Scan(&s.StockId ,&s.Name, &s.Price, &s.Company); err != nil{
            http.Error(w , "Error Scanning Data" , http.StatusForbidden)
            return
        }

        stocks = append(stocks, s)    
    }
    
    w.Header().Set("Content-type" ,"application/json")
    w.WriteHeader(http.StatusFound)

    if err := json.NewEncoder(w).Encode(stocks); err != nil{
        http.Error(w , err.Error(),http.StatusBadRequest)
        return
    }
}

func CreateStock(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

    var payload models.Stock


    if err :=json.NewDecoder(r.Body).Decode(&payload); err != nil{
        http.Error(w, err.Error(),http.StatusBadRequest)
        return
    }
    
    con := db.CreateConection()
    defer con.Close()

    query := "INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3)"
    _, err := con.Exec(query , payload.Name , payload.Price , payload.Company)
    if err != nil{
        http.Error(w, "Failed To Insert Data" , http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    response := Response{
        Stock: payload,
        Message: "Stock created Successfully",
    }

    if err :=json.NewEncoder(w).Encode(response);err != nil{
        http.Error(w, "Something went Wrong" , http.StatusInternalServerError)
    }

}

func GetById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type" ,"application/json")

    params := mux.Vars(r)
    n , _ := strconv.Atoi(params["id"])

    var stock models.Stock

    con := db.CreateConection()
    query := "SELECt * FROM stocks where stockid = $1"
    defer con.Close()

        row := con.QueryRow(query, n)
        if err := row.Scan(&stock.StockId , &stock.Name,&stock.Price, &stock.Company); err != nil{
               http.Error(w, "Stock Not Found" , http.StatusBadRequest)
               return
        }

        if err := json.NewEncoder(w).Encode(stock); err != nil{
               http.Error(w,err.Error() , http.StatusInternalServerError)
               return
           }

        response := Response{
            Stock: stock,
            Message: "Found",
        }

        if err :=json.NewEncoder(w).Encode(response);err != nil{
            http.Error(w, "Something goes Wrong", http.StatusInternalServerError)
            return
        }
}

func UpdateStock(w http.ResponseWriter,r *http.Request){
        w.Header().Set("Content-type" ,"application/json")
        
        var payload models.Stock
        if err := json.NewDecoder(r.Body).Decode(&payload); err != nil{
            http.Error(w,err.Error(),http.StatusBadRequest)
            return
        }
        
        params := mux.Vars(r)
        n , _ := strconv.Atoi(params["id"])

        con := db.CreateConection()
        query := "UPDATE stocks SET name=$1, price=$2, company=$3 WHERE stockid = $4;"
        defer con.Close()
        
        _, err := con.Exec(query ,payload.Name , payload.Price, payload.Company, n)
        if err != nil{
            http.Error(w, "Failed to Update" , http.StatusInternalServerError)
            return
        }
        
        var stock models.Stock

        row := con.QueryRow("Select * from stocks where stockid = $1", n)
        if err := row.Scan(&stock.StockId , &stock.Name, &stock.Price, &stock.Company); err != nil{
            http.Error(w, "Stock Not Found" , http.StatusBadRequest)
            return
        }
        
        response := Response{
            Stock: stock,
            Message: "Stock Updated Successfully",
        }

        if err :=json.NewEncoder(w).Encode(response);err != nil{
            http.Error(w, "Something goes Wrong", http.StatusInternalServerError)
            return
        }


}

func DeleteStock(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Content-type" ,"application/json")
    
    params := mux.Vars(r)
    n , _ := strconv.Atoi(params["id"])

    con := db.CreateConection()
    query := "DELETE FROM stocks  WHERE stockid = $1;"
    defer con.Close()

    
    var st models.Stock

    row := con.QueryRow("Select * from stocks where stockid = $1", n)
    if err := row.Scan(&st.StockId , &st.Name, &st.Price, &st.Company); err != nil{
        http.Error(w, "Stock Not Found" , http.StatusBadRequest)
        return
    }

    _,err := con.Exec(query , n)
    if err != nil{
        http.Error(w,"Stock not found",http.StatusInternalServerError)
        return
    }

    Response := Response{
        Stock: st,
        Message: "Stock deleted SuccessFully",
    }
    json.NewEncoder(w).Encode(Response)


}