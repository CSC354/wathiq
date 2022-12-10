package main

import (
	context "context"
	"log"
	"net"
	"time"

	"github.com/CSC354/wathiq/proto"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
)

var KEY = []byte("DEBATE")

type Waithq struct {
	proto.UnimplementedWathiqServer
}

// GetToken implements proto.WathiqServer
func (*Waithq) GetToken(ctx context.Context, req *proto.TokenRequest) (response *proto.TokenResponse, err error) {
	expirationTime := time.Now().Add(60 * 24 * time.Hour)
	claims := &Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(KEY)
	if err != nil {
		response.Error = int32(proto.Zalh_ErrGeneratingToken)
		return
	}
	response.Token = tokenString
	response.Error = int32(proto.Zalh_Ok)
	return

}

// Validate implements proto.WathiqServer
func (*Waithq) Validate(ctx context.Context, req *proto.ValidateRequest) (res *proto.ValidateResponse, err error) {
	tknStr := req.Token
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return KEY, nil
	})
	if err != nil || !tkn.Valid {
		res.Valid = false
		return
	}
	res.Valid = true
	res.Id = claims.Username
	return
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterWathiqServer(grpcServer, &Waithq{})
	go print("listening..")
	grpcServer.Serve(lis)
}
