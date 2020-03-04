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

	handlers HandlersChain
	index int
}

func (c *Context) reset() {
	c.handlers = nil
	c.index = 0
}

func (c *Context) Next() {
	// handlers not allowed dynamic change
	// index indicates next handler
	for c.index < len(c.handlers) {
		h := c.handlers[c.index]

		// inc before handle, may use in handler
		c.index++
		h(c)
	}
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

