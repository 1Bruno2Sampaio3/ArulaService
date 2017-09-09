package main

import "os"

func main() {

	a := App{}
	a.Init(
		os.Getenv("ARULA_DB_USERNAME"),
		os.Getenv("ARULA_DB_PASSWORD"),
		os.Getenv("ARULA_DB_NAME"))

	a.Run(":7777")
}
