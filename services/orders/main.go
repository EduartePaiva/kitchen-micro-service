package main

func main() {
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()

	httpServer := NewHttpServer(":9001")
	httpServer.Run()
}
