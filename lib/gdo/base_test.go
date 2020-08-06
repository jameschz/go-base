package gdo

import (
	"go-base/lib/util"
	"testing"
	"time"
)

func TestD(t *testing.T) {
	util.Dump(D("demo"))
}

func TestC(t *testing.T) {
	util.Dump(C("user"))
	util.Dump(C("user"))
	util.Dump(C("log"))
	util.Dump(C("log"))
}

func TestSql(t *testing.T) {
	res, err := D("demo").T("story").Select("*", "1=1")
	util.Dump(res, err)
}

func TestTx(t *testing.T) {
	// db init
	db := D("demo")
	defer db.Close()
	// tx begin
	db.Begin()
	// tx select
	id, err := db.T("story").Max("id")
	if err != nil {
		util.Dump(err)
	} else {
		util.Dump("before insert", id)
	}
	// tx insert
	id, err = db.T("story").Insert("title=?,content=?,dtime=?", "title N", "content N", time.Now().Unix())
	if err != nil {
		util.Dump("insert fail", id, err)
		db.Rollback()
	}
	// tx select
	id, err = db.T("story").Max("id")
	if err != nil {
		util.Dump(err)
	} else {
		util.Dump("before commit", id)
	}
	// tx commit
	util.Dump("insert ok", id)
	db.Commit()
	// tx select
	id, err = db.T("story").Max("id")
	if err != nil {
		util.Dump(err)
	} else {
		util.Dump("after commit", id)
	}
}

func TestTxNest(t *testing.T) {
	// db init
	db1 := D("demo")
	defer db1.Close()
	// tx1 begin
	db1.Begin()
	util.Dump("tx1 begin")
	// call tx2 logic
	res := func() bool {
		// tx2 begin
		db2 := D("demo")
		defer db2.Close()
		db2.Begin()
		util.Dump("tx2 begin")
		// tx2 commit & rollback
		id, err := db2.T("story").Insert("title=?,content=?,dtime=?", "title N", "content N", time.Now().Unix())
		if err != nil {
			util.Dump("tx2 rollback")
			db2.Rollback()
			return false
		} else {
			util.Dump("tx2 commit", id)
			db2.Commit()
			return true
		}
	}()
	// tx1 logic
	if res {
		// tx1 commit & rollback
		id, err := db1.T("story").Insert("title=?,content=?,dtime=?", "title N", "content N", time.Now().Unix())
		if err != nil {
			util.Dump("tx1 rollback")
			db1.Rollback()
		} else {
			util.Dump("tx1 commit", id)
			db1.Commit()
		}
	}
}

func TestShard(t *testing.T) {
	// db init
	db := C("user")
	defer db.Close()
	// insert
	id, err := db.T("user_info").Insert("id=?,name=?,dtime=?", 101, "james", time.Now().Unix())
	if err != nil {
		util.Dump("insert fail", id, err)
	}
	// select by cluster
	res, err := db.T("user_info").Select("*", "id=?", 101)
	util.Dump(res, err)
	// select by specify
	res, err = D("user_shard_1").T("user_info").Select("*", "1=1")
	util.Dump(res, err)
}

func TestShardTx(t *testing.T) {
	// db init
	db := C("user")
	defer db.Close()
	// tx begin
	db.Begin()
	// tx insert
	id, err := db.T("user_info").Insert("id=?,name=?,dtime=?", 106, "james", time.Now().Unix())
	if err != nil {
		util.Dump("insert fail", id, err)
		db.Rollback()
	}
	// tx commit
	util.Dump("insert ok", id)
	db.Commit()
}
