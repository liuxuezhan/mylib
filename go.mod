module mylib

go 1.14

require (
	Utils/db v0.0.0-00010101000000-000000000000
	Utils/signal v0.0.0-00010101000000-000000000000
	Utils/wlog v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/chenjiandongx/go-echarts v0.0.0-20190915064101-cbb3b43ade5d
	github.com/chromedp/cdproto v0.0.0-20200209033844-7e00b02ea7d2
	github.com/chromedp/chromedp v0.5.3
	github.com/gin-gonic/gin v1.6.2
	github.com/go-echarts/go-echarts v0.0.0-20190915064101-cbb3b43ade5d // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/valyala/fasthttp v1.12.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace Utils/signal => ./Utils/signal

replace Utils/db => ./Utils/db

replace Utils/wlog => ./Utils/wlog

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace proto_micro => ./proto_micro
