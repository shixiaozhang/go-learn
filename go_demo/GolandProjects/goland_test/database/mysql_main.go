package main
import "goland_test/dbs"

func main() {

	dbs.StructInsert()
	dbs.StructUpdate()
	dbs.StructQueryField()
	dbs.StructQueryAllField()
	dbs.StructDel()
	dbs.StructTx()
	dbs.RawQueryField()
	dbs.RawQueryAllField()
}