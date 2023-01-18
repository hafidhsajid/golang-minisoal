package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/deepaksinghvi/catalogwf/docs"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var baseURL = "http://localhost:5000"

const SESSION_ID = "id"

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type produk struct {
	KodeProduk int
	Nama       string
	Kuantitas  int
	Message    string
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /

// func dbConn() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "produk"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// func getCart() {
// 	var err error
// 	var client = &http.Client{}
// 	var data []cart

// 	request, err := http.NewRequest("POST", baseURL+"/cart", nil)

// 	response, err := client.Do(request)
// 	defer response.Body.Close()

// 	err = json.NewDecoder(response.Body).Decode(&data)
// 	if err != nil {
// 		println(err)
// 	}

// 	return cart, nil
// }

// func storeCart(w http.ResponseWriter, r *http.Request) {

// 	// w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	if r.Method == "POST" {
// 		KodeProduk := r.FormValue("kodeProduk")
// 		Nama := r.FormValue("namaProduk")
// 		Kuantitas := r.FormValue("kuantitas")
// 		fmt.Println(KodeProduk,
// 			Nama,
// 			Kuantitas)

// 		db := dbConn()

// 		insForm, err := db.Prepare("INSERT INTO cart (kodeProduk, namaProduk, kuantitas) VALUES(?,?,?)")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		insForm.Exec(KodeProduk, Nama, Kuantitas)

// 		defer db.Close()
// 	}

// }

// func getCart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// w.Header().Set("Content Type", "text/html")
// 	var result []byte
// 	db := dbConn()
// 	selDB, err := db.Query("SELECT * FROM cart ORDER BY kodeProduk DESC")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	data := produk{}
// 	res := []produk{}
// 	for selDB.Next() {
// 		var kodeProduk, kuantitas int
// 		var name string
// 		err = selDB.Scan(&kodeProduk, &name, &kuantitas)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		data.KodeProduk = kodeProduk
// 		data.Nama = name
// 		data.Kuantitas = kuantitas
// 		res = append(res, data)
// 	}
// 	// tmpl.ExecuteTemplate(w, "Index", res)
// 	result, err = json.Marshal(res)

// 	w.Write(result)
// 	// defer db.Close()

// 	// http.Error(w, "User not found", http.StatusNotFound)
// 	// return
// }

func addCart(c echo.Context) error {
	c.Request().Header.Set("Content Type", "text/html")
	// w.Header().Set("Content Type", "text/html")
	// 	html := `
	// 	<!DOCTYPE html>
	// <html>
	// <head>
	// <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
	// </head>
	// <body>
	// 	<form>
	//   <div class="form-group">GET
	//     <label for="exampleInputEmail1">Email address</label>
	//     <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
	//     <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
	//   </div>
	//   <div class="form-group">
	//     <label for="exampleInputPassword1">Password</label>
	//     <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
	//   </div>
	//   <div class="form-check">
	//     <input type="checkbox" class="form-check-input" id="exampleCheck1">
	//     <label class="form-check-label" for="exampleCheck1">Check me out</label>
	//   </div>
	//   <button type="submit" class="btn btn-primary">Submit</button>
	// </form>

	// </body>
	// </html>
	// 	`

	// w.Write([]byte(html))
	return c.File("./add.html")
}
func getall(c echo.Context) error {
	return c.File("./add.html")
}
func gatdata(c echo.Context) error {
	session, _ := store.Get(c.Request(), SESSION_ID)

	if len(session.Values) == 0 {
		// return echo.NewHTTPError(http.StatusOK, "empty result")

		return c.JSON(http.StatusOK, "")
	}
	var res []produk
	var finalres []produk
	jsonstring := fmt.Sprintf("%s", session.Values["data"])
	err := json.Unmarshal([]byte(jsonstring), &res)
	if err != nil {
		println(err)
	}
	nama := c.FormValue("nama")
	kuantitas, _ := strconv.Atoi(c.FormValue("kuantitas"))
	fmt.Println(c.FormValue("nama"), reflect.TypeOf(nama), strings.Contains(strings.ToLower(""), strings.ToLower(nama)))
	fmt.Println(c.FormValue("kuantitas"), reflect.TypeOf(kuantitas), c.FormValue("kuantitas") == "")

	for _, p := range res {
		if p.Kuantitas == kuantitas {
			finalres = append(finalres, p)
		}
		if strings.Contains(strings.ToLower(p.Nama), strings.ToLower(nama)) && c.FormValue("nama") != "" {
			finalres = append(finalres, p)
		}
		if c.FormValue("kuantitas") == "" && c.FormValue("nama") == "" {
			finalres = res
		}

	}

	return c.JSON(http.StatusOK, finalres)
}
func adddata(c echo.Context) error {

	session, _ := store.Get(c.Request(), SESSION_ID)
	var jsonstring string
	if session.Values["data"] != nil {
		jsonstring = fmt.Sprintf("%s", session.Values["data"])
		// session.Values["data"] = nil
		var res []produk
		data := produk{}
		kodeProduk := c.FormValue("kodeProduk")
		namaProduk := c.FormValue("namaProduk")
		kuantitas := c.FormValue("kuantitas")
		data.KodeProduk, _ = strconv.Atoi(kodeProduk)
		data.Nama = namaProduk
		data.Kuantitas, _ = strconv.Atoi(kuantitas)

		err := json.Unmarshal([]byte(jsonstring), &res)
		if err != nil {
			fmt.Println("Error unmarshalling")
		}
		var isInsert = true
		for i, p := range res {
			if p.KodeProduk == data.KodeProduk {
				p.Nama = data.Nama
				p.Kuantitas += data.Kuantitas
				// fmt.Println(p, i)
				res[i] = p
				isInsert = false

			}
		}
		if isInsert {
			res = append(res, data)
			fmt.Println("insert")

			result, _ := json.Marshal(res)
			session.Values["data"] = result
		} else {
			result, _ := json.Marshal(res)
			session.Values["data"] = result

		}

	} else {

		var res []produk
		data := produk{}
		kodeProduk := c.FormValue("kodeProduk")
		namaProduk := c.FormValue("namaProduk")
		kuantitas := c.FormValue("kuantitas")
		data.KodeProduk, _ = strconv.Atoi(kodeProduk)
		data.Nama = namaProduk
		data.Kuantitas, _ = strconv.Atoi(kuantitas)

		res = append(res, data)
		result, _ := json.Marshal(res)
		session.Values["data"] = result
		fmt.Print("ksoosng")
	}
	// var _ jsonData := []byte(jsonstring)
	// mes := session.Values["message1"]

	// fmt.Println(jsonstring)
	// session.Values["data"] = fmt.Sprintf("%s", data)
	// session.Values["message1"] = fmt.Sprintf("%s %s", mes, "asdf")
	session.Save(c.Request(), c.Response())

	// return c.Redirect(http.StatusTemporaryRedirect, "/get")
	return c.String(http.StatusCreated, fmt.Sprintf("%s", session.Values["data"]))
}
func deleteCart(c echo.Context) error {
	session, _ := store.Get(c.Request(), SESSION_ID)
	var jsonstring string
	kode, _ := strconv.Atoi(c.FormValue("kodeProduk"))
	fmt.Println(kode)
	if session.Values["data"] != nil {

		jsonstring = fmt.Sprintf("%s", session.Values["data"])
		var res []produk
		// data := produk{}
		err := json.Unmarshal([]byte(jsonstring), &res)
		if err != nil {
			fmt.Println("Error unmarshalling")
		}
		for i, p := range res {
			if p.KodeProduk == kode {
				// res[i] = res[i+1]
				res = append(res[0:i], res[i+1:]...)
			}
		}
		// fmt.Printf("res: %v\n", res)
		result, _ := json.Marshal(res)
		session.Values["data"] = result
	} else {
		jsonstring = ""
	}
	session.Save(c.Request(), c.Response())
	// return c.JSON(http.StatusOK,)

	return c.Redirect(http.StatusTemporaryRedirect, "/api/get")
	// return nil
}

func main() {
	// http.HandleFunc("/api/add", storeCart)
	// http.HandleFunc("/api", getCart)
	// http.HandleFunc("/", getCart)
	// http.HandleFunc("/delete", func(c echo.Context) error {
	// 	session, _ := store.Get(c.Request(), SESSION_ID)
	// 	session.Options.MaxAge = -1
	// 	session.Save(c.Request(), c.Response())

	// 	return c.Redirect(http.StatusTemporaryRedirect, "/get")
	// })
	// http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "add.html")
	// })GET

	// e.GET("/api/get", func(c echo.Context) error {
	// 	session, _ := store.Get(c.Request(), SESSION_ID)

	// 	if len(session.Values) == 0 {
	// 		// return echo.NewHTTPError(http.StatusOK, "empty result")

	// 		return c.JSON(http.StatusOK, "")
	// 	}
	// 	var res []produk
	// 	jsonstring := fmt.Sprintf("%s", session.Values["data"])
	// 	err := json.Unmarshal([]byte(jsonstring), &res)
	// 	if err != nil {
	// 		println(err)
	// 	}

	// 	return c.JSON(http.StatusOK, res)
	// 	// return c.JSON(http.StatusOK, fmt.Sprintf(
	// 	// 	"%s ",
	// 	// 	// session.Values["message2"],
	// 	// 	// session.Values["message2"],
	// 	// 	session.Values["data"],
	// 	// ))
	// })
	// e.POST("/api/add", func(c echo.Context) error {
	// 	session, _ := store.Get(c.Request(), SESSION_ID)
	// 	var jsonstring string
	// 	if session.Values["data"] != nil {
	// 		jsonstring = fmt.Sprintf("%s", session.Values["data"])
	// 		// session.Values["data"] = nil
	// 		var res []produk
	// 		data := produk{}
	// 		kodeProduk := c.FormValue("kodeProduk")
	// 		namaProduk := c.FormValue("namaProduk")
	// 		kuantitas := c.FormValue("kuantitas")
	// 		data.KodeProduk, _ = strconv.Atoi(kodeProduk)
	// 		data.Nama = namaProduk
	// 		data.Kuantitas, _ = strconv.Atoi(kuantitas)

	// 		err := json.Unmarshal([]byte(jsonstring), &res)
	// 		if err != nil {
	// 			fmt.Println("Error unmarshalling")
	// 		}
	// 		var isInsert = true
	// 		for i, p := range res {
	// 			if p.Nama == data.Nama {
	// 				p.KodeProduk = data.KodeProduk
	// 				p.Kuantitas += data.Kuantitas
	// 				// fmt.Println(p, i)
	// 				res[i] = p
	// 				isInsert = false

	// 			}
	// 		}
	// 		if isInsert {
	// 			res = append(res, data)
	// 			fmt.Println("insert")

	// 			result, _ := json.Marshal(res)
	// 			session.Values["data"] = result
	// 		} else {
	// 			result, _ := json.Marshal(res)
	// 			session.Values["data"] = result

	// 		}

	// 	} else {

	// 		var res []produk
	// 		data := produk{}

	// 		data.KodeProduk = 112
	// 		data.Nama = "name"
	// 		data.Kuantitas = 10

	// 		res = append(res, data)
	// 		result, _ := json.Marshal(res)
	// 		session.Values["data"] = result
	// 		fmt.Print("ksoosng")
	// 		fmt.Print(result)
	// 	}
	// 	// var _ jsonData := []byte(jsonstring)
	// 	// mes := session.Values["message1"]

	// 	// fmt.Println(jsonstring)
	// 	// session.Values["data"] = fmt.Sprintf("%s", data)
	// 	// session.Values["message1"] = fmt.Sprintf("%s %s", mes, "asdf")
	// 	session.Save(c.Request(), c.Response())

	// 	// return c.Redirect(http.StatusTemporaryRedirect, "/get")
	// 	return c.String(http.StatusOK, fmt.Sprintf("%s %s", jsonstring, session.Values["data"]))
	// })
	e := echo.New()
	// e.GET("/", )
	e.GET("/*", echoSwagger.WrapHandler)
	e.GET("/add", addCart)
	e.GET("/api/get", gatdata)
	e.POST("/api/add", adddata)
	e.GET("/api/delete", deleteCart)
	e.Start(":5000")
	// var carts, err = getCart()kuantitas

	// fmt.Println("starting web server at http://localhost:5000/")
	// http.ListenAndServe(":5001", nil)

}
