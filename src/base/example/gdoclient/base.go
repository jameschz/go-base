package gdoclient

import (
	"base/gdo"
	"base/util"
	"fmt"
	"strconv"
	"time"
)

func MysqlQueryBasic() {
	// init by driver
	db := gdo.D("demo")
	defer db.Close()
	// test select
	rows, err := db.T("story").Select("*", "1=1")
	if err != nil {
		fmt.Println("> mysql query basic err", err)
	} else {
		for rows.Next() {
			var id string
			var title string
			var content string
			var time string
			rows.Scan(&id, &title, &content, &time)
			fmt.Println("> mysql query basic [id]", id, "[title]", title, "[content]", content, "[time]", time)
		}
	}
}

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

func MysqlUpdateBasic() {
	// init by driver
	db := gdo.D("demo")
	defer db.Close()
	// test max
	maxId, _ := db.T("story").Max("id")
	if maxId > 0 {
		// test update
		title := "title " + strconv.FormatInt(maxId, 10)
		content := "content " + strconv.FormatInt(maxId, 10)
		affect, err := db.T("story").Update("title=?,content=? where id=?", title, content, maxId)
		if err != nil {
			fmt.Println("> mysql update basic err", err)
		} else {
			fmt.Println("> mysql update basic affect", affect)
		}
	}
}

func MysqlDeleteBasic() {
	// init by driver
	db := gdo.D("demo")
	defer db.Close()
	// test max
	maxId, _ := db.T("story").Max("id")
	if maxId > 0 {
		// test delete
		affect, err := db.T("story").Delete("id=?", maxId)
		if err != nil {
			fmt.Println("> mysql delete basic err", err)
		} else {
			fmt.Println("> mysql delete basic affect", affect)
		}
	}
}

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
