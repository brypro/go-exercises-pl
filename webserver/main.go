package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
