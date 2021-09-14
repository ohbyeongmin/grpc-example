package main

import (
	"log"
	"net"

	pb "productinfo/service/ecommerce" // 프로토버프 컴파일러로 생성된 코드가 포함된 패키지를 임포트

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// gRPC 서버가 바인딩하고자 하는 TCP 리스너는 지정 포트(50051)로 생성된다.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// RPC Go API 를 호출해 새 gRPC 서버 인스턴스를 생성한다.
	s := grpc.NewServer()
	// 이전에 구현된 서비스로 생성된 API를 호출해 새로 작성된 gRPC 서버에 등록한다.
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port " + port)
	// 포트(50051)에서 수신되는 메시지를 리스닝하기 시작한다.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
