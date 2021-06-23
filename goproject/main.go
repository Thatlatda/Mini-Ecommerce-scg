package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Person struct {
	Admin_id int    `json:"admin_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type Product struct {
	Product_id   int     `json:"product_id"`
	Product_name string  `json:"product_name"`
	Price        float32 `json:"price"`
	Descriptions string  `json:"descriptions"`
	Image        string  `json:"image"`
	Categories   string  `json:"categories"`
	Stock        int     `json:"stock"`
}

type Order struct {
	Order_id    int `json:"order_id"`
	Address     int `json:"address"`
	Quantity    int `json:"quantity"`
	Status      int `json:"status"`
	Customer_id int `json:"customer_id"`
}
type Order_Details struct {
	Id         int `json:"id"`
	Order_id   int `json:"order_id"`
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
	Price      int `json:"price"`
}

type Customers struct {
	Customer_id   int    `json:"customer_id"`
	Customer_name string `json:"customer_name"`
	Address       string `json:"address"`
	Phone         int    `json:"phone"`
	Email         string `json:"email"`
	Role          int    `json:"role"`
}

const (
	host     = "full_db_postgres"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "test"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

// Get admin_account
func GETHandler(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM admin_account")
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	for rows.Next() {
		var person Person
		rows.Scan(&person.Admin_id, &person.Name, &person.Username, &person.Password, &person.Role)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

// start products api
func GETHandlerProducts(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	var prod []Product

	for rows.Next() {
		var product Product
		rows.Scan(&product.Product_id, &product.Product_name, &product.Price, &product.Descriptions, &product.Image, &product.Categories, &product.Stock)
		prod = append(prod, product)
	}

	prodBytes, _ := json.MarshalIndent(prod, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(prodBytes)

	defer rows.Close()
	defer db.Close()
}

// Del product
func POSTHandlerDelProducts(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `DELETE FROM products WHERE product_id = $1`
	_, err = db.Exec(sqlStatement, p.Product_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

// Insert customers
func POSTHandlerInsertProducts(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	var a Product
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO products (product_id,product_name,price,descriptions,image,categories,stock) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err = db.Exec(sqlStatement, a.Product_id, a.Product_name, a.Price, a.Descriptions, a.Image, a.Categories, a.Stock)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

// Update products
func POSTHandlerUpdateProducts(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `UPDATE  products SET product_name = $1,price = $3,descriptions = $4,image = $5,categories = $6,stock = $6 WHERE product_id = $1`
	_, err = db.Exec(sqlStatement, p.Product_id, p.Product_name, p.Price, p.Descriptions, p.Image, p.Categories, p.Stock)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

// start orders api
func GETHandlerOrders(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}

	var orders []Order

	for rows.Next() {
		var order Order
		rows.Scan(&order.Order_id, &order.Customer_id, &order.Quantity, &order.Address, &order.Status)
		orders = append(orders, order)
	}

	ordersBytes, _ := json.MarshalIndent(orders, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(ordersBytes)

	defer rows.Close()
	defer db.Close()
}

// start customers api
func GETHandlerCustomers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		log.Fatal(err)
	}

	var cus []Customers

	for rows.Next() {
		var customers Customers
		rows.Scan(&customers.Customer_id, &customers.Customer_name, &customers.Address, &customers.Phone, &customers.Email, &customers.Role)
		cus = append(cus, customers)
	}

	cusBytes, _ := json.MarshalIndent(cus, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(cusBytes)

	defer rows.Close()
	defer db.Close()
}

// Insert customers
func POSTHandlerInsertCustomers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	db := OpenConnection()
	var a Customers
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO customers (customer_id,customer_name,address,phone,email,role VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err = db.Exec(sqlStatement, a.Customer_id, a.Customer_name, a.Address, a.Phone, a.Email, a.Role)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
func POSTHandlerProducts(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var p Product

	db := OpenConnection()

	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sqlStatement := `SELECT *
	FROM products WHERE product_id = $1;`
	rows, err := db.Query(sqlStatement, p.Product_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	var prod []Product

	for rows.Next() {
		var product Product
		rows.Scan(&product.Product_id, &product.Product_name, &product.Price, &product.Descriptions, &product.Image, &product.Categories, &product.Stock)
		prod = append(prod, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prod[0])
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func main() {
	http.HandleFunc("/getadmin", GETHandler)
	http.HandleFunc("/getproducts", GETHandlerProducts)
	http.HandleFunc("/insertproducts", POSTHandlerInsertProducts)
	http.HandleFunc("/delproducts", POSTHandlerDelProducts)
	http.HandleFunc("/updateproducts", POSTHandlerUpdateProducts)
	http.HandleFunc("/productdetails", POSTHandlerProducts)
	http.HandleFunc("/getorders", GETHandlerOrders)
	http.HandleFunc("/getcustomers", GETHandlerCustomers)
	http.HandleFunc("/insertcustomers", POSTHandlerInsertCustomers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
