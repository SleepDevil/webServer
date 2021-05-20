package main

import (
	"yasi_audio/core"
	"yasi_audio/global"
	"yasi_audio/initialize"
)

func main() {
	global.GVA_DB = initialize.Gorm()
	// initialize.MysqlTables(global.GVA_DB)
	core.RunWindowsServer()
	initialize.Redis()
}
