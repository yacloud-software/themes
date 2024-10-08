package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/h2gproxy"
	pb "golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	CSS_PREFIX = `/* served by themes-server */
`
)

var (
	templ_dir  = flag.String("templates", "templates/", "template directory")
	use_cache  = flag.Bool("use_cache", true, "if false, never cache, load from disk each time")
	port       = flag.Int("port", 4100, "The grpc server port")
	debug      = flag.Bool("debug", false, "debug mode")
	file_cache = cache.New("filecontent", time.Duration(150)*time.Hour, 100)
	fdir       = ""
)

type fcache struct {
	Data  []byte
	Error error
}
type echoServer struct {
}

func main() {
	flag.Parse()
	fmt.Printf("Starting ThemesServer...\n")

	sd := server.NewServerDef()
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterThemesServer(server, e)
			return nil
		},
	))
	err := server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) GetThemeByHost(ctx context.Context, req *pb.HostThemeRequest) (*pb.ThemeResponse, error) {
	debugf("Theme for host \"%s\" for useragent \"%s\"\n", req.Host, req.UserAgent)
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
	} else if strings.Contains(req.Host, "max") {
		res = &pb.ThemeResponse{
			SmallLogoName: "max_logo.png",
			FavIconName:   "favicon.ico",
			HeaderText:    "MaxWoodArt Fashion",
			ThemeName:     "maxwoodart",
		}
	} else if strings.Contains(req.Host, "planetary") {
		res = &pb.ThemeResponse{
			SmallLogoName: "logo.png",
			FavIconName:   "favicon.ico",
			HeaderText:    "Planetary Processing",
			ThemeName:     "planetaryprocessing",
		}
	} else if strings.Contains(req.Host, "youritguru") {
		res = &pb.ThemeResponse{
			SmallLogoName: "guru.jpeg",
			FavIconName:   "favicon.ico",
			HeaderText:    "The IT Guru - All of Information Technology as a Service",
			ThemeName:     "youritguru",
		}
	} else if strings.Contains(req.Host, "ionos") {
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
	b, err := e.getSingleFileForTheme(ctx, req, filename)
	if err != nil {
		return nil, err
	}
	if strings.Contains(strings.ToLower(req.UserAgent), "android") {
		b2, err := e.getSingleFileForTheme(ctx, req, "android_"+filename)
		if err == nil {
			b = append(b, b2...)
		}
	}
	return b, nil
}
func (e *echoServer) getSingleFileForTheme(ctx context.Context, req *pb.HostThemeRequest, filename string) ([]byte, error) {
	debugf("File \"%s\" for host \"%s\" for useragent \"%s\"\n", filename, req.Host, req.UserAgent)
	if fdir == "" {
		f, err := utils.FindFile(*templ_dir)
		if err != nil {
			return nil, err
		}
		fdir = f
	}
	t, err := e.GetThemeByHost(ctx, req)
	if err != nil {
		return nil, err
	}
	i := strings.Index(filename, "/")
	if i != -1 {
		filename = filename[i+1:]
	}
	d_filename := fdir + "/" + t.ThemeName + "/" + filename
	fc := file_cache.Get(d_filename)
	if fc != nil && *use_cache {
		return fc.(*fcache).Data, fc.(*fcache).Error
	}
	if !utils.FileExists(d_filename) {
		d_filename = fdir + "/default/" + filename
	}
	if !utils.FileExists(d_filename) {
		fmt.Printf("Warning file \"%s\" does not exist\n", d_filename)
	}
	debugf("[%s] File: %s\n", t.ThemeName, d_filename)
	u, err := utils.ReadFile(d_filename)
	if err != nil {
		return nil, err
	}
	file_cache.Put(d_filename, &fcache{Data: u})
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
	p := strings.TrimPrefix(req.Path, "/")
	f, err := e.getFileForTheme(ctx, &pb.HostThemeRequest{Host: req.Host}, p)
	if err == nil {
		res.IsAuthoritative = true
		res.Response = buildResponse(p, f)
		return res, err
	}

	return res, nil
}
func buildResponse(filename string, content []byte) *h2gproxy.ServeResponse {
	mt := mime.TypeByExtension(filepath.Ext(filename))
	res := &h2gproxy.ServeResponse{
		MimeType: mt,
		Body:     content,
	}
	return res
}
func debugf(format string, args ...interface{}) {
	if !*debug {
		return
	}
	fmt.Printf(format, args...)
}
