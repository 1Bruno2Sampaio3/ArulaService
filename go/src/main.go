package main

import "os"
import "fmt"

func main() {

	fmt.Println(os.Getenv("ARULA_DB_USERNAME"),
		os.Getenv("ARULA_DB_PASSWORD"),
		os.Getenv("ARULA_DB_NAME"))

	a := App{}
	a.Init(
		os.Getenv("ARULA_DB_USERNAME"),
		os.Getenv("ARULA_DB_PASSWORD"),
		os.Getenv("ARULA_DB_NAME"))

	a.Run(":7777")
}
