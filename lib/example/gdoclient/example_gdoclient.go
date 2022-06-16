package examplegdoclient

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jameschz/go-base/lib/gdo"
	gdopool "github.com/jameschz/go-base/lib/gdo/pool"
	"github.com/jameschz/go-base/lib/gutil"
)

// MysqlQueryBasic :
func MysqlQueryBasic() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// init result struct
	type story struct {
		ID      int
		Title   string
		Content string
		Dtime   string
	}
	// test FetchMaps
	rows, err := db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		r, _ := db.T("story").FetchMaps(rows)
		for k, v := range *r {
			fmt.Println("> mysql query basic : FetchMaps - id :", k)
			gutil.Dump(v)
		}
	}
	// test FetchMap
	rows, err = db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		r, _ := db.T("story").FetchMap(rows)
		fmt.Println("> mysql query basic : FetchMap")
		gutil.Dump(r)
	}
	// test FetchStructs
	rows, err = db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		s := &story{}
		r, _ := db.T("story").FetchStructs(rows, s)
		for k, v := range *r {
			fmt.Println("> mysql query basic : FetchStructs - id :", k)
			gutil.Dump(v.(story))
		}
	}
	// test FetchStruct
	rows, err = db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		s := &story{}
		r, _ := db.T("story").FetchStruct(rows, s)
		fmt.Println("> mysql query basic : FetchStruct")
		gutil.Dump(r.(story))
	}
}

// MysqlInsertBasic :
func MysqlInsertBasic() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// test insert
	id, err := db.T("story").Insert("title=?,content=?,dtime=?", "title N", "content N", time.Now().Unix())
	if err != nil {
		fmt.Println("> mysql insert basic err", err)
	} else {
		fmt.Println("> mysql insert basic id", id)
	}
}

// MysqlUpdateBasic :
func MysqlUpdateBasic() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// test max
	maxID, _ := db.T("story").Max("id")
	if maxID > 0 {
		// test update
		title := "title " + strconv.FormatInt(maxID, 10)
		content := "content " + strconv.FormatInt(maxID, 10)
		affect, err := db.T("story").Update("title=?,content=? where id=?", title, content, maxID)
		if err != nil {
			fmt.Println("> mysql update basic err", err)
		} else {
			fmt.Println("> mysql update basic affect", affect)
		}
	}
}

// MysqlDeleteBasic :
func MysqlDeleteBasic() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// test max
	maxID, _ := db.T("story").Max("id")
	if maxID > 0 {
		// test delete
		affect, err := db.T("story").Delete("id=?", maxID)
		if err != nil {
			fmt.Println("> mysql delete basic err", err)
		} else {
			fmt.Println("> mysql delete basic affect", affect)
		}
	}
}

// MysqlTxBasic :
func MysqlTxBasic() {
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// tx begin
	db.Begin()
	// tx select
	id, err := db.T("story").Max("id")
	if err != nil {
		gutil.Dump(err)
	} else {
		gutil.Dump("before insert", id)
	}
	// tx insert
	id, err = db.T("story").Insert("title=?,content=?,dtime=?", "title N", "content N", time.Now().Unix())
	if err != nil {
		gutil.Dump("insert fail", id, err)
		db.Rollback()
	}
	// tx select
	id, err = db.T("story").Max("id")
	if err != nil {
		gutil.Dump(err)
	} else {
		gutil.Dump("before commit", id)
	}
	// tx commit
	gutil.Dump("insert ok", id)
	db.Commit()
	// tx select
	id, err = db.T("story").Max("id")
	if err != nil {
		gutil.Dump(err)
	} else {
		gutil.Dump("after commit", id)
	}
}

// MysqlQueryShard :
func MysqlQueryShard() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.C("user")
	defer db.Close()
	// init result struct
	type UserInfo struct {
		ID    int
		Name  string
		Dtime string
	}
	// test update (get latest seq id from gb_base.seq_user)
	maxID := MysqlMaxUserSeqId()
	// test FetchMaps
	rows, err := db.T("user_info").Select("id,name,dtime", "id=?", maxID)
	if err != nil {
		fmt.Println("> mysql query shard err", err)
	} else {
		r, _ := db.T("user_info").FetchMaps(rows)
		for k, v := range *r {
			fmt.Println("> mysql query shard : FetchMaps - id :", k)
			gutil.Dump(v)
		}
	}
	// test FetchMap
	rows, err = db.T("user_info").Select("id,name,dtime", "id=?", maxID)
	if err != nil {
		fmt.Println("> mysql query shard err", err)
	} else {
		r, _ := db.T("user_info").FetchMap(rows)
		fmt.Println("> mysql query shard : FetchMap")
		gutil.Dump(r)
	}
	// test FetchStructs
	rows, err = db.T("user_info").Select("id,name,dtime", "id=?", maxID)
	if err != nil {
		fmt.Println("> mysql query shard err", err)
	} else {
		s := &UserInfo{}
		r, _ := db.T("user_info").FetchStructs(rows, s)
		for k, v := range *r {
			fmt.Println("> mysql query shard : FetchStructs - id :", k)
			gutil.Dump(v.(UserInfo))
		}
	}
	// test FetchStruct
	rows, err = db.T("user_info").Select("id,name,dtime", "id=?", maxID)
	if err != nil {
		fmt.Println("> mysql query shard err", err)
	} else {
		s := &UserInfo{}
		r, _ := db.T("user_info").FetchStruct(rows, s)
		fmt.Println("> mysql query shard : FetchStruct")
		gutil.Dump(r.(UserInfo))
	}
}

// MysqlInsertShard :
func MysqlInsertShard() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.C("user")
	defer db.Close()
	// test insert (push seq id to gb_base.seq_user)
	id, err := db.T("user_info").Insert("id=?,name=?,dtime=?", MysqlUserSeqId(), "Name N", time.Now().Unix())
	if err != nil {
		fmt.Println("> mysql insert shard err", err)
	} else {
		fmt.Println("> mysql insert shard id", id)
	}
}

// MysqlUpdateShard :
func MysqlUpdateShard() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.C("user")
	defer db.Close()
	// test update (get latest seq id from gb_base.seq_user)
	maxID := MysqlMaxUserSeqId()
	if maxID > 0 {
		// test update
		name := "name " + strconv.FormatInt(maxID, 10)
		affect, err := db.T("user_info").Update("name=? where id=?", name, maxID)
		if err != nil {
			fmt.Println("> mysql update shard err", err)
		} else {
			fmt.Println("> mysql update shard affect", affect)
		}
	}
}

// MysqlDeleteShard :
func MysqlDeleteShard() {
	// print debug info
	gdopool.SetDebug(true)
	// init by driver
	db := gdo.C("user")
	defer db.Close()
	// test delete (get latest seq id from gb_base.seq_user)
	maxID := MysqlMaxUserSeqId()
	if maxID > 0 {
		// test delete
		affect, err := db.T("user_info").Delete("id=?", maxID)
		if err != nil {
			fmt.Println("> mysql delete shard err", err)
		} else {
			fmt.Println("> mysql delete shard affect", affect)
		}
	}
}

func MysqlUserSeqId() int64 {
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// test insert
	id, err := db.T("seq_user").Insert("dtime=?", time.Now().Unix())
	if err != nil {
		fmt.Println("> mysql user seq err", err)
	} else {
		fmt.Println("> mysql user seq id", id)
	}
	// return
	return int64(id)
}

func MysqlMaxUserSeqId() int64 {
	// init by driver
	db := gdo.D("base")
	defer db.Close()
	// test insert
	maxID, _ := db.T("seq_user").Max("id")
	// return
	return int64(maxID)
}
