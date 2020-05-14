package proto_micro

import (
	"Utils/db"
	"Utils/wlog"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/klauspost/compress/gzip"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"golang.org/x/crypto/xtea"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type HelloWorld  struct{}
type MessageCacheProcessParam struct {
	Bytes       []byte
	Version     uint64
	RepeatTimes uint32
}

type LoginResult struct {
	Result LoginReCode
	Uid    string
	Host   string
	Port   string
}
type LoginParam struct {
	AccountName  string
	IDToken      string
	Password     string
	Secret       string
	LoginType    uint32
	ResultChanId int64
	RemoteIP     string
}
const WindPlayGzipON uint32 = 0xabababab
const WindPlayGzipOff uint32 = 0xababcccc

type GzipWriter struct {
	buffer bytes.Buffer
	writer *gzip.Writer
}
type ST_LoginServer struct {
	xTeaInstance  *xtea.Cipher
	TwitterConfig *oauth1.Config
}
var (
	b     broker.Broker
	Login_Instance = &ST_LoginServer{}
	IsGzipOpen = false
	GzipMiniSize int32 = 512
	Conf = &AutoGenerated{}
)

type AutoGenerated struct {
	Loginserver struct {
		IP          string `toml:"ip"`
		Sid         int64    `toml:"sid"`
		Port        string `toml:"port"`
		Mongodbinfo struct {
			ConnectionString string   `toml:"connection_string"`
			DbName           string   `toml:"db_name"`
			Table            []string `toml:"table"`
		} `toml:"mongodbinfo"`
	} `toml:"loginserver"`
	Log struct {
		Loglevel uint32  `toml:"loglevel"`
		Tostd    bool `toml:"tostd"`
	} `toml:"log"`
	Google struct {
		Topic            string `toml:"topic"`
		GoogleClientID   string `toml:"GoogleClientID"`
		XTeaKey          string `toml:"XTeaKey"`
		FBAppID          string `toml:"FBAppID"`
		FBAppSecret      string `toml:"FBAppSecret"`
		TTConsumerKey    string `toml:"TTConsumerKey"`
		TTConsumerSecret string `toml:"TTConsumerSecret"`
	} `toml:"google"`
}
func LoginServer() {//主函数
	db.TomlDecode("./server.toml", Conf)
	wlog.Initialize(Conf.Log.Loglevel,Conf.Log.Tostd,Conf.Loginserver.Sid )
	//>>>初始化数据库连接
	db.InitDBConnection(Conf.Loginserver.Mongodbinfo.DbName,Conf.Loginserver.Mongodbinfo.ConnectionString,
		Conf.Loginserver.Mongodbinfo.Table)
	c, err := xtea.NewCipher([]byte(Conf.Google.XTeaKey))
	if err != nil {
		wlog.Error(err.Error())
		panic(err.Error())
	}

	Login_Instance.xTeaInstance = c

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/quicklogin", quickLogin)
	router.GET("/quicklogin", quickLogin)
	router.GET("/checkstatus", func (ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	router.GET("/akamai/sureroute-test-object.html", func (ctx *gin.Context) {
		ctx.File("./Data/sureroute-test-object.html")
	})

	port := Conf.Loginserver.Port
	addr := ":" + port
	wlog.Info("Server Started @ ", addr)
	if err := router.RunTLS(":"+port, "./Data/server.pem", "./Data/server.key"); err == nil {
		wlog.Info("Server Started success!!")
	} else {
		wlog.Error(err)
	}

	//-------------------------------------------------------------------------------
	if err := broker.Init(); err != nil {
		wlog.Info("broker.Init() error:", err)
	}
	if err := broker.Connect(); err != nil {
		wlog.Info("broker.Connect() error:", err)
	}

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		micro.Flags(&cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
			return nil
		}),
	)
	runServer(service)
}
func (g *HelloWorld) Hello(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Greeting = "Hello World: " + req.Name
	return nil
} // 实现hello_world service中Hello方法
type rpcServer struct{}
func (*rpcServer) AddList(ctx context.Context, msg *ST_CS2LS_Add_List) (*ST_LS2CS_Add_List_Result, error) {
	return nil, nil
}
func (*rpcServer) RemoveList(ctx context.Context, msg *ST_CS2LS_Remove_List) (*ST_LS2CS_Remove_List_Result, error) {
	return nil, nil
}
func (*rpcServer) OpenList(ctx context.Context, msg *ST_Open_List) (*ST_Open_List_Result, error) {
	return nil, nil
}
func (*rpcServer) GetListStatus(ctx context.Context, msg *ST_Get_List_Status) (*ST_Get_List_Status_Result, error) {
	return nil, nil
}
func quickLogin(ctx *gin.Context) {
	req, res := Prepare(ctx)
	defer res.Send()

	userName := req.GetString("username")
	password := req.GetString("password")
	loginType := req.GetUint32("logintype")
	idToken := req.GetString("idtoken")
	secret := req.GetString("secret")
	puk := req.GetString("puk")

	remoteIP := ctx.ClientIP()

	result := Login_Instance.Login(userName, password, idToken, secret, remoteIP, loginType, puk)
	res.Add("status", result.Result)
	res.Add("uid", result.Uid)
	res.Add("host", result.Host)
	res.Add("port", result.Port)
}
func (this *ST_LoginServer) Login(userName, password, idToken, secret, remoteIP string, loginType uint32, puk string) *LoginResult {
	request := &LoginParam{
		AccountName: userName,
		Password:    password,
		IDToken:     idToken,
		LoginType:   loginType,
		Secret:      secret,
		RemoteIP:    remoteIP,
	}
	result := &LoginResult{
		Result: LoginReCode_LRC_Unknow,
	}
	if len(userName) <= 0 {
		wlog.Info("Login err!!!! userName err", userName)
		return result
	}

	// 黑名单禁止登录
	if this.IsInList(ListType_LT_BLACK, remoteIP) ||
		this.IsInList(ListType_LT_BLACK, userName) {
		result.Result = LoginReCode_LRC_FORBID
		wlog.Info("Login err!!!! in blacklist err", userName)
		return result
	}

	user := Get(request,puk)
	if user ==nil {
		nextId := uint64(wlog.GetSFUID())
		now := uint32(wlog.GetNow())
		user := &PB_User{
			AccountId:          nextId,
			DeviceId:      		request.AccountName,
			CurPlayerID:        nextId,
		}
		user.Player =make(map[uint64]*PB_Player)
		if request.LoginType == uint32(LoginType_LT_Visitor) {
			// 游客处理
			user.Player[nextId] = &PB_Player{
				Name:          request.AccountName,
				LastLoginTime: now,
				CreateTime:    now,
			}
			if db.UpsertById("users", user.GetAccountId(),user) != nil {
				return result
			}
		} else {
			user.Player[nextId] = &PB_Player{
				Name:          request.AccountName,
				LastLoginTime: now,
				CreateTime:    now,
			}
			if db.UpsertById("users", user.GetAccountId(),user) != nil {
				return result
			}
		}
	}


	if nil == user {
		wlog.Info("Login err!!!! usr nil")
		return result
	}

	// 黑名单禁止登录
	if this.IsInList(ListType_LT_BLACK, wlog.ConvertToString(user.GetAccountId())) {
		result.Result = LoginReCode_LRC_FORBID
		return result
	}

	serverInfo:=&PB_ServerList{}
	db.GetById("server_list",user.Player[user.CurPlayerID].GetSid(),serverInfo)

	if serverInfo == nil {
		wlog.Error("Alloc Server err", user.GetAccountId())
		result.Result = LoginReCode_LRC_Server_Offline
		return result
	}

	if serverInfo.Status == int32(ServerStatus_SST_Upgrading) {
		// 维护中但没有开白名单
		if serverInfo.WhiteList != 1 {
			result.Result = LoginReCode_LRC_Server_Upgrading
			return result
		}

		if this.IsInList(ListType_LT_WHITE, wlog.ConvertToString(user.GetAccountId())) ||
			this.IsInList(ListType_LT_WHITE, userName) ||
			this.IsInList(ListType_LT_WHITE, remoteIP) {
			result.Result = LoginReCode_LRC_Ok
		} else {
			result.Result = LoginReCode_LRC_Server_Upgrading
		}
	} else {
		result.Result = LoginReCode_LRC_Ok
	}

	// 检查服务器是否在线 todo


	if result.Result == LoginReCode_LRC_Ok {
		wlog.Info(fmt.Sprintf("Login success, accountId:%d, sid: %d, Logintype: %d,  Saddr: %s", user.AccountId, serverInfo.ID, loginType, serverInfo.SAddr))
		result.Uid = encryptUid(user.GetAccountId())
		result.Host = serverInfo.SAddr
	} else {
		wlog.Info("Login err!!! Unknow")
	}

	return result
}
func Get(request *LoginParam,puk string) (*PB_User) {
	query := bson.M{"name": request.AccountName}
	switch request.LoginType {
	case uint32(LoginType_LT_Google):
		// 验证token
		name :=  wlog.VerifyGoogleIDToken(request.IDToken,Conf.Google.GoogleClientID)
		query = bson.M{"google_name": name}
	case uint32(LoginType_LT_FaceBook):
		// 验证token
		name :=  wlog.VerifyFBTokenID(request.IDToken,Conf.Google.FBAppID, Conf.Google.FBAppSecret)
		query = bson.M{"fb_name": name}
	case uint32(LoginType_LT_Twitter):
		// 验证token
		name :=  wlog.VerifyTwitterCredentials(request.IDToken, request.Secret,Conf.Google.TTConsumerKey,Conf.Google.TTConsumerSecret)
		query = bson.M{"twitter_name": name}
	case uint32(LoginType_LT_GameCenter):
		// 验证token
		name :=  wlog.VerifyGameCenter(request.AccountName, request.IDToken, request.Secret, puk)
		query = bson.M{"gamecenter_name": name}
	}
	doc := PB_User{}
	if err := db.Get("uers", query, &doc); nil != err {
		wlog.Error(err.Error())
		return nil
	}
	return &doc
}
func encryptUid(uid uint64) string {
	t := uint32(wlog.GetNow())

	uid = uid<<32 + uint64(t)
	source := make([]byte, 8)
	binary.BigEndian.PutUint64(source, uid)
	dest := make([]byte, xtea.BlockSize)
	Login_Instance.xTeaInstance.Encrypt(dest, source)

	return wlog.Base64Encode(string([]byte(dest)))
}

func (this *ST_LoginServer) IsInList(lType ListType, value string) bool {
	type ListDocument struct {
		ID    string "_id"
		Value string "value"
	}
	if lType == ListType_LT_WHITE {
		query := bson.M{"value": value}
		err := db.Get("white_list", query, &ListDocument{})
		if err != nil {
			return false
		}
		return true
	} else {
		query := bson.M{"value": value}
		err := db.Get("black_list", query, &ListDocument{})
		if err != nil {
			return false
		}
		return true
	}
}
func runServer(service micro.Service) { //接受方
	go test_pub()
	RegisterGreeterHandler(service.Server(), new(HelloWorld))
	if err := service.Run(); err != nil {
		wlog.Error(err)
	}
}
func runClient(service micro.Service) { //发送方
	go test_sub()
	greeter := NewGreeterService("greeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &Request{Name: "John"})
	if err != nil {
		wlog.Error(err)
		return
	}
	wlog.Info(rsp.Greeting)
	<-time.After(time.Second * 20)
}
func test_pub() {//发送同步数据
	// 该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
	// 它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d:%s", i, time.Now().String())),
		}
		wlog.Info(broker.String())
		// 发布消息
		if err := broker.Publish(Conf.Google.Topic, msg); err != nil {
			wlog.Info("[pub] Message publication failed: %v", err)
		} else {
			wlog.Info("[pub] Message published: ", string(msg.Body))
		}
		i++
	}
}
func test_sub() {//接受同步数据
	_, err := broker.Subscribe(Conf.Google.Topic, func(p broker.Event) error {
		wlog.Info("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (this *GzipWriter) init() {
	this.buffer = bytes.Buffer{}
	this.writer = gzip.NewWriter(&this.buffer)
}

func (this *GzipWriter) reset() {
	this.buffer.Reset()
	this.writer.Reset(&this.buffer)
}

func (this *GzipWriter) write(in []byte) (error, []byte) {
	_, err := this.writer.Write(in)
	err = this.writer.Close()

	return err, this.buffer.Bytes()
}

func (this *GzipWriter) bytes() []byte {
	return this.buffer.Bytes()
}

var WriterPool = sync.Pool{
	New: func() interface{} {
		writer := &GzipWriter{}
		writer.init()
		return writer
	},
}

//func init() {
//	binary.BigEndian.PutUint32(BitHeaderOn, WindPlayGzipON)
//	binary.BigEndian.PutUint32(BitHeaderOff, WindPlayGzipOff)
//}

func Initialize(openGzip bool, minSize int32) {
	IsGzipOpen = openGzip
	GzipMiniSize = minSize
	wlog.Info("Initialize gzip %t %d", IsGzipOpen, GzipMiniSize)
}

type MessageCacheProcessHandler interface {
	Handle(userId uint64, input map[string]*MessageCacheProcessParam)
}

type MessageCachePostHandler interface {
	Handle(userId uint64, form *ST_ResponForm_PB)
}

type ST_Request struct {
	ctx *gin.Context
}

func Prepare(ctx *gin.Context) (*ST_Request, *Respon) {
	req := newST_Request(ctx)
	res := newRespon(ctx)
	/*req.GetString("")
	wlog.InfoDepth(1, req.ctx.Request.Form.Encode())*/
	return req, res
}

func newST_Request(ctx *gin.Context) *ST_Request {
	retval := &ST_Request{
		ctx: ctx,
	}
	return retval
}
func (this *ST_Request) Body() []byte {
	body := make([]byte, 0)
	_, err := this.ctx.Request.Body.Read(body)
	if err != nil {
		wlog.Error(err.Error())
	}
	return body
}
func (this *ST_Request) String() string {
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		return this.ctx.Request.URL.Query().Encode()
	}
	return this.ctx.Request.URL.Query().Encode()
}
func (this *ST_Request) GetString(key string) string {
	var str string
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		str = this.ctx.Query(key)
	} else {
		str = this.ctx.PostForm(key)
	}
	return str
}
func (this *ST_Request) GetUint64(key string) uint64 {
	var str string
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		str = this.ctx.Query(key)
	} else {
		str = this.ctx.PostForm(key)
	}
	return wlog.StringToUint64(str)
}
func (this *ST_Request) GetInt32(key string) int32 {
	var str string
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		str = this.ctx.Query(key)
	} else {
		str = this.ctx.PostForm(key)
	}
	return wlog.StringToInt32(str)
}
func (this *ST_Request) GetInt64(key string) int64 {
	var str string
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		str = this.ctx.Query(key)
	} else {
		str = this.ctx.PostForm(key)
	}
	return wlog.StringToInt64(str)
}
func (this *ST_Request) GetUint32(key string) uint32 {
	var str string
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		str = this.ctx.Query(key)
	} else {
		str = this.ctx.PostForm(key)
	}
	return wlog.StringToUint32(str)
}
func (this *ST_Request) Set(key, value string) {
	if this.ctx.Request.Method == "" || this.ctx.Request.Method == "GET" {
		this.ctx.Set(key, value)
	} else {
		this.ctx.Set(key, value)
	}
}

type Respon struct {
	ctx  *gin.Context
	data map[string]interface{}
}

func newRespon(ctx *gin.Context) *Respon {
	retval := &Respon{
		ctx:  ctx,
		data: make(map[string]interface{}),
	}
	return retval
}

func (this *Respon) Add(key string, value interface{}) {
	if this.data == nil {
		return
	}
	if v, ok := value.(proto.Message); ok {
		this.data[key],_ = proto.Marshal(v)
	} else {
		this.data[key] = wlog.ConvertToBytes(value)
	}
}
func (this *Respon) AddMap(in map[string]interface{}) {
	if this.data == nil {
		return
	}
	for k, v := range in {
		if v1, ok := v.(proto.Message); ok {
			this.data[k],_ = proto.Marshal(v1)
		} else {
			this.data[k] = wlog.ConvertToBytes(v)
		}
	}
}

func (this *Respon) enCode() []byte {
	//构建响应
	form := &ST_ResponForm_PB{}
	for k, v := range this.data {
		pair := &ST_ResponPair_PB{
			Key:   *proto.String(k),
			Value: v.([]byte),
			//RepeatTimes: SetUint32(v.RepeatTimes),
		}
		form.ValueList = append(form.ValueList, pair)
	}

	bytesArray,_ := proto.Marshal(form)
	encodeArray, _ := GzipEncode(bytesArray)
	if len(encodeArray) > 500 {
		tmpMap := make(map[string]int, 0)
		for _, v := range form.ValueList {
			tmpMap[v.GetKey()] = len(v.GetValue())
		}
		wlog.Info(this.ctx.Request.URL.Path, tmpMap)
	}
	return encodeArray
}

func GzipEncode(in []byte) ([]byte, error) {
	if IsGzipOpen && int32(len(in)) >= GzipMiniSize {
		var (
			out []byte
			err error
		)
		writer := WriterPool.Get().(*GzipWriter)
		defer func() { writer.reset(); WriterPool.Put(writer) }()
		err, out = writer.write(in)
		if err != nil {
			wlog.Error("Gzip Encode err", err.Error())
			return out, err
		}

		if len(out) == 0 {
			return out, err
		}

		BitHeaderOn := make([]byte, 4+len(out))
		binary.BigEndian.PutUint32(BitHeaderOn, WindPlayGzipON)
		copy(BitHeaderOn[4:], out)
		return BitHeaderOn, nil

		/*
			var (
				buffer bytes.Buffer
				out []byte
				err error
			)
			writer := gzip.NewWriter(&buffer)

			_, err=writer.Write(in)
			if err != nil {
				writer.Close()
				return out, err
			}
			err = writer.Close()
			if err != nil {
				return out, err
			}
			BitHeaderOn := make([]byte, 4)
			binary.BigEndian.PutUint32(BitHeaderOn, WindPlayGzipON)
			return append(BitHeaderOn, buffer.Bytes()...), nil
		*/
	} else {
		BitHeaderOff := make([]byte, 4)
		binary.BigEndian.PutUint32(BitHeaderOff, WindPlayGzipOff)
		return append(BitHeaderOff, in...), nil
	}
}

func GzipDecode(in []byte) ([]byte, error) {
	if len(in) < 4 {
		return in, nil
	}
	header := binary.BigEndian.Uint32(in)
	if header == WindPlayGzipON {
		reader, err := gzip.NewReader(bytes.NewReader(in[4:]))
		if err != nil {
			var out []byte
			return out, err
		}
		defer reader.Close()
		return ioutil.ReadAll(reader)
	} else if header == WindPlayGzipOff {
		return in[4:], nil
	} else {
		return in, nil
	}
}

func (this *Respon) Send() {
	this.ctx.Writer.Write(this.enCode())
}
func (this *Respon) SendBody(body []byte) {
	this.ctx.Writer.Write(body)
}