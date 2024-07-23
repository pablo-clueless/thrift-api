package main

import "thrift.com/m/api"

// @title Thrift
// @version 1.0
// @description Thrift is a software framework for managing personal finances.
// @contact.name API Support
// @contact.url XXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http https
// @accept json multipart/form-data
// @produce json

func main() {
	err := api.Setup()
	if err != nil {
		panic(err)
	}
}
