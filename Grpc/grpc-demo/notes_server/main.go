package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/xNok/go-grpc-demo/notes"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type notesServer struct {
	notes.UnimplementedNotesServer
}

func (s *notesServer) Save(ctx context.Context, n *notes.Note) (*notes.NoteSaveReply, error) {
	log.Printf("Saving note %v", n.Title)
	err := notes.SaveToDisk(n, "testdata")
	if err != nil {
		return &notes.NoteSaveReply{Saved: false}, err
	}
	return &notes.NoteSaveReply{Saved: true}, nil
}
func (s *notesServer) Load(ctx context.Context, search *notes.NoteSearch) (*notes.Note, error) {
	log.Printf("Received a note to load: %v", search.Keyword)
	n, err := notes.LoadFromDisk(search.Keyword, "testdata")

	if err != nil {
		return &notes.Note{}, err
	}

	return n, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// Register server method (actions the server will do)
	notes.RegisterNotesServer(s, &notesServer{})
	log.Printf("Server listening on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
