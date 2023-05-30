package main

import (
	"context"
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	pb "github.com/mariobenissimo/my_grpc/routeguide" // Update with your actual package path
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetSum(ctx context.Context, point *pb.Point) (*pb.Sum, error) {
	z := point.X + point.Y
	return &pb.Sum{Z: z}, nil
}
func quadratic(x int32) int32 {
	return int32(math.Pow(float64(x), 2))
}
func cubic(x int32) int32 {
	return int32(math.Pow(float64(x), 3))
}
func (s *routeGuideServer) GetYFromFunction(x *pb.X, stream pb.RouteGuide_GetYFromFunctionServer) error {
	operations := [2]func(int32) int32{
		quadratic,
		cubic,
	}
	for _, operation := range operations {
		result := operation(x.X)
		if err := stream.Send(&pb.Y{Y: result}); err != nil {
			return err
		}
	}
	return nil
}

func (s *routeGuideServer) GetSumFromPoints(stream pb.RouteGuide_GetSumFromPointsServer) error {

	var sum int32
	sum = 0
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&pb.Sum{
					Z: sum,
				})
		}
		if err != nil {
			return err
		}
		sum += point.X + point.Y
	}
}
func main() {
	lis, err := net.Listen("tcp", ":50051") // Choose your desired port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRouteGuideServer(server, &routeGuideServer{})

	log.Println("Server started on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
