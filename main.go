package main

import (
	"runtime"
	"net/http"
	"flag"
	"log"
	"os"
	"strconv"
	"golang.org/x/time/rate"
	"github.com/lattecake/emoticon/handler"
	"github.com/lattecake/emoticon/core"
)

var (
	dir, path string
	port      int
	err       error
	limiter   *rate.Limiter
	burst     int
	setLimit  float64
	fs        = flag.NewFlagSet("emoticon", flag.ExitOnError)
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fs.IntVar(&port, "port", 8080, "服务器端口")
	fs.StringVar(&dir, "static", "./static", "静态目录")
	fs.StringVar(&path, "path", "./upload", "静态目录")
	fs.Float64Var(&setLimit, "limit", 2, "每几秒")
	fs.IntVar(&burst, "burst", 5, "并发数")
}

func main() {
	fs.Parse(os.Args[1:])

	limiter = rate.NewLimiter(rate.Limit(setLimit), burst)
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	mux.Handle("/", &handler.IndexHandler{})

	mux.Handle("/upload/image", &handler.ImageHandler{Path: path})

	log.Println("service", "start", "httpPort", port)

	if err = http.ListenAndServe(":"+strconv.Itoa(port), core.NewLimiter(limiter).Limit(mux)); err != nil {
		panic(err)
	}

	//g := &group.Group{}
	//httpListener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	//if err != nil {
	//	log.Println("transport", "HTTP", "during", "Listen", "err", err.Error())
	//	panic(err)
	//}
	//g.Add(func() error {
	//	return http.ListenAndServe(":"+strconv.Itoa(port), limit(mux))
	//}, func(err error) {
	//	log.Println("transport", "HTTP", "during", "Listen", "err", err.Error())
	//})
}
