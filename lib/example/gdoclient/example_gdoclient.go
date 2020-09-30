package examplegdoclient

import (
	"fmt"
	"github.com/jameschz/go-base/lib/gdo"
	"github.com/jameschz/go-base/lib/util"
	"strconv"
	"time"
)

// MysqlQueryBasic :
func MysqlQueryBasic() {
	// init by driver
	db := gdo.D("demo")
	defer db.Close()
	// init result struct
	type story struct {
		ID      int
		Title   string
		Content string
		Dtime   string
	}
	// test select
	rows, err := db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		stu1 := &story{}
		res1, _ := db.T("story").FetchRow(rows, stu1)
		fmt.Println("> mysql query basic : FetchRow")
		util.Dump(res1.(story))
		rows.Close()
	}
	// test select
	rows, err = db.T("story").Select("id,title,content,dtime", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		stu2 := &story{}
		res2, _ := db.T("story").FetchAll(rows, stu2)
		for k, v := range *res2 {
			fmt.Println("> mysql query basic : FetchAll", k)
			util.Dump(v.(story))
		}
		rows.Close()
	}

}

// MysqlInsertBasic :
func MysqlInsertBasic() {
	// init by driver
	db := gdo.D("demo")
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
	// init by driver
	db := gdo.D("demo")
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
	// init by driver
	db := gdo.D("demo")
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
	db := gdo.D("demo")
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
