package main

import (
	"github.com/crudecample/inittializers"
	model "github.com/crudecample/module"
)

func init() {
	// inittializers.LoadEnvvariables()
	inittializers.ConnectToDB()
}
func main() {
	inittializers.DB.AutoMigrate(&model.Post{})

}
