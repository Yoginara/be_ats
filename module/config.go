package module

import "os"

package be_ats

import (
"github.com/aiteung/atdb"
"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "ats",
}

var MongoConn = atdb.MongoConnect(MongoInfo)
