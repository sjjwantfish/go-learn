package main

import (
	"database/sql"
	//"errors"
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	// 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
	//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
	DoSthWithDB()
}

type Students struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func DoSthWithDB() {
	execute_sql := "select * from students;"
	students, err := GetOtherError(execute_sql)
	//students, err := GetNoRowsError(execute_sql)
	if errors.Is(err, sql.ErrNoRows) {
		// do nothing because students is nil
		return
	} else if err != nil {
		// log
		fmt.Printf("errorType:[%T] Message:[%v]\nStack Trace: %+v", errors.Unwrap(err), err, err)
		return
	}
	// do sth else with students
	fmt.Println("students: ",students)
	return
}

func GetNoRowsError(s string) ([]Students, error) {
	return nil, errors.Wrap(sql.ErrNoRows, "inputSql:" + s)
}

func GetOtherError(s string) ([]Students, error)  {
	return nil, errors.Wrap(sql.ErrConnDone, "inputSql:" + s)
}

