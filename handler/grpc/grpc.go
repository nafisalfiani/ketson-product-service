package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-playground/validator/v10"
	jwtAuth "github.com/nafisalfiani/ketson-go-lib/auth"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/security"
	"github.com/nafisalfiani/ketson-product-service/handler/grpc/category"
	"github.com/nafisalfiani/ketson-product-service/handler/grpc/region"
	"github.com/nafisalfiani/ketson-product-service/handler/grpc/ticket"
	"github.com/nafisalfiani/ketson-product-service/handler/grpc/wishlist"
	"github.com/nafisalfiani/ketson-product-service/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Interface interface {
	Run()
}

type Config struct {
	Base string `env:"GRPC_BASE"`
	Port int    `env:"GRPC_PORT"`
}

type grpcServer struct {
	cfg    Config
	log    log.Interface
	auth   jwtAuth.Interface
	server *grpc.Server
}

func Init(cfg Config, log log.Interface, uc *usecase.Usecases, sec security.Interface, jwtAuth jwtAuth.Interface, validator *validator.Validate) Interface {
	srv := &grpcServer{
		cfg:  cfg,
		log:  log,
		auth: jwtAuth,
	}

	s := grpc.NewServer(
	// grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(srv.authFunc)),
	)

	category.RegisterCategoryServiceServer(s, category.Init(log, uc.Category))
	region.RegisterRegionServiceServer(s, region.Init(log, uc.Region))
	ticket.RegisterTicketServiceServer(s, ticket.Init(log, uc.Ticket))
	wishlist.RegisterWishlistServiceServer(s, wishlist.Init(log, uc.Wishlist))

	reflection.Register(s)

	srv.server = s

	return srv
}

func (g *grpcServer) Run() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", g.cfg.Base, g.cfg.Port))
	if err != nil {
		g.log.Fatal(ctx, err)
	}

	g.log.Info(ctx, fmt.Sprintf("Listening and Serving GRPC on %v:%v", g.cfg.Base, g.cfg.Port))
	if err := g.server.Serve(listener); err != nil {
		g.log.Fatal(ctx, err)
	}
}

// wrapper to connect to grpc package
func Dial(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(target, opts...)
}

// wrapper to connect to grpc package
func WithInsecure() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}

// func (g *grpcServer) authFunc(ctx context.Context) (context.Context, error) {
// 	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
// 	if err != nil {
// 		return nil, status.Error(codes.Unauthenticated, "token not provided")
// 	}

// 	user, err := g.auth.VerifyToken(ctx, token)
// 	if err != nil {
// 		return nil, status.Error(codes.Unauthenticated, err.Error())
// 	}

// 	return g.auth.SetUserAuthInfo(ctx, user, &jwtAuth.Token{TokenType: header.AuthorizationBearer, AccessToken: token}), nil
// }
