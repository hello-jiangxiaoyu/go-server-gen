package orm

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateGormModel() {
	g, err := getGenerator()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		var result []byte
		for i := 0; i < len(columnName); i++ {
			if columnName[i] != '_' {
				result = append(result, columnName[i])
				continue
			}
			i++
			if i >= len(columnName) {
				continue
			}
			if 'a' <= columnName[i] && columnName[i] <= 'z' {
				result = append(result, columnName[i]+'A'-'a') // 单词首字母大写
			} else {
				result = append(result, columnName[i])
			}
		}
		return string(result)
	})
	g.GenerateAllTable(jsonField)
	g.Execute()
	println("==== success ====")
}

func getGenerator() (*gen.Generator, error) {
	db, err := gorm.Open(mysql.Open("gorm:gorm@tcp(127.0.0.1:3306)/im?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "query", // 相对执行`go run`时的路径, 会自动创建目录
		ModelPkgPath:      "biz/model/mysql",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)
	return g, err
}
