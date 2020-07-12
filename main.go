package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/atomotic/annotorius-desktop/annotations"
	_ "github.com/atomotic/annotorius-desktop/statik"
	"github.com/rakyll/statik/fs"
	"github.com/zserge/webview"
)

func main() {
	api := annotations.Api{}
	err := api.Init()
	if err != nil {
		fmt.Println(err)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// fs := http.FileServer(http.Dir("./public"))
		fs := http.FileServer(statikFS)
		http.Handle("/", fs)

		fmt.Println("annotorius-desktop running on: http://" + ln.Addr().String())
		log.Fatal(http.Serve(ln, nil))
	}()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("EDL")
	w.SetSize(1000, 800, webview.HintNone)
	w.Bind("save", api.Save)
	w.Bind("get", api.Get)
	w.Bind("update", api.Update)
	w.Bind("del", api.Delete)

	w.Navigate("http://" + ln.Addr().String())
	w.Run()
}
