package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}

func Init(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true, // 关闭默认事务
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// CursorPage 游标翻页
func CursorPage(cursor, num int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id > ?", cursor).Limit(int(num))
	}
}
func CursorPageWithIDName(cursor, num int64, idName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(idName+" > ?", cursor).Limit(int(num))
	}
}

// BigCursorPage 大端模式游标翻页
func BigCursorPage(cursor, num int64) func(db *gorm.DB) *gorm.DB {
	return BigCursorPageWithIDName(cursor, num, "id")
}
func BigCursorPageWithIDName(cursor, num int64, idName string) func(db *gorm.DB) *gorm.DB {
	if cursor > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(idName+" < ?", cursor).Limit(int(num))
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db.Limit(int(num)) }
}
