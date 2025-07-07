package main

import "log"

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {

	return nil
}
