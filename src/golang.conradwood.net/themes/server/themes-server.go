package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
	"strings"
)

var (
	port = flag.Int("port", 4100, "The grpc server port")
)

type echoServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting ThemesServer...\n")

	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterThemesServer(server, e)
			return nil
		},
	)
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) Ping(ctx context.Context, req *common.Void) (*pb.PingResponse, error) {
	resp := &pb.PingResponse{Response: "pingresponse"}
	return resp, nil
}
func (e *echoServer) GetThemeByHost(ctx context.Context, req *pb.HostThemeRequest) (*pb.ThemeResponse, error) {

	// default
	res := &pb.ThemeResponse{
		SmallLogoName: "cnw_logo.png",
		FavIconName:   "cnw_favicon.ico",
		HeaderText:    "Conrad Wood (cnw)",
	}

	if strings.Contains(req.Host, "singingcat") {
		res = &pb.ThemeResponse{
			SmallLogoName: "singingcat_logo.png",
			FavIconName:   "favicon.ico",
			HeaderText:    "SingingCat - IoT as a Service",
		}
	}
	if strings.Contains(req.Host, "youritguru") {
		res = &pb.ThemeResponse{
			SmallLogoName: "guru.jpeg",
			FavIconName:   "favicon.ico",
			HeaderText:    "The IT Guru - All of Information Technology as a Service",
		}
	}
	if strings.Contains(req.Host, "ionos") {
		res = &pb.ThemeResponse{
			SmallLogoName: "ionos_logo.png",
			FavIconName:   "ionos.ico",
			HeaderText:    "IONOS IoT",
			CorporateCss:  "ionos.css",
		}
	}

	return res, nil
}
func (e *echoServer) GetHTMLTheme(ctx context.Context, req *h2gproxy.ServeRequest) (*pb.ThemeResponse, error) {
	return e.GetThemeByHost(ctx, &pb.HostThemeRequest{Host: req.Host})
}
