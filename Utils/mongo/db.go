package mongo

import (
	"Utils/wlog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	sess *mgo.Session
	db   *mgo.Database
}

func NewMgoSession(collectString, DBName string) *Session {
	session, err := mgo.Dial(collectString)
	if err != nil {
		wlog.Error(err)
		panic("Connet To " + collectString + " failed!!!")
	}

	return &Session{session, session.DB(DBName)}
}
func (object *Session) Close() {
	object.sess.Close()
}

//获取集合
func (object *Session) withCollection(collectionName string, callback func(collection *mgo.Collection)) {
	callback(object.db.C(collectionName))
}
func (object *Session) UpdateByID(collectionName string, id uint64, update interface{}) error {
	var ret error = nil
	object.withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.UpdateId(id, update); nil != ret {
			wlog.Error(ret.Error(), collectionName, id, update)
		}
	})
	return ret
}

//通过id更新或插入
func (object *Session) UpsertByID(collectionName string, id, update interface{}) error {
	var ret error = nil
	object.withCollection(collectionName, func(collection *mgo.Collection) {
		if _, ret = collection.UpsertId(id, update); nil != ret {
			wlog.Error(ret.Error(), collectionName, id, update)
		}
	})
	return ret
}

//删除文档
func (object *Session) RemoveByID(collectionName string, id interface{}) error {
	var ret error = nil
	object.withCollection(collectionName, func(collection *mgo.Collection) {
		if ret = collection.RemoveId(id); nil != ret {
			wlog.Error(ret.Error(), collectionName, id)
		}
	})
	return ret
}

//获取迭代器
func (object *Session) Foreach(collectionName string, query interface{}, sort string, doc interface{}, callback func(error, bool, interface{})) {
	object.withCollection(collectionName, func(collection *mgo.Collection) {
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
