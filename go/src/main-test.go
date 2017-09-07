package main

import (
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {

	a = App{}
	a.Init(
		os.Getenv("ARULA_DB_USERNAME"),
		os.Getenv("ARULA_DB_PASSWORD"),
		os.Getenv("ARULA_DB_NAME"))

	ensureTablesExists()

	code := m.Run()

	clearTable()

	apiTest()

	os.Exit(code)

}

// FIXME: Paulo
func ensureTablesExists() {
	a.Session.
}

// FIXME: Paulo
func clearTable() {}

// FIXME: Paulo
func apiTest() {}
