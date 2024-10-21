package main

import (
	"bytes"
	"context"

	"html/template"

	pb "golang.conradwood.net/apis/themes"
	"golang.conradwood.net/go-easyops/utils"
)

type defaults struct {
	CorporateTitle string
	Title          string
	UserID         string
	IP             string
	Links          []*Link
}
type Link struct {
	LinkName string
	LinkID   string
}

func apply_defaults(ctx context.Context, req *pb.HostThemeRequest, text []byte) []byte {
	def := &defaults{
		CorporateTitle: "YACloud",
		Title:          "yet-another-cloud",
		UserID:         "",
		IP:             "",
	}
	t, err := template.New("foo").Parse(string(text))
	utils.Bail("invalid template in config", err)
	res := &bytes.Buffer{}
	err = t.Execute(res, def)
	utils.Bail("failed to fix up default template", err)
	return res.Bytes()
}
