package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "github.com/mariobenissimo/my_grpc/routeguide" // Update with your actual package path
)

func randomPoint(r *rand.Rand) *pb.Point {
	x := r.Int31n(10)
	y := r.Int31n(10)
	log.Printf("%d , %d\n", x, y)
	return &pb.Point{X: x, Y: y}
}

func runSumPoint(client pb.RouteGuideClient) {
	// Create a random number of random points
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(10)) + 2 // Traverse at least two points
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	log.Printf("Il numero di punti generati è:  %d ", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetSumFromPoints(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("La somma è: %d", reply.Z)
}
func printFunction(client pb.RouteGuideClient, x *pb.X) {
	log.Printf("Osservazioni delle funzioni quadratiche e cubiche del valore %d", x.X)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetYFromFunction(ctx, x)
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("Output : %d", feature.Y)
	}
}
func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	req := &pb.Point{X: 1, Y: 2}

	resp, err := client.GetSum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call getSum: %v", err)
	}
	fmt.Println("la risposta è: ", resp.Z)

	runSumPoint(client)
	printFunction(client, &pb.X{X: 5})
}
