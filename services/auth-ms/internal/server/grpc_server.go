package server

//
//func StartGRPCServer(cfg *config.Config, userRepo repository.UserRepository) error {
//	lis, err := net.Listen("tcp", cfg.GRPCPort)
//	if err != nil {
//		return err
//	}
//
//	grpcServer := grpc.NewServer()
//	authHandler := handlers.NewAuthGRPCHandler(userRepo)
//	authpb.RegisterAuthServiceServer(grpcServer, authHandler)
//
//	return grpcServer.Serve(lis)
//}
