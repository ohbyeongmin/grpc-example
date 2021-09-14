package main

// 비즈니스 로직 구현

import (
	"context"
	pb "productinfo/service/ecommerce" // 프로토버프 컴파일러로 생성된 코드가 포함된 패키지를 임포트

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server 구조체는 서버에 대한 추상화로 서버에 서비스 메서드를 지정한다.
type server struct {
	productMap map[string]*pb.Product
}

// 두 메서드 모두 Context 파라미터를 갖는데, Context 객체에는 최종 사용자
// 인증 토큰의 ID와 ㅇ쵸청 데드라인과 같은 메타데이터가 포함되며, 서비스 요청 동안 존재한다

// 두 메서드 모두 원격 메서드의 반환값 외에 에러를 반환한다
// 이 에러는 소비자에게 전파되며 소비자 측에서 에러 처리에 사용된다.

// AddProduct는 ecommerce.AddProduct 를 구현한다.
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct는 ecommerce.GetProduct 를 구현한다.
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist", in.Value)
}
