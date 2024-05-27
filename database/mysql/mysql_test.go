package mysql

import (
	"fmt"
	"testing"
)

func TestQueryRow(t *testing.T) {
	err := initDB()
	if err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	queryRowDemo()
}

func TestQueryMulti(t *testing.T) {
	err := initDB()
	if err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	queryMultiDemo()
}
