package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wasuken/todo-jsay/alert"
)

type AlertJsonResponse struct{
	Data []alert.IntervalAlert `json:"data"`
	Status int `json:"status"`
	Msg string `json:"msg"`
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./public/assets/")
	r.LoadHTMLGlob("public/*.html")
	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/api/alert", func(c *gin.Context) {
		title := c.PostForm("title")
		interval_second, err := strconv.Atoi(c.PostForm("interval"))
		count, e := strconv.Atoi(c.PostForm("count"))
		if err != nil {
			panic(err)
		}
		if e != nil {
			panic(e)
		}
		go func() {
			alt := alert.IntervalAlert{Title: title, Interval_second: interval_second, Count: count}
			alert.WriteAlertHist(alt)
			alert.AddAlert(alt)
		}()

	})
	r.GET("/api/alert", func(c *gin.Context){
		mp := *alert.GetAlertMap()
		rst := []alert.IntervalAlert{}
		for _, alt := range mp {
			rst = append(rst, alt)
		}
		data := AlertJsonResponse{
			Data: rst,
			Status: 200,
			Msg: "",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
