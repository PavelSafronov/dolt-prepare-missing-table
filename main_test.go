package test_dolt_prepare

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func TestMain(t *testing.T) {
	// setup
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/obsidian")
	if err != nil {
		t.Fatalf("can't connect to mysql: %s", err)
	}
	defer db.Close()
	_, err = db.Exec("use obsidian")
	if err != nil {
		t.Fatalf("unable to use obsidian: %s", err)
	}
	_, err = db.Exec("drop table if exists _pavel_testing")
	if err != nil {
		t.Fatalf("unable to drop _pavel_testing: %s", err)
	}

	// test
	stmt, err := db.Prepare("SELECT `id` FROM `_pavel_testing` ORDER BY `id` ASC")

	// verify
	if err != nil {
		if mse, ok := err.(*mysql.MySQLError); ok {
			if mse.Number == 1146 {
				t.Log("SUCCESS!!")
			} else if mse.Number == 1105 {
				t.Fatalf("still getting 1105: %s", err)
			} else {
				t.Fatalf("unexpected number %v for error %s", mse.Number, err)
			}
		} else {
			t.Fatalf("unexpected error: %s", err)
		}
	} else {
		t.Fatalf("expected error, got none, stmt: %v", stmt)
	}
}
