package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
	"strings"
)

const (
	CSS_PREFIX = `/* served by themes-server */
`
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

func (e *echoServer) GetThemeByHost(ctx context.Context, req *pb.HostThemeRequest) (*pb.ThemeResponse, error) {

	// default
	res := &pb.ThemeResponse{
		SmallLogoName: "cnw_logo.png",
		FavIconName:   "cnw_favicon.ico",
		HeaderText:    "Conrad Wood (cnw)",
		ThemeName:     "default",
	}

	if strings.Contains(req.Host, "singingcat") {
		res = &pb.ThemeResponse{
			SmallLogoName: "singingcat_logo.png",
			FavIconName:   "favicon.ico",
			HeaderText:    "SingingCat - IoT as a Service",
			ThemeName:     "singingcat",
		}
	}
	if strings.Contains(req.Host, "youritguru") {
		res = &pb.ThemeResponse{
			SmallLogoName: "guru.jpeg",
			FavIconName:   "favicon.ico",
			HeaderText:    "The IT Guru - All of Information Technology as a Service",
			ThemeName:     "youritguru",
		}
	}
	if strings.Contains(req.Host, "ionos") {
		res = &pb.ThemeResponse{
			SmallLogoName: "ionos_logo.png",
			FavIconName:   "ionos.ico",
			HeaderText:    "IONOS IoT",
			CorporateCss:  "ionos.css",
			ThemeName:     "ionos",
		}
	}

	return res, nil
}
func (e *echoServer) GetHTMLTheme(ctx context.Context, req *h2gproxy.ServeRequest) (*pb.ThemeResponse, error) {
	return e.GetThemeByHost(ctx, &pb.HostThemeRequest{Host: req.Host})
}
func (e *echoServer) getFileForTheme(ctx context.Context, req *pb.HostThemeRequest, filename string) ([]byte, error) {
	f, err := utils.FindFile("templates/")
	if err != nil {
		return nil, err
	}
	t, err := e.GetThemeByHost(ctx, req)
	if err != nil {
		return nil, err
	}
	d_filename := f + "/" + t.ThemeName + "/" + filename
	if !utils.FileExists(d_filename) {
		d_filename = f + "/default/" + filename
	}
	if !utils.FileExists(d_filename) {
		fmt.Printf("Warning file \"%s\" does not exist", d_filename)
	}
	fmt.Printf("File: %s\n", d_filename)
	u, err := utils.ReadFile(d_filename)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (e *echoServer) GetLogo(ctx context.Context, req *pb.HostThemeRequest) (*pb.Image, error) {
	f, err := e.getFileForTheme(ctx, req, "logo.png")
	if err != nil {
		return nil, err
	}
	res := &pb.Image{
		Filename: "logo.png",
		MimeType: "image/png",
		Data:     f,
	}
	return res, nil
}
func (e *echoServer) GetCSS(ctx context.Context, req *pb.HostThemeRequest) (*pb.CSS, error) {
	f, err := e.getFileForTheme(ctx, req, "stylesheet.css")
	if err != nil {
		return nil, err
	}
	s := CSS_PREFIX + string(f)
	res := &pb.CSS{
		Filename: "stylesheet.css",
		Data:     []byte(s),
	}
	return res, nil
}
func (e *echoServer) GetFavIcon(ctx context.Context, req *pb.HostThemeRequest) (*pb.Image, error) {
	f, err := e.getFileForTheme(ctx, req, "favicon.ico")
	if err != nil {
		return nil, err
	}
	res := &pb.Image{
		Filename: "favicon.ico",
		MimeType: "image/png",
		Data:     f,
	}
	return res, nil
}
func (e *echoServer) GetHeaderText(ctx context.Context, req *pb.HostThemeRequest) (*pb.Text, error) {
	f, err := e.getFileForTheme(ctx, req, "heading.txt")
	if err != nil {
		return nil, err
	}
	s := string(f)
	s = strings.Trim(s, "\n")
	res := &pb.Text{
		Text: s,
	}
	return res, nil
}
func (e *echoServer) ServeURL(ctx context.Context, req *h2gproxy.ServeRequest) (*pb.ServeResponse, error) {
	res := &pb.ServeResponse{
		IsAuthoritative: false,
	}
	return res, nil
}
