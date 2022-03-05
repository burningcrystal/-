package main

import (
	"calloflife/model"
	"calloflife/routes"
)

func main(){
	model.InitDb()
	routes.InitRouter()
}

