package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xNok/go-grpc-demo/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	saveCmd := flag.NewFlagSet("save", flag.ExitOnError)
	saveTitle := saveCmd.String("title", "", "The title of the note")
	saveBody := saveCmd.String("body", "", "The body of the note")

	loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
	loadKeyword := loadCmd.String("keyword", "", "The keyword to search for")

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close()
	c := notes.NewNotesClient(conn)

	if len(os.Args) < 2 {
		fmt.Println("save or load subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "save":
		saveCmd.Parse(os.Args[2:])
		_, err := c.Save(ctx, &notes.Note{Title: *saveTitle, Body: []byte(*saveBody)})
		if err != nil {
			log.Fatalf("could not save note: %v", err)
		}
		fmt.Printf("Note saved: %s", saveTitle)
	case "load":
		loadCmd.Parse(os.Args[2:])
		n, err := c.Load(ctx, &notes.NoteSearch{Keyword: *loadKeyword})
		if err != nil {
			log.Fatalf("could not load note: %v", err)
		}
		fmt.Printf("Note loaded: %s", n.Title)

	default:
		fmt.Println("Expected save or load subcommand")
		os.Exit(1)
	}

}
