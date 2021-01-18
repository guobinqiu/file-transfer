package main

import "os"

func main() {
	f, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	f.Truncate(1*1024*1024*1024) //1GB

	f.Close()
}
