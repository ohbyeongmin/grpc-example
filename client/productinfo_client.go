package main

import (
	"context"
	"log"
	pb "productinfo/client/ecommerce" // 프로토버프 컴파일러로 생성한 코드가 포함된 패키지를 임포트한다.
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// 제공된 주소로 서버와의 커넥션을 설정한다.
	// 여기서는 클라이언트와 서버 사이에 보안되지 않은 커넥션을 만든다.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 커넥션을 전달해 스텁을 생성한다. 이 스텁 인스턴스에는 서버를 호출하는 모든 원격 메서드가 포함돼 있다.
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := "iPhone 11 version"

	// 원격 호출과 함께 전달할 Context를 생성한다. 이 Context 객체에는 최종 사용자 인증 토큰에 대한 ID와
	// 요청 데드라인 같은 메타데이터가 포함되며 서비스 요청 동안 유지된다.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 제품 정보와 함께 addProduct 메서드를 호출한다.
	// 호출이 완료되면 제품 ID를 반환하고 그렇지 않으면 에러가 발생한다.
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add Product: %v", err)
	}
	log.Printf("Product ID: %s added syccessfully", r.Value)

	// 제품 ID로 getProduct를 호출한다.
	// 호출이 완료되면 제품 정보를 반환하고 그렇지 않으면 에러가 반환된다.
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get Product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
