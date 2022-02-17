package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

var (
	echoClient pb.ThemesClient
	logo       = flag.Bool("logo", false, "get logo")
	heading    = flag.Bool("heading", false, "get heading")
)

func main() {
	flag.Parse()

	echoClient = pb.GetThemesClient()
	if *logo {
		Logo()
	}
	if *heading {
		Heading()
	}
	fmt.Printf("Done\n")
	os.Exit(0)
}
func Logo() {
	ctx := authremote.Context()
	res, err := echoClient.GetLogo(ctx, getThemeRequest())
	utils.Bail("failed to get logo", err)
	fmt.Printf("Filename: %s\n", res.Filename)
}
func Heading() {
	ctx := authremote.Context()
	res, err := echoClient.GetHeaderText(ctx, getThemeRequest())
	utils.Bail("failed to get heading", err)
	fmt.Printf("Heading: \"%s\"\n", res.Text)
}
func getThemeRequest() *pb.HostThemeRequest {
	if len(flag.Args()) == 0 {
		fmt.Printf("Missing host\n")
		os.Exit(10)
	}
	res := &pb.HostThemeRequest{Host: flag.Args()[0]}
	return res
}
