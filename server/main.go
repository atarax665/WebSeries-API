package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	pb "webseries.com/grpc/protos"
	"net"
	"strconv"
)

const (
	port = ":50051"
)

var series []*pb.SeriesInfo

type seriesServer struct {
	pb.UnimplementedSeriesServer
}

func main() {
	initSeries()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterSeriesServer(s, &seriesServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSeries() {
	movie1 := &pb.SeriesInfo{Id: "1", Isbn: "0593310438",
		Title: "Money Hiest", Character: &pb.Character{
			Firstname: "Alvaro", Lastname: "Morte"}}
	movie2 := &pb.SeriesInfo{Id: "2", Isbn: "3430220302",
		Title: "Broken but Beautiful", Character: &pb.Character{
			Firstname: "Vikrant", Lastname: "Massey"}}
	series = append(series, movie1)
	series = append(series, movie2)
}

func (s *seriesServer) GetAllSeries(in *pb.Empty,
	stream pb.Series_GetAllSeriesServer) error {
	log.Printf("Received: %v", in)
	for _, serie := range series {
		if err := stream.Send(serie); err != nil {
			return err
		}
	}
	return nil
}

func (s *seriesServer) GetSeries(ctx context.Context,
	in *pb.Id) (*pb.SeriesInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.SeriesInfo{}

	for _, serie := range series {
		if serie.GetId() == in.GetValue() {
			res = serie
			break
		}
	}

	return res, nil
}

func (s *seriesServer) AddSeries(ctx context.Context,
	in *pb.SeriesInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000000))
	in.Id = res.GetValue()
	series = append(series, in)
	return &res, nil
}

func (s *seriesServer) UpdateSeries(ctx context.Context,
	in *pb.SeriesInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, serie := range series {
		if serie.GetId() == in.GetId() {
			series = append(series[:index], series[index+1:]...)
			in.Id = serie.GetId()
			series = append(series, in)
			res.Value = 1
			break
		}
	}

	return &res, nil
}

func (s *seriesServer) DeleteSeries(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for index, movie := range series {
		if movie.GetId() == in.GetValue() {
			series = append(series[:index], series[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}
