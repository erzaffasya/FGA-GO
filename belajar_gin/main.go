package main

import (
	"belajar_gin/routers"
)

func main()  {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)	
}