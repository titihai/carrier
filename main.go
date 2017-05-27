package main

import (
	"fmt"

	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	t1 := time.Now()
	session, err := mgo.Dial("mongodb://zahai.pub:27017/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Println(time.Since(t1))

	t2 := time.Now()
	session.Clone()
	fmt.Println("Clone cost:", time.Since(t2))

	t3 := time.Now()
	session.Copy()
	fmt.Println("Copy cost:", time.Since(t3))

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("student")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t1))
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(t1))
	fmt.Println("Phone:", result.Phone)
}
