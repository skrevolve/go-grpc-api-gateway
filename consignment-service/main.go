package main

import (
	pb "grpc/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

