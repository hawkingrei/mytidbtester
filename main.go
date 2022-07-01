package main

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var wg sync.WaitGroup
	for c := 1; c <= 10; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sbtest")
			defer db.Close()

			if err != nil {
				log.Fatal(err)
			}
			start := time.Now()
			sqlstr := "CREATE INDEX index_c5 ON sbtest" + strconv.Itoa(c) + "(c5)"
			defer func() {
				elapsed := time.Since(start)
				log.Printf("(%d) %s: %s", c, sqlstr, elapsed)
			}()

			// db create index
			_, err = db.Exec(sqlstr)

			if err != nil {
				log.Fatal(err)
			}
		}(c)
	}
	wg.Wait()
}
