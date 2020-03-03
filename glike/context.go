package glike

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	responseWriter http.ResponseWriter
	request *http.Request
}

func (c *Context) JSON(code int,messages map[string]interface{}){
	data , err := json.MarshalIndent(messages,"","	")
	if err!=nil{
		fmt.Println(err)
		c.responseWriter.WriteHeader(400)
		return
	}
	c.responseWriter.WriteHeader(code)
	c.responseWriter.Write(data)
}

func (c *Context) String(code int,message string) {
	c.responseWriter.WriteHeader(code)
	c.responseWriter.Write([]byte(message))
}

