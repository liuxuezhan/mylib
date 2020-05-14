module proto_micro

go 1.14

require (
	Utils/db v0.0.0-00010101000000-000000000000
	Utils/wlog v0.0.0-00010101000000-000000000000
	github.com/dghubble/oauth1 v0.6.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/protobuf v1.3.5
	github.com/klauspost/compress v1.10.5
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.0
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	google.golang.org/grpc v1.29.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/square/go-jose.v1 v1.1.2 // indirect
)

replace Utils/wlog => ../Utils/wlog

replace Utils/db => ../Utils/db

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
