package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "webseries.com/grpc/protos"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSeriesClient(conn)

	runGetAllSeries(client)
	// runGetMovie(client, "1")
	// runCreateMovie(client, "2", "Deadpool",
	// 	"Abhinav", "Anand")
	// runUpdateMovie(client, "98498081", "24325645", "Spiderman Spiderverse",
	// 	"Peter", "Parker")
	// runDeleteMovie(client, "98498081")
}

func runGetAllSeries(client pb.SeriesClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetAllSeries(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetMovies(_) = _, %v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetAllSeries(_) = _, %v", client, err)
		}
		log.Printf("SeriesInfo: %v", row)
	}
}

func runGetMovie(client pb.SeriesClient, movieid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: movieid}
	res, err := client.GetSeries(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetSeries(_) = _, %v", client, err)
	}
	log.Printf("SeriesInfo: %v", res)
}

func runCreateMovie(client pb.SeriesClient, isbn string,
	title string, firstname string, lastname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.SeriesInfo{Isbn: isbn, Title: title,
		Character: &pb.Character{Firstname: firstname,
			Lastname: lastname}}
	res, err := client.AddSeries(ctx, req)
	if err != nil {
		log.Fatalf("%v.AddSeries(_) = _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("AddSeries Id: %v", res)
	} else {
		log.Printf("AddSeries Failed")
	}

}

func runUpdateMovie(client pb.SeriesClient, movieid string,
	isbn string, title string, firstname string, lastname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.SeriesInfo{Id: movieid, Isbn: isbn,
		Title: title, Character: &pb.Character{
			Firstname: firstname, Lastname: lastname}}
	res, err := client.UpdateSeries(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateSeries(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateSeries Success")
	} else {
		log.Printf("UpdateSeries Failed")
	}
}

func runDeleteMovie(client pb.SeriesClient, movieid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: movieid}
	res, err := client.DeleteSeries(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteSeries(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteSeries Success")
	} else {
		log.Printf("DeleteSeries Failed")
	}
}
