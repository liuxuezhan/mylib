package wlog

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/futurenda/google-auth-id-token-verifier"
	"github.com/golang/glog"
	"github.com/kyle-wang/GameCenter"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
	random "math/rand"
	"errors"
	"gopkg.in/square/go-jose.v1/json"
)

const (
	log_info uint32 = 3
	log_warn uint32 = 2
	log_err  uint32 = 1
)

var (
	loglevel          uint32 = 3
	tostd             bool   = true
	Count_handler     sync.Map
	Count_handler_num sync.Map
)

//Info Info
func Info(args ...interface{}) {
	InfoDepth(1, args...)
}

func InfoDepth(depth int, args ...interface{}) {
	if loglevel >= log_info {
		if tostd {
			_, file, line, _ := runtime.Caller(1 + depth)
			if i := strings.LastIndex(file, "/"); i >= 0 {
				file = file[i+1:]
			}

			fmt.Println("INFO ", time.Unix(int64(GetNow()), 0).Format(time.RFC3339), file, ":", line, "] ", args)
		}
		glog.InfoDepth(1+depth, args...)
	}
}

//Error Error
func Error(args ...interface{}) {
	ErrorDepth(1, args...)
}
func ErrorDepth(depth int, args ...interface{}) {
	if loglevel >= log_warn {
		if tostd {
			_, file, line, _ := runtime.Caller(1 + depth)
			if i := strings.LastIndex(file, "/"); i >= 0 {
				file = file[i+1:]
			}
			fmt.Println("ERROR ", time.Unix(int64(GetNow()), 0).Format(time.RFC3339), file, ":", line, "] ", args)

		}
		glog.ErrorDepth(1+depth, args...)
	}
}

//Warning Warning
func Warning(args ...interface{}) {
	if loglevel >= log_err {
		if tostd {
			_, file, line, _ := runtime.Caller(1)
			if i := strings.LastIndex(file, "/"); i >= 0 {
				file = file[i+1:]
			}
			fmt.Println("WARN ", time.Unix(int64(GetNow()), 0).Format(time.RFC3339), file, ":", line, "] ", args)
		}
		glog.WarningDepth(1, args...)
	}
}

func FlushNow() {
	glog.Flush()
}

func flush() {
	for {
		glog.Flush()
		time.Sleep(time.Duration(10) * time.Second)
	}

}
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

}
func Initialize(logLevel uint32, toSTD bool,serverid int64) {
	loglevel = logLevel
	tostd = toSTD
	//创建日志目录
	if !IsDirExists("log") {
		if err := os.Mkdir("log", os.ModeDir); nil != err {
			fmt.Fprintf(os.Stderr, "create log dir failed: %v\n", err.Error())
		}
	}
	flag.Set("log_dir", "log")
	flag.Parse()
	//go flush()

	InitSnowFlake(serverid)
}
var (
	Interval int64 = 0
)

func GetNow() int64 {
	return time.Now().Unix() + Interval
}
func GetNowTime(tm int64,offset int) time.Time {
	now := time.Unix(int64(tm), 0).In(time.FixedZone("CST", offset))
	return now
}
func GetTodayZeroClock(tm int64,offset int) int64 {
	now := GetNowTime(tm,offset)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("CST", offset))
	return today.Unix()
}
func ParseTimeToUnixTime(in string,offset int ) int64 {
	timeT, _ := time.ParseInLocation("2006-01-02 15:04:05", in,time.FixedZone("CST", offset))
	Info(timeT.Hour())
	return timeT.Unix()
}

func GetNextMonday(tm int64,offset int) int64 { //周一0点
	now := GetNowTime(tm,offset)
	next := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("CST", offset))
	if now.Weekday() == time.Sunday {
		next = next.AddDate(0, 0, 1)
	} else {
		offset := int(time.Saturday - now.Weekday())
		next = next.AddDate(0, 0, offset+2)
	}
	return next.Unix()
}

//返回当前时间是当年的第几周
func GetWeek(tm int64,offset int) uint32 {
	t := GetNowTime(tm,offset)
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	//fmt.Println("%d第%d周", t.Year(), week)
	return uint32(week)
}

var (
	randGenerate *random.Rand
	randomLock   sync.Mutex
	BuffPool     = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

const (
	MapLength           int32  = 990
	MapHeight           int32  = 990
	ChunkSize           int32  = 30
	MainCityLength      uint32 = 200
	MainCityHeight      uint32 = 200
	MainCitySize               = MainCityLength * MainCityHeight
	WorldMainCityWidth  int32  = 2
	WorldMainCityHeight int32  = 2
	tileConvert         int32  = 1000
)

func initRandSeek() {
	randGenerate = random.New(random.NewSource(time.Now().UnixNano()))
}

func GetValueByPR(pr, values []int32) int32 {
	if len(pr) != len(values) {
		return int32(math.MinInt32)
	}
	totalPr := ComputeInt32Slice(pr)
	shouldAdd := 10000 - totalPr
	if shouldAdd < 0 {
		return int32(math.MinInt32)
	}
	for k, v := range pr {
		pr[k] = v + shouldAdd*v/totalPr
	}
	randomNum := GetRandomNum(10000)
	for k, v := range pr {
		if randomNum < v {
			return values[k]
		}
		randomNum -= v
	}
	return values[len(values)-1]
}

func IsFileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func Int32ToString(input int32) string {
	return strconv.FormatInt(int64(input), 10)
}

func Int64ToString(input int64) string {
	return strconv.FormatInt(input, 10)
}

func Uint64ToString(input uint64) string {
	return strconv.FormatUint(input, 10)
}
func Float64ToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}
func Uint32ToString(input uint32) string {
	return strconv.FormatUint(uint64(input), 10)
}

func StringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}

func StringToInt32(input string) int32 {
	output, _ := strconv.ParseInt(input, 10, 64)
	return int32(output)
}

func BoolToString(input bool) string {
	return strconv.FormatBool(input)
}
func StringToBool(input string) bool {
	ret, _ := strconv.ParseBool(input)
	return ret
}
func StringToInt64(input string) int64 {
	output, _ := strconv.ParseInt(input, 10, 64)
	return output
}
func StringToUint64(input string) uint64 {
	output, _ := strconv.ParseUint(input, 10, 64)
	return output
}
func StringToUint32(input string) uint32 {
	output, _ := strconv.ParseUint(input, 10, 64)
	return uint32(output)
}
func StringToFloat64(input string) float64 {
	ret, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0
	}
	return ret
}

func CombineMap(m1 ...map[string]string) map[string]string {
	tmpMap := make(map[string]string, 0)
	for i := 0; i < len(m1); i++ {
		for k, v := range m1[i] {
			tmpMap[k] = v
		}
	}
	return tmpMap
}
func CopyUint64Slice(input []uint64) []uint64 {
	output := make([]uint64, len(input))
	copy(output, input)
	return output
}
func CopyUint64Map(in map[uint64]uint64) map[uint64]uint64 {
	ret := make(map[uint64]uint64, len(in))
	for k, v := range in {
		ret[k] = v
	}
	return ret
}



func ParseStringToMap(input string) (map[uint64]uint64, map[uint64]uint64) {
	ret := make(map[uint64]uint64, 0)
	ret2 := make(map[uint64]uint64, 0)
	tmp := strings.Split(input, "|")
	for _, v := range tmp {
		t_tmp := strings.Split(v, " ")
		if len(t_tmp) < 2 {
			continue
		}
		ret[StringToUint64(t_tmp[0])] = StringToUint64(t_tmp[1])
		ret2[StringToUint64(t_tmp[1])] = StringToUint64(t_tmp[0])
	}
	return ret, ret2
}

func ParseStringToUint32List(input string) []uint32 {
	ret := make([]uint32, 0)
	tmp := strings.Split(input, "|")
	for _, v := range tmp {
		ret = append(ret, StringToUint32(v))
	}
	return ret
}

func ParseStringToUint64List(input string) []uint64 {
	ret := make([]uint64, 0)
	tmp := strings.Split(input, "|")
	for _, v := range tmp {
		ret = append(ret, StringToUint64(v))
	}
	return ret
}
func ParseStringToFloat64List(input string) []float64 {
	ret := make([]float64, 0)
	tmp := strings.Split(input, "|")
	for _, v := range tmp {
		ret = append(ret, StringToFloat64(v))
	}
	return ret
}
func ParseStringToInt64List(input string) []int64 {
	ret := make([]int64, 0)
	tmp := strings.Split(input, "|")
	for _, v := range tmp {
		ret = append(ret, StringToInt64(v))
	}
	return ret
}



func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GenGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func Base64Decode(in string) string {
	tmp, _ := base64.StdEncoding.DecodeString(in)
	return string(tmp)
}

func Base64Encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func ConverToInt64(in interface{}) int64 {
	var retval int64
	switch in.(type) {
	case int64:
		retval = in.(int64)
	case int32:
		retval = int64(in.(int32))
	case int:
		retval = int64(in.(int))
	case float32:
		retval = int64(math.Floor(float64(in.(float32))))
	case float64:
		retval = int64(math.Floor(in.(float64)))
	default:
		retval = 0
	}
	return retval
}
func ConverToUint64(in interface{}) uint64 {
	var retval uint64
	switch in.(type) {
	case uint64:
		retval = in.(uint64)
	case uint32:
		retval = uint64(in.(uint32))
	case int64:
		retval = uint64(in.(int64))
	case int32:
		retval = uint64(in.(int32))
	case int:
		retval = uint64(in.(int))
	case float32:
		retval = uint64(math.Floor(float64(in.(float32))))
	case float64:
		retval = uint64(math.Floor(in.(float64)))
	default:
		retval = 0
	}
	return retval
}
func ConverToUint32(in interface{}) uint32 {
	var retval uint32
	switch in.(type) {
	case uint64:
		retval = uint32(in.(uint64))
	case uint32:
		retval = in.(uint32)
	case int64:
		retval = uint32(in.(int64))
	case int32:
		retval = uint32(in.(int32))
	case int:
		retval = uint32(in.(int))
	case float32:
		retval = uint32(math.Floor(float64(in.(float32))))
	case float64:
		retval = uint32(math.Floor(in.(float64)))
	default:
		retval = 0
	}
	return retval
}

func ConverToInt32(in interface{}) int32 {
	var retval int32
	switch in.(type) {
	case int64:
		retval = int32(in.(int64))
	case int32:
		retval = int32(in.(int32))
	case int:
		retval = int32(in.(int))
	case float32:
		retval = int32(math.Floor(float64(in.(float32))))
	case float64:
		retval = int32(math.Floor(in.(float64)))
	default:
		retval = 0
	}
	return retval
}

func ConvertToString(in interface{}) string {
	var retval string
	switch in.(type) {
	case int64:
		retval = fmt.Sprintf("%d", in.(int64))
	case string:
		retval = in.(string)
	case int:
		retval = fmt.Sprintf("%d", in.(int))
	case int32:
		retval = fmt.Sprintf("%d", in.(int32))
	case float32:
		retval = fmt.Sprintf("%d", int64(in.(float32)))
	case float64:
		retval = fmt.Sprintf("%d", int64(in.(float64)))
	case []byte:
		retval = string(in.([]byte))
	case uint64:
		retval = fmt.Sprintf("%d", in.(uint64))
	case uint32:
		retval = fmt.Sprintf("%d", in.(uint32))
	default:
		retval = ""
	}
	return retval
}

func ConvertToBytes(in interface{}) []byte {
	if v, ok := in.(string); ok {
		return []byte(v)
	}

	buf := BuffPool.Get().(*bytes.Buffer)
	//buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, in); nil != err {
		fmt.Errorf("write error: %s", err.Error())
	}
	data := make([]byte, buf.Len())
	copy(data, buf.Bytes())
	buf.Reset()
	BuffPool.Put(buf)
	return data
}
func AddVersionInfoBeforeData(data []byte, version int32) []byte {
	return append(ConvertToBytes(version), data...)
}

func ConvertCoordinateToId(x int32, y int32, sid int32) int32 {
	if x > MapLength || y > MapHeight || sid < 0 {
		return -1
	}
	id := y*tileConvert + x
	id = sid<<20 | id
	return id
}

func CheckTileidValid(id int32) bool {
	x, y, _ := ConvertIdToCoordinate(id)
	if x == -1 || y == -1 {
		return false
	}
	return true
}
func ConvertIdToCoordinate(id int32) (int32, int32, int32) {
	sid := id >> 20
	id = id & 0xFFFFF
	if id < 0 || id >= tileConvert*tileConvert {
		return -1, -1, -1
	}
	x, y := id%tileConvert, id/tileConvert
	if x < 0 || x >= MapLength || y < 0 || y >= MapHeight {
		return -1, -1, -1
	}
	return x, y, sid
}

func ConvertMainCityCoordinateToId(x, y uint32) uint32 {
	if x > MainCityLength || y > MainCityHeight {
		return 0
	}
	return y*MainCityLength + x
}

func ConvertMainCityIdToCoordinate(id uint32) (uint32, uint32) {
	if id < 0 || id > MainCitySize {
		return 0, 0
	}
	return id % MainCityLength, id / MainCityHeight
}

func ParseParameterStringToMap(in string) map[string]string {
	ret := make(map[string]string)
	tmp := strings.Split(in, ";")
	for _, v := range tmp {
		tmp1 := strings.Split(v, "=")
		if len(tmp1) > 1 {
			ret[tmp1[0]] = tmp1[1]
		}
	}
	return ret
}
func StatueFuncMapToString(d map[string]string) string {
	arr := make([]string, 0)
	for k, v := range d {
		arr = append(arr, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(arr, ";")
}


func GetRandomNum(in int32) int32 {
	if in <= 0 {
		return 0
	}
	randomLock.Lock()
	defer randomLock.Unlock()
	return randGenerate.Int31n(in)
}

func GetRandPerm(in int) []int {
	randomLock.Lock()
	defer randomLock.Unlock()
	return randGenerate.Perm(in)
}


func Uint64SliceContant(slice []uint64, value uint64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func StringSliceContant(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func Uint32SliceContant(slice []uint32, value uint32) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func Int32SliceContant(slice []int32, value int32) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func ComputeInt32Slice(slice []int32) int32 {
	var ret int32
	for _, v := range slice {
		ret += v
	}
	return ret
}
func ComputeUint32Slice(slice []uint32) uint32 {
	var ret uint32
	for _, v := range slice {
		ret += v
	}
	return ret
}
func BitIsSet(value uint32, position uint32) bool {
	if 32 < position {
		return false
	}

	return 0 != value&uint32(0x01<<position)
}

func SetBit(value *uint32, position uint32) {
	if 32 > position {
		*value |= uint32(0x01 << position)
	}
}

func ResetBit(value *uint32, position uint32) {
	if 32 > position {
		*value &= ^uint32(0x01 << position)
	}
}
func InsertSplitK(u64 uint64) string {
	tmpStr := Uint64ToString(u64)
	var needAdd = 0
	if len(tmpStr)%3 != 0 {
		needAdd = 3 - len(tmpStr)%3
	}
	tmp := make([]string, 0)
	for i := 0; i < needAdd; i++ {
		tmp = append(tmp, ",")
	}
	for k, v := range tmpStr {
		if (k+needAdd)%3 == 0 {
			tmp = append(tmp, ",")
		}
		tmp = append(tmp, string(v))
	}
	retStr := strings.Join(tmp, "")
	return strings.TrimLeft(retStr, ",")
}

func AddUint64ToSlice(slice []uint64, value uint64) []uint64 {
	var find = false
	for _, v := range slice {
		if v == value {
			find = true
			break
		}
	}
	if !find {
		slice = append(slice, value)
	}
	return slice
}
func DeleteUint64ToSlice(slice []uint64, value uint64) []uint64 {
	for k, v := range slice {
		if v == value {
			slice = append(slice[:k], slice[k+1:]...)
			break
		}
	}
	return slice
}
func DeleteInt64ToSlice(slice []int64, value int64) []int64 {
	for k, v := range slice {
		if v == value {
			slice = append(slice[:k], slice[k+1:]...)
			break
		}
	}
	return slice
}
func DeleteUint32FromSlice(slice []uint32, value uint32) []uint32 {
	for k, v := range slice {
		if v == value {
			slice = append(slice[:k], slice[k+1:]...)
			break
		}
	}
	return slice
}

// parse string(DATE) like 20170804 to UTC timestamp
func StringFormatGtime(dateStr string) int64 {
	strSize := len(dateStr)
	if strSize != 8 {
		return 0 //	wrong date size
	}

	y := int(StringToUint32(string(dateStr[0:4])))
	m := time.Month(StringToUint32(string(dateStr[4:6])))
	d := int(StringToUint32(string(dateStr[6:8])))
	newTime := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return newTime.Unix()
}

func TimeFormatString(duration uint64) string {
	var tmp string
	if duration/86400 > 0 {
		tmp += fmt.Sprintf("%dD", duration/86400)
		duration = duration % 86400
	}

	if len(tmp) > 0 {
		tmp += fmt.Sprintf("%dH", duration/3600)
		duration = duration % 3600
	} else {
		if duration/3600 > 0 {
			tmp += fmt.Sprintf("%dH", duration/3600)
			duration = duration % 3600
		}
	}

	if len(tmp) > 0 {
		tmp += fmt.Sprintf("%dM", duration/60)
		duration = duration % 60
	} else {
		if duration/60 > 0 {
			tmp += fmt.Sprintf("%dM", duration/60)
			duration = duration % 60
		}
	}
	if duration > 0 {
		tmp += fmt.Sprintf("%dS", duration)
	}
	return tmp
}

/**
@brief 雕像加成key-value array转换为string
*/
func StatueFuncArrayToString(keys []string, values []uint32, perValue uint32) string {
	if nil == keys || 0 == len(keys) || nil == values || 0 == len(values) {
		return ""
	}
	arr := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		arr = append(arr, fmt.Sprintf("%s=%d", keys[i], values[i]*perValue/10000))
	}
	return strings.Join(arr, ";")
}

/**
@brief 雕像加成key-value array转换为map
*/
func StatueFuncArrayToMap(keys []string, values []uint32) map[string]uint32 {
	if nil == keys || 0 == len(keys) || nil == values || 0 == len(values) {
		return nil
	}
	arr := make(map[string]uint32)
	for i := 0; i < len(keys); i++ {
		arr[keys[i]] = values[i]
	}
	return arr
}

func ProcessTime(duration uint64) string {
	var tmp string
	if duration/86400 > 0 {
		tmp += fmt.Sprintf("%dD", duration/86400)
		duration = duration % 86400
	}

	if len(tmp) > 0 {
		tmp += fmt.Sprintf("%dH", duration/3600)
		duration = duration % 3600
	} else {
		if duration/3600 > 0 {
			tmp += fmt.Sprintf("%dH", duration/3600)
			duration = duration % 3600
		}
	}

	if len(tmp) > 0 {
		tmp += fmt.Sprintf("%dM", duration/60)
		duration = duration % 60
	} else {
		if duration/60 > 0 {
			tmp += fmt.Sprintf("%dM", duration/60)
			duration = duration % 60
		}
	}

	if duration > 0 {
		tmp += fmt.Sprintf("%dS", duration)
	}
	return tmp
}

func init() {
	initRandSeek()
}

func DelNum(a, b int32) int32 {
	if a > b {
		return a - b
	} else {
		return 0
	}
}
func Max(a, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a, b int32) int32 {
	if a > b {
		return b
	} else {
		return a
	}
}
func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

func GetChunkIdx(tileId int32) int {
	x, y, _ := ConvertIdToCoordinate(tileId)
	idx := (x/ChunkSize)*(MapHeight/ChunkSize) + y/ChunkSize
	return int(idx)
}

/*
雪花算法golang实现，用于生成全局唯一ID
snowflakeID(int64) = 1bit(无用) + 41bit(时间戳:最大2^41 - 1) + 10bit(机器ID) + 12bit(序列号)
 */

var SFFactory *SnowFlake
var err error
const(
	seriaNumBits uint8 = 12
	serverIDBits uint8 = 10

	seriaNumMax int64 = -1 ^ (-1 << seriaNumBits) //最大序列号：4096
	serverIDMax int64 = -1 ^ (-1 << serverIDBits) //最大机器ID：1024

	serverIDShift  uint8 = seriaNumBits				//机器ID左偏移量
	timeStampShift uint8 = seriaNumBits + serverIDBits //时间戳左偏移量

	startTimeStamp int64 = 1559543323827				//开启算法时的时间戳(毫秒)，不允许修改
)


type SnowFlake struct {
	mu 			sync.Mutex 	//锁
	timeStamp 	int64 		//时间戳(毫秒),上一次生成ID的时间戳
	serverID 	int64	  	//服务器编号
	seriaNum 	int64 		//序列号(自增)，一毫秒最大累计至4096
}

/**
@brief 生成一个新的NewSnowFlake
*/
func newSnowFlake(serverID int64)(*SnowFlake, error){
	if 0 > serverID || serverID > serverIDMax{
		return nil, errors.New("new SnowFlakeID Failed:serverID error")
	}

	snowFlake := &SnowFlake{
		timeStamp:0,
		serverID:serverID,
		seriaNum:0,
	}

	return snowFlake, nil
}

/**
@brief 计算出唯一ID
*/
func (obj *SnowFlake)getSFUID()(int64, bool){
	obj.mu.Lock()
	defer obj.mu.Unlock()

	if nil == SFFactory{
		return 0, false
	}

	//当前时间戳(毫秒)
	nowTime := time.Now().UnixNano()/1e6

	if obj.timeStamp == nowTime{
		obj.seriaNum++

		//如果当前毫秒内序列号达到最大值
		//则等到下一毫秒再生成
		if obj.seriaNum > seriaNumMax{
			for nowTime <= obj.timeStamp    {
				nowTime = time.Now().UnixNano()/1e6
			}
		}
	}else{
		obj.seriaNum = 0
	}

	obj.timeStamp = nowTime

	//通过位移得出唯一ID
	SFUID := (nowTime - startTimeStamp) << timeStampShift | (obj.serverID << serverIDShift) | (obj.seriaNum)

	return SFUID, true
}


func InitSnowFlake(serverID int64){
	SFFactory, err = newSnowFlake(serverID)
	if nil != err{
		panic(err)
		return
	}
}

func GetSFUID()int64{
	if nil == SFFactory{
		return 0
	}

	uid,_:= SFFactory.getSFUID()

	return uid
}


/**
@brief 测试
*/
func StartTestSFUID(){
	obj , error := newSnowFlake(1)
	if nil != error{
		fmt.Println(error)
		return
	}

	count := 10
	chSFUID := make(chan int64, count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++{
		wg.Add(1)
		go func() {
			SFUID, ok := obj.getSFUID()
			if ok{
				chSFUID <- SFUID
			}
		}()
	}


	close(chSFUID)

	testMap := make(map[int64]int)

	for i := 0; i < count; i++{
		go func(){
			SFUID := <-chSFUID
			if _, ok:=testMap[SFUID]; ok{
				fmt.Println("have save SFUID...\n", SFUID)
			}
			testMap[SFUID] = i
			wg.Done()
		}()

	}
	wg.Wait()

	print("TestSFUID End\n")
}


func VerifyGoogleIDToken(idToken string,GoogleClientID string) (string) {
	var uid string
	v := googleAuthIDTokenVerifier.Verifier{}
	err := v.VerifyIDToken(idToken, []string{
		GoogleClientID,
	})
	claimSet, err2 := googleAuthIDTokenVerifier.Decode(idToken)
	if err != nil { //Google验证失败
		Error(err)
		if err2 == nil { //解析token成功
			Info(claimSet.Email, "  ", claimSet.Sub)
			return claimSet.Sub
		} else {
			return uid
		}
	} else {
		if err2 == nil {
			return claimSet.Sub
		} else {
			Error("googleAuthIDTokenVerifier.Decode Failed")
			return uid
		}
	}
}

func VerifyFBTokenID(idToken,FBAppID, FBAppSecret string) (string) {
	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)
	args.Add("access_token", fmt.Sprintf("%s|%s", FBAppID, FBAppSecret))
	args.Add("input_token", idToken)

	resp, err := http.Get(fmt.Sprintf("%s%s", "https://graph.facebook.com/debug_token?", args.String()))
	if err != nil {
		return ""
	}

	type resultData struct {
		App_id      string   `json:"app_id"`
		Type        string   `json:"type"`
		Application string   `json:"application"`
		Expires_at  uint64   `json:"expires_at"`
		Is_valid    bool     `json:"is_valid"`
		Issued_at   uint64   `json:"issued_at"`
		Scopes      []string `json:"scopes"`
		User_id     string   `json:"user_id"`
	}

	type Result struct {
		Data resultData `json:"data"`
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	result := &Result{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return ""
	}

	if !result.Data.Is_valid {
		return ""
	}

	return result.Data.User_id
}
func VerifyGameCenter(username, data, sig, puk string) (string) {
	if err := GameCenter.Verify(puk, data, sig); err == nil {
		return username
	} else {
		return ""
	}
}
func VerifyTwitterCredentials(token, secret,TTConsumerKey,TTConsumerSecret string) (string) {
	token1 := oauth1.NewToken(token, secret)
	// OAuth1 http.Client will automatically authorize Requests
	TwitterConfig := oauth1.NewConfig(TTConsumerKey,TTConsumerSecret)
	httpClient := TwitterConfig.Client(oauth1.NoContext, token1)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return ConvertToString(user.ID)
}