package main

import (
	"flag"
	"main/Notes/NotesServer/controller/httpserver"
	"main/Notes/NotesServer/gates/storage"
	"main/Notes/NotesServer/gates/storage/list"
	"main/Notes/NotesServer/gates/storage/mp"
)

func main() {
	var st storage.Storage
	var useList bool
	flag.BoolVar(&useList, "list", false, "A boolean flag")
	flag.Parse()

	if useList {
		st = list.NewList()
	} else {
		st = mp.NewMap()
	}
	hs := httpserver.NewHttpServer(":8080", st)
	hs.Start()
}
