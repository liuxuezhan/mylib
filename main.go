package main

import (
	"Utils/db"
	"Utils/signal"
	"Utils/wlog"
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	Exit    = make(chan bool, 1)
	wg_exit sync.WaitGroup
	srv     = &http.Server{Addr: ":8081"}
	csv_data = map[string]*map[string][]string{}
	old      = map[string]int32{}
	path     = flag.String("dir", "./", "filepath")

)
func main() {
	//sh_sz()
		open_chrome()
	//proto_micro.LoginServer()
	signal.Init(exit, reload)
}
func open_chrome() {
	wlog.Info("")
	ctx := context.Background()
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.ExecPath("C:\\Users\\Administrator\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context

	ctx, cancel := chromedp.NewContext(c)
	defer cancel()
	var site, res string
	err := chromedp.Run(ctx,googleSearch("site:sohu.com", "Easy Money Management", &site, &res))
	if err != nil {
		wlog.Error(err)
	}

}
func googleSearch(q, text string, site, res *string) chromedp.Tasks {
	var buf []byte
	sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.sohu.com`),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#hplogo`, chromedp.ByID),
		chromedp.SendKeys(`#lst-ib`, q+"\n", chromedp.ByID),
		chromedp.WaitVisible(`#res`, chromedp.ByID),
		chromedp.Text(sel, res),
		chromedp.Click(sel),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#footer`, chromedp.ByQuery),
		chromedp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, chromedp.ByQuery),
		chromedp.Location(site),
		chromedp.Screenshot(`#testimonials`, &buf, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			return ioutil.WriteFile("testimonials.png", buf, 0644)
		}),
	}
}
func VisitWeb(url string, cookies ...string) chromedp.Tasks {// 获取登录账号后的页面
	//创建一个chrome任务
	return chromedp.Tasks{
		//ActionFunc是一个适配器，允许使用普通函数作为操作。
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置Cookie存活时间
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// 添加Cookie到chrome
			for i := 0; i < len(cookies); i += 2 {
				//SetCookie使用给定的cookie数据设置一个cookie； 如果存在，可能会覆盖等效的cookie。
				success, err := network.SetCookie(cookies[i], cookies[i+1]).
					// 设置cookie到期时间
					WithExpires(&expr).
					// 设置cookie作用的站点
					WithDomain("dl.xzg01.com:83"). //访问网站主体
					// 设置httponly,防止XSS攻击
					WithHTTPOnly(true).
					//Do根据提供的上下文执行Network.setCookie。
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", cookies[i], cookies[i+1])
				}
			}
			return nil
		}),
		// 跳转指定的url地址
		chromedp.Navigate(url),
	}
}
func DoCrawler(res *string) chromedp.Tasks {//执行爬虫任务
	return chromedp.Tasks{
		//下面注释掉的 Navigate 不要随便添加，如果添加上每次执行都相当于刷新，这样就永远翻不了页
		//chromedp.Navigate("http://dl.xzg01.com:83/OpRoot/MemberScoreList.aspx?uid=0&op=0&uname=003008"),
		chromedp.Sleep(1000),                // 等待
		chromedp.WaitVisible(`#form1`, chromedp.ByQuery), //等待id=from1页面可见  ByQuery是使用DOM选择器查找
		chromedp.Sleep(2*time.Second),
		// Click 是元素查询操作，它将鼠标单击事件发送到与选择器匹配的第一个元素节点。
		chromedp.Click(`.pagination li:nth-last-child(4) a`, chromedp.ByQuery), //点击翻页
		chromedp.OuterHTML(`tbody`, res, chromedp.ByQuery),        //获取tbody标签的html
	}
}
func WirteText(savefile string,txt string) {
	f, err := os.OpenFile(savefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString(txt + "\n")
	bw.Flush()
}
func IsMobile(text string) bool {//是否是手机号
	match,_:=regexp.MatchString(`^((\+86)|(86))?(1(([35][0-9])|[8][0-9]|[7][01356789]|[4][579]|[6][2567]))\d{8}$`,text)
	return  match
}
func GetAccount(text string) {//选择网页数据
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(strings.Replace(text, "tbody", "table", -1)))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find("tr").Each(func(i int, selection *goquery.Selection) {
		s:= selection.Find("td").Eq(6).Text()
		fmt.Println(s)
		WirteText("Acount.txt",s)
		if IsMobile(s) {
			WirteText("Mobile.txt",s)
		}
	})
}

func chromedp_ex() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	//执行任务
	url := "http://dl.xzg01.com:83/OpRoot/MemberScoreList.aspx?uid=0&op=0&uname=003008"
	err:= chromedp.Run(ctx, VisitWeb(url,
		"ASP.NET_SessionId", "zkamxkic4oiuwyc5obzgl2oj",
		"__cfduid", "d04d769b567cbe9e6f24369423b440f0d1575981989",
		"security_session_verify", "af027d69fbfbf4c925819043a50740b5",
	))
	if err != nil {
		log.Fatal(err)
	}
	var res string
	for i := 1; i < 27170; i++ {
		//执行
		err = chromedp.Run(ctx, DoCrawler(&res)) //执行爬虫任务
		if err != nil {
			log.Fatal(err)
		}
		GetAccount(res)
	}
}



func sh_sz() {
	flag.Parse()
	name := *path + "sh_sz.toml"
	conf := map[string][]string{}
	db.TomlDecode(name, &conf)
	for _, v := range conf["code"] {
		go Gethttp(v)
	}
	router := gin.Default()
	router.LoadHTMLGlob("page.html")
	router.GET("/", page)
	router.Run(":8123")
	signal.Init(exit, reload)
}


func Gethttp(code string) {
	wg_exit.Add(1)
	filename := *path + code + ".toml"
	csv_data[code] = &map[string][]string{}
	db.TomlDecode(filename, csv_data[code])
loop:
	for {
		select {
		case <-Exit:
			break loop
		case <-time.After(1 * time.Second):
			now := wlog.GetNowTime(wlog.GetNow(),8*60*60)
			if now.Hour() < 9 || now.Hour() > 15 {
				continue
			}
			d := []string{}
			res := Gethttp_One("http://qt.gtimg.cn/q="+code, "")
			if res == "" {
				continue
			}
			tmp1 := strings.Split(res, "=")
			tmp1 = strings.Split(tmp1[1], "\"")
			tmp := map[string][]string{code: strings.Split(tmp1[1], "~")}
			if len(tmp[code]) < 36 {
				continue
			}

			d = append(d, tmp[code][2])
			d = append(d, now.Format("2006/01/02 15:04:05"))
			dian := tmp[code][32]
			if wlog.StringToFloat64(dian) < 2 && wlog.StringToFloat64(dian) > -2 {
				continue
			}
			d = append(d, dian)

			in := wlog.StringToFloat64(tmp[code][9]) * wlog.StringToFloat64(tmp[code][10]) * 100
			out := wlog.StringToFloat64(tmp[code][19]) * wlog.StringToFloat64(tmp[code][20]) * 100

			t := strings.Split(tmp[code][35], "/")
			if len(t) < 3 {
				fmt.Println(3)
				continue
			}

			cur := wlog.StringToInt32(t[2])
			n := cur - old[d[0]]
			d = append(d, t[0])
			d = append(d, wlog.Int32ToString(n))
			d = append(d, wlog.Float64ToString(in))
			d = append(d, wlog.Float64ToString(out))
			if _, ok := old[d[0]]; !ok {
				old[d[0]] = cur
				(*csv_data[code])[d[1]] = d
				wlog.Info(d)
			} else if n != 0 {
				old[d[0]] = cur
				(*csv_data[code])[d[1]] = d
				wlog.Info(d)
			} else {

			}
			db.TomlEncode(filename, csv_data[code])
		}
	}
	fmt.Println("exit:", code)
	wg_exit.Done()
}
func Gethttp_One(Url, Body string) string {
	wlog.Info(Url + "?" + Body)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/x-www-form-urlencoded;charset=utf-8")

	req.SetBody([]byte(Body))
	req.SetRequestURI(Url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.DoTimeout(req, resp, time.Second*10)
	if err == nil {

		return (string(resp.Body()))
	}
	return ""
}
func exit() {
	Exit <- true
	close(Exit)
	wg_exit.Wait()
	srv.Shutdown(nil)
}
func page(ctx *gin.Context) {
	//req, _ := httputils.Prepare(ctx)
	p := charts.NewPage()
	for _, v := range csv_data {
		p.Add(
			//NewThemeRiver(v),
			bar(*v),
		)
	}
	p.Add(
		Sankey(),
	)
	f, err := os.Create("page.html")
	if err != nil {
		log.Println(err)
	}
	p.Render(f)
	ctx.HTML(http.StatusOK, "page.html", gin.H{
		"title": "hello Go",
	})
}

func bar(csv_data map[string][]string) *charts.Line {
	bar := charts.NewLine()
	data1 := make([]string, 0)
	data2 := make([]float64, 0)
	data3 := make([]float64, 0)
	data4 := make([]float64, 0)
	data5 := make([]float64, 0)
	data6 := make([]float64, 0)
	var name string
	for _, v := range csv_data {
		name = v[0]
		data1 = append(data1, v[1])
		data2 = append(data2, wlog.StringToFloat64(v[2]))
		data3 = append(data3, wlog.StringToFloat64(v[3]))
		data4 = append(data4, wlog.StringToFloat64(v[4]))
		data5 = append(data5, wlog.StringToFloat64(v[5]))
		data6 = append(data6, wlog.StringToFloat64(v[6]))
	}
	bar.AddXAxis(data1).AddYAxis("up_down", data2).AddYAxis("price", data3).AddYAxis("ok(万)", data4).
		AddYAxis("in(万)", data5).AddYAxis("out(万)", data6)
	bar.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	bar.SetGlobalOptions(charts.TitleOpts{Title: name}, charts.ToolboxOpts{Show: true})
	return bar
}
func NewThemeRiver(csv_data map[string][]string) *charts.ThemeRiver {
	trTime := [][]interface{}{}
	var name string
	for _, v := range csv_data {
		name = v[0]
		n := wlog.StringToFloat64(v[5])
		if n > 0 {
			trTime = append(trTime, []interface{}{wlog.StringToUint64(v[1]), n, "up"})
		} else {
			trTime = append(trTime, []interface{}{wlog.StringToUint64(v[1]), -n, "down"})
		}
		trTime = append(trTime, []interface{}{wlog.StringToUint64(v[1]), wlog.StringToFloat64(v[4]), "pos"})
		trTime = append(trTime, []interface{}{wlog.StringToUint64(v[1]), wlog.StringToFloat64(v[3]), "in"})
		trTime = append(trTime, []interface{}{wlog.StringToUint64(v[1]), wlog.StringToFloat64(v[2]), "out"})
	}
	tr := charts.NewThemeRiver()
	tr.SetGlobalOptions(
		charts.TitleOpts{Title: name},
		charts.SingleAxisOpts{Min: "dataMin", Max: "dataMax"},
		charts.TooltipOpts{Trigger: "axis"},
	)
	tr.Add("themeRiver", trTime)
	return tr
}
func Sankey() *charts.Sankey {
	var sankeyNode = []charts.SankeyNode{
		{Name: "category1"},
		{Name: "category2"},
		{Name: "category3"},
		{Name: "category4"},
		{Name: "category5"},
		{Name: "category6"},
	}

	var sankeyLink = []charts.SankeyLink{
		{Source: "category1", Target: "category2", Value: 10},
		{Source: "category1", Target: "category3", Value: 15},
		{Source: "category3", Target: "category4", Value: 20},
		{Source: "category4", Target: "category5", Value: 25},
	}
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(charts.TitleOpts{Title: "流量图"})
	sankey.Add("sankey", sankeyNode, sankeyLink, charts.LabelTextOpts{Show: true})
	return sankey
}
func reload() {

}
