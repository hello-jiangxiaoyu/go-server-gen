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
func CursorPage(cursor, limit int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id > ?", cursor).Limit(int(limit))
	}
}
func CursorPageWithIDName(cursor, limit int64, idName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(idName+" > ?", cursor).Limit(int(limit))
	}
}

// BigCursorPage 大端模式游标翻页
func BigCursorPage(cursor, limit int64) func(db *gorm.DB) *gorm.DB {
	return BigCursorPageWithIDName(cursor, limit, "id")
}
func BigCursorPageWithIDName(cursor, limit int64, idName string) func(db *gorm.DB) *gorm.DB {
	if cursor > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(idName+" < ?", cursor).Limit(int(limit))
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db.Limit(int(limit)) }
}
