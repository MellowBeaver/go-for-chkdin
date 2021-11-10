package main

//Using a docker container of mysql setting up employee db and
//make an API to CRUD an employee table.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()
	r.GET("/employee", readEmployeeHandler)
	r.POST("/employee", createEmplyeeHandler)
	r.PUT("/employee", updateEmployeeHandler)
	r.DELETE("/employee/:id", deleteEmployeeHAndler)

	r.Run(":8000")
}

func readEmployeeHandler(c *gin.Context) {

	//password = sam123
	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	emp := employee{}

	id := c.Request.URL.Query().Get("id")

	log.Println(id)
	i, _ := strconv.Atoi(id)

	row := db.QueryRow("Select * FROM employee where id=?", i)

	//error in row

	err = row.Scan(&emp.Id, &emp.Name, &emp.City)

	if err != nil {
		log.Println(err)
		c.JSON(500, "Row does not exist")
		return
	}

	c.JSON(200, emp)

}

func createEmplyeeHandler(c *gin.Context) {

	//password = sam123
	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	// Getting data from POST request body
	decoder := json.NewDecoder(c.Request.Body)

	type body_struct struct {
		Name string
		City string
	}

	var one body_struct

	err = decoder.Decode(&one)
	if err != nil {
		log.Println(err)
		return
	}

	//Adding row to the table
	rowadd, err := db.Exec("Insert into employee (name,city) VALUES (?,?)", one.Name, one.City)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error in addition in table!")
		return
	}

	output, err := rowadd.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(500, "Conversion")
		return
	}
	c.JSON(200, fmt.Sprintf("Added = %v ", output))
}

func updateEmployeeHandler(c *gin.Context) {

	//password = sam123
	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	// Getting data from PUT request body
	decoder := json.NewDecoder(c.Request.Body)

	type body_struct struct {
		Id   int
		Name string
		City string
	}

	var one body_struct

	err = decoder.Decode(&one)
	if err != nil {
		log.Println(err)
		return
	}

	var query string
	var params = make([]interface{}, 3)

	//Updating existing row
	if one.Name == "" {
		query = "Update employee set city = ? where id = ?"
		params = []interface{}{one.City, one.Id}
	} else if one.City == "" {
		query = "Update employee set name = ? where id = ?"
		params = []interface{}{one.Name, one.Id}
	} else {
		query = "Update employee set name = ?, city= ? where id = ?"
		params = []interface{}{one.Name, one.City, one.Id}
	}
	_, err = db.Exec(query, params...)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error in updation in table!")
		return
	}

	c.JSON(200, fmt.Sprintf("Updated = %v ", one.Id))

}

func deleteEmployeeHAndler(c *gin.Context) {

	//password = sam123
	db, err := sql.Open("mysql", "root:sam123@tcp(localhost:2000)/test1")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	//Creating employee struct

	type employee struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		City string `json:"city"`
	}

	deleteId := c.Params.ByName("id")

	//Delete
	//DELETE FROM CUSTOMERS WHERE ID = 6;
	_, err = db.Exec("Delete from employee where id = ?", deleteId)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error in deletion from table!")
		return
	}

	c.JSON(200, fmt.Sprintf("Deleted = %v ", deleteId))

}
