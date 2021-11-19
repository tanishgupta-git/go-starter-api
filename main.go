package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	Coursename  string  `json:"coursename"`
	Courseprice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.Coursename == ""
}

func main() {

}

// controllers - file

// serve home route
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api by LearnCodeOnline</h1>"))

}

// getting all the courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
