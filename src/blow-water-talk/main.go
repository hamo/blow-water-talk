package main

import (
)

func main() {

	serv := newAPIServer()
	panic(serv.ListenAndServe())
}
