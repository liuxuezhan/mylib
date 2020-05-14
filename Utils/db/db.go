package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"os"
	"reflect"
	"Utils/wlog"
)

//Conn truct
type Conn struct {
	db *sql.DB
}
//id
type IdDocument struct {
	Id     uint64 "_id"
	NextId uint32 "next_id"
}
//配置
type ConfigDocument struct {
	Id       uint64 "_id"
	ServerId uint32 "server_id"
}


//数据库父类
type DBObject struct {
	dbName string
	config ConfigDocument //配置
	collectionName string //集合名
}
var (
	globalSession *mgo.Session //全局会话
	serverConfig                    = "./conf/server.ini"
	DefaultUserIdCollectionName     = "user_id"
	DBList = map[string]*DBObject{}
)
const (
	panicMode                                  = false //panic模式
)
func InitDBConnection(dbName,mongodbConnectionString string,table []string){
	var err error
	globalSession, err = mgo.Dial(mongodbConnectionString)
	if nil != err {
		panic(err)
	}
	db := globalSession.DB(dbName)
	names, err := db.CollectionNames()
	if nil != err {
		panic(err)
	}
	nameBooleanMap := make(map[string]bool)
	for _, name := range names {
		nameBooleanMap[name] = true
	}

	names = []string{DefaultUserIdCollectionName}
	for _, name := range names {
		if _, ok := nameBooleanMap[name]; !ok {
			doc := &IdDocument{Id: 0, NextId: 1}
			db.C(name).Insert(doc)
		}
	}
	for _, one := range table {
		indexes := []mgo.Index{{
			Key:        []string{"_id", "read"},
			Background: true,
		}, {
			Key:        []string{"name"},
			Background: true,
			Unique:     true,
			Sparse:     true,
		}, {
			Key:        []string{"google_name"},
			Background: true,
			Unique:     true,
			Sparse:     true,
		}, {
			Key:        []string{"fb_name"},
			Background: true,
			Unique:     true,
			Sparse:     true,
		}, {
			Key:        []string{"twitter_name"},
			Background: true,
			Unique:     true,
			Sparse:     true,
		}, {
			Key:        []string{"device_id"},
			Background: true,
		}, {
			Key:        []string{"gamecenter_name"},
			Background: true,
			Unique:     true,
			Sparse:     true,
		}, {
			Key:        []string{"accountid"},
			Background: true,
			Sparse:     true,
		}}
		object:=&DBObject{collectionName:one }
		DBList[one]=object
		EnsureIndex(object.collectionName, indexes)
	}
}

func Decode(fpath string, v interface{}) {
	if _, err := toml.DecodeFile(fpath, v); err != nil {
	}
}
func Encode(fpath string, v interface{}) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(v); err != nil {
	}
	if err := ioutil.WriteFile(fpath, buf.Bytes(), os.ModeType); err != nil {
	}
}
func TomlDecode(fpath string, v interface{}) {
	if _, err := toml.DecodeFile(fpath, v); err != nil {
		wlog.Error(err.Error())
		panic("")
	}
}
func TomlEncode(fpath string, v interface{}) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(v); err != nil {
	}
	if err := ioutil.WriteFile(fpath, buf.Bytes(), os.ModeType); err != nil {
	}
}

//Connect Connect to database
func (conn *Conn) Connect(dbhost string, dbport string, database string, user string, password string) bool {
	var err error
	conn.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbhost, dbport, database))
	if err != nil || dbhost == "" || dbport == "" || database == "" {
		return false
	}
	return true
}

//Close Close Connection
func (conn *Conn) Close() {
	conn.db.Close()
}

//Excute the insert and update sql
func (conn *Conn) Excute(sql string, args ...interface{}) (sql.Result, bool) {
	if conn == nil {
		return nil, false
	}
	v, err := conn.db.Exec(sql, args...)
	if err != nil {
		return nil, false
	}
	return v, true
}

//Query 执行查询语句，返回结果
func (conn *Conn) Query(sql string, args ...interface{}) ([]map[string]string, bool) {
	var reval bool
	if conn == nil {
		return make([]map[string]string, 0), false
	}
	rows, err := conn.db.Query(sql, args...)
	if err != nil || rows == nil {
		fmt.Println(sql, args)
		reval = false
	} else {
		reval = true
	}
	if !reval {
		return make([]map[string]string, 0), reval
	}
	return parseResult(rows), reval
}

//QueryNum return num
func (conn *Conn) QueryNum(sql string, args ...interface{}) int64 {
	if conn == nil {
		return 0
	}
	var reval int64
	rows, err := conn.db.Query(sql, args...)
	if err != nil || rows == nil {
		return 0
	}
	for rows.Next() {
		rows.Scan(&reval)
	}
	rows.Close()
	return reval
}

//QueryStringOne return string result
func (conn *Conn) QueryStringOne(sql string, args ...interface{}) string {
	if conn == nil {
		return ""
	}
	var reval string
	rows, err := conn.db.Query(sql, args...)
	if err != nil || rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&reval)
	}
	rows.Close()
	return reval
}

func parseResult(rows *sql.Rows) []map[string]string {
	result := make([]map[string]string, 0)
	if rows == nil {
		return result
	}
	column, err := rows.Columns()
	if err == nil {
		for rows.Next() {
			columnsMp := make(map[string]interface{}, len(column))
			tmpresult := make(map[string]string, len(column))
			refs := make([]interface{}, 0, len(column))
			for _, col := range column {
				var ref interface{}
				columnsMp[col] = &ref
				refs = append(refs, &ref)
			}

			rows.Scan(refs...)
			for k, v := range columnsMp {
				value := reflect.ValueOf(v).Elem().Interface()
				switch value.(type) {
				case int64:
					tmpresult[k] = fmt.Sprintf("%d", value.(int64))
				case string:
					tmpresult[k] = value.(string)
				case int:
					tmpresult[k] = fmt.Sprintf("%d", value.(int))
				case int32:
					tmpresult[k] = fmt.Sprintf("%d", value.(int32))
				case nil:
					tmpresult[k] = ""
				default:
					tmpresult[k] = string(value.([]uint8))
				}

			}
			result = append(result, tmpresult)
		}
		rows.Close()
	}
	return result
}
////>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
//获取集合
func withCollection(collectionName string, callback func(collection *mgo.Collection)) {
	object := DBList[collectionName]
	sessionCopy := globalSession.Clone()
	defer sessionCopy.Close()
	callback(sessionCopy.DB(object.dbName).C(collectionName))
}

//创建索引
func  EnsureIndex(collectionName string, indexes []mgo.Index) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		for _, index := range indexes {
			if ret = collection.EnsureIndex(index); nil != ret {
				if panicMode {
					panic(ret)
				} else {
					wlog.Error(fmt.Sprintf("(%v, %v) ensure index error: %v", collectionName, index, ret.Error()))
				}
				break
			}

		}
	})
	return ret
}
//获取
func Get(collectionName string, query, res interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.Find(query).One(res); nil != ret {
			if panicMode {
				panic(ret)
			} else if ret != mgo.ErrNotFound {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) get error: %v", collectionName, query, res, ret.Error()))
			}
		}
	})
	return ret
}

//获取
func GetMany(collectionName, sortkey string, query, res interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.Find(query).Sort(sortkey).All(res); nil != ret {
			if panicMode {
				panic(ret)
			} else if ret != mgo.ErrNotFound {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) get error: %v", collectionName, query, res, ret.Error()))
			}
		}
	})
	return ret
}

//获取所有查询结果
func  GetById(collectionName string, id, res interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.FindId(id).One(res); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) get by id error: %v", collectionName, id, res, ret.Error()))
			}
		}
	})
	return ret
}

//获取所有
func  GetAll(collectionName, sort string, skip, limit int, query, selector, res interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if nil == query {
			query = bson.M{}
		}
		q := collection.Find(query).Select(selector)
		if 0 != len(sort) {
			q = q.Sort(sort)
		}
		if ret = q.Skip(skip).Limit(limit).All(res); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v, %v, %v, %v, %v) get all error: %v",
					collectionName, sort, skip, limit, query, selector, res, ret.Error()))
			}
		}
	})
	return ret
}

//计数
func  Count(collectionName string, query interface{}) (error, int) {
	var count int = 0
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if nil == query {
			count, ret = collection.Count()
		} else {
			count, ret = collection.Find(query).Count()
		}
		if nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v) count error: %v", collectionName, query, ret.Error()))
			}
		}
	})
	return ret, count
}

//查找并修改
func FindAndModify(collectionName string, id interface{}, change mgo.Change, doc interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.FindId(id).Apply(change, doc); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v, %v) find and modify error: %v", collectionName, id, change, doc, ret.Error()))
			}
		}
	})
	return ret
}

//更新
func  Update(collectionName string, query, update interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.Update(query, update); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) update error: %v", collectionName, query, update, ret.Error()))
			}
		}
	})
	return ret
}

//更新
func UpdateById(collectionName string, id uint64, update interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.UpdateId(id, update); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) update error: %v", collectionName, id, update, ret))
			}
		}
	})
	return ret
}

//更新所有
func UpdateAll(collectionName string, query, update interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.UpdateAll(query, update); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) update all error: %v", collectionName, query, update, ret.Error()))
			}
		}
	})
	return ret
}

//通过id更新或插入
func UpsertById(collectionName string, id, update interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.UpsertId(id, update); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) upsert by id error: %v", collectionName, id, update, ret.Error()))
			}
		}
	})
	return ret
}

//更新或插入
func upsert(collectionName string, query, update interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.Upsert(query, update); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v, %v) upsert error: %v", collectionName, query, update, ret.Error()))
			}
		}
	})
	return ret
}

//设置新的置
func  UpdateKeyValue(collectionName string, id uint64, key string, value interface{}) error {
	update := bson.M{"$set": bson.M{key: value}}
	return UpdateById(collectionName, id, update)
}

//删除文档
func RemoveById(collectionName string, id interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.RemoveId(id); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v) remove by id error: %v", collectionName, id, ret.Error()))
			}
		}
	})
	return ret
}

//删除文档
func RemoveAll(collectionName string, query interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.RemoveAll(query); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v) remove all error: %v", collectionName, query, ret.Error()))
			}
		}
	})
	return ret
}

//删除文档
func  Remove(collectionName string, query interface{}) error {
	var ret error = nil
	withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.Remove(query); nil != ret {
			if panicMode {
				panic(ret)
			} else {
				wlog.Error(fmt.Sprintf("(%v, %v) remove error: %v", collectionName, query, ret.Error()))
			}
		}
	})
	return ret
}

//获取迭代器
func Foreach(collectionName string, query interface{}, sort string, doc interface{}, callback func(error, bool, interface{})) {
	withCollection(collectionName, func(collection *mgo.Collection) {
		if nil == query {
			query = bson.M{}
		}
		q := collection.Find(query)
		if nil == q {
			return
		}
		if 0 != len(sort) {
			q = q.Sort(sort)
		}
		iterator := q.Iter()
		if nil == iterator {
			return
		}
		for iterator.Next(doc) {
			callback(nil, true, doc)
		}
		iterator.Close()
	})
}