//package main

//import (
//"fmt"
//"html/template"
//"net/http"
//)

//type NewsAggPage struct {
//Title string
//News  string
//}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//p := NewsAggPage{
//Title: "karllarkzs gaming",
//News:  "edi wow",
//}
//t, _ := template.ParseFiles("template.html")
//t.Execute(w, p)
//}
//func newAggHandler(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintf(w, "<h1>lol</h1>")
//}
//func main() {
//http.HandleFunc("/", indexHandler)
//http.HandleFunc("/agg/", newAggHandler)
//http.ListenAndServe(":8080", nil)
//}
//package main

//import (
//"encoding/json"
//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/labstack/echo"
//)

//Companies ...
//type Companies []struct {
//ID           int           `json:"id"`
//Title        string        `json:"title"`
//Description  string        `json:"description"`
//Founded      interface{}   `json:"founded"`
//IsVerified   interface{}   `json:"isVerified"`
//Email        string        `json:"email"`
//Telephone    string        `json:"telephone"`
//Website      string        `json:"website"`
//Revenue      string        `json:"revenue"`
//Size         string        `json:"size"`
//Headquarters string        `json:"headquarters"`
//Industry     string        `json:"industry"`
//Longitude    string        `json:"longitude"`
//Latitude     string        `json:"latitude"`
//Address      string        `json:"address"`
//Slug         string        `json:"slug"`
//IsClaimed    interface{}   `json:"isClaimed"`
//CreatedAt    time.Time     `json:"created_at"`
//UpdatedAt    time.Time     `json:"updated_at"`
//Addresses    []interface{} `json:"addresses"`
//Reviews      []struct {
//ID                 int         `json:"id"`
//Body               string      `json:"body"`
//Rating             interface{} `json:"rating"`
//JobTitle           string      `json:"jobTitle"`
//LastYear           interface{} `json:"lastYear"`
//NumberOfYears      interface{} `json:"numberOfYears"`
//ReviewTitle        string      `json:"reviewTitle"`
//Pros               string      `json:"pros"`
//Cons               string      `json:"cons"`
//AdviceToManagement string      `json:"adviceToManagement"`
//EmploymentStatus   string      `json:"employmentStatus"`
//Company            int         `json:"company"`
//User               int         `json:"user"`
//CreatedAt          time.Time   `json:"created_at"`
//UpdatedAt          time.Time   `json:"updated_at"`
//} `json:"reviews"`
//}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//p := Companies{}
//t, _ := template.ParseFiles("template.html")
//t.Execute(w, p)
//}

//func main() {
//fetch data from api
//response, err := http.Get("https://review.a-fis.net/companies")
//if err != nil {
//log.Println(err)
//return
//}

//var companies Companies

//if err := json.NewDecoder(response.Body).Decode(&companies); err != nil {
//log.Println(err)
//return
//}

//for _, company := range companies {
//log.Println(company.Title)
//}
//start server
//e := echo.New()

//e.GET("/", func(c echo.Context) error {
//var companies Companies

//if err := json.NewDecoder(response.Body).Decode(&companies); err != nil {
//log.Println(err)
//return
//}

//for _, company := range companies {
//log.Println(company.Title)
//}

//return c.JSON(http.StatusOK, companies)
//})

//e.Start(":8000")

//fmt.Println("Starting the application...")

//someTemplate["companies"] = companies
//}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

func main() {
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// initialize http client
	client := &http.Client{}

	// marshal User to json
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, "http://api.example.com/v1/user", bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
}
