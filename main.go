package main

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
