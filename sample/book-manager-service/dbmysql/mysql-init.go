package dbmysql

import (
	"examples/sample/book-manager-service/db"
)

func Init() db.MasterDB {
	return db.MasterDB{
		BookDBImpl: initBookDBImpl(),
	}
}
