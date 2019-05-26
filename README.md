# autosql
gorm自动生成表相对应的model

#使用方法
	data := database.Database{User:"root", Password:"root", Addr:"127.0.0.1", Port:"3306", Db:"ay_wisdom", Prefix:"ay_"}
	database.Init(data)

#做的一个小工具，大家如果觉得用的到可以试试
