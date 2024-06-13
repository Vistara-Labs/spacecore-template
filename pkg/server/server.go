package server

// "spacecore/pkg/proto"
// "spacecore/pkg/utils"

// func (g *config.MySpacecore) Start(ctx context.Context) (string, error) {
// 	// I want this start to be a long running process
// 	// like running a binary that exposes a REST API
// 	// or a long running process that listens to a port
// 	// run vimana run celestia light-node command

// 	binPath := "/Users/mayurchougule/development/vistara/vimana-cli/tmp/vimana/vimana-darwin-arm64/vimana"
// 	c := exec.Command(binPath, "run", "celestia", "light-node")
// 	err := c.Run()
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// 	return c.Path, nil

// 	// return "Statred updated", nil
// }

// func (g *MySpacecore) Stop(ctx context.Context) (string, error) {
// 	fmt.Println("Spacecore stopping...")
// 	return "Stopped updated", nil
// }

// func (g *MySpacecore) Status(ctx context.Context) (string, error) {
// 	return "Running blah blah \n nows updated \n", nil
// }

// func (g *MySpacecore) Logs(ctx context.Context) (string, error) {
// 	return "Logs updated", nil
// }

// // func (g *config.GRPCServer) Start(ctx context.Context, req *proto.StartRequest) (*proto.StartResponse, error) {
// // 	msg, err := g.Impl.Stop(ctx)
// // 	fmt.Printf("Starting Spacecore logs...\n %s", msg)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &proto.StartResponse{}, nil
// // }

// // func (g *config.GRPCServer) Stop(context.Context, *proto.StopRequest) (*proto.StopResponse, error) {
// // 	msg, err := g.Impl.Logs(context.Background())
// // 	fmt.Printf("Stopping Spacecore logs...\n %s", msg)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &proto.StopResponse{}, nil
// // }

// // func (g *config.GRPCServer) Logs(context.Context, *proto.LogsRequest) (*proto.LogsResponse, error) {
// // 	msg, err := g.Impl.Logs(context.Background())
// // 	fmt.Printf("Checking Spacecore logs...\n %s", msg)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &proto.LogsResponse{}, nil
// // }

// // func (g *config.GRPCServer) Status(context.Context, *proto.StatusRequest) (*proto.StatusResponse, error) {
// // 	msg, err := g.Impl.Logs(context.Background())
// // 	fmt.Printf("Checking Spacecore status...\n %s", msg)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &proto.StatusResponse{}, nil
// // }
