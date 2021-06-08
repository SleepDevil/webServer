package core

import (
	"yasi_audio/api"
	"yasi_audio/initialize"
)

// type server interface {
// 	ListenAndServe() error
// }

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.MaxMultipartMemory = 8 << 20
	Router.POST("/register", api.Register)
	Router.POST("/login", api.Login)
	Router.GET("/invite", api.Invite)
	Router.POST("/checkinvite", api.Check_Invited)
	Router.POST("/task", api.CreateTask)
	Router.GET("/task", api.GetTasks)
	Router.POST("/checktoken", api.CheckJwt)
	Router.POST("/room", api.CreateRoom)
	Router.POST("/avator", api.CreateAvator)
	Router.GET("/avator/update/:username/:nickname", api.UpdaeUserHeaderImg)
	Router.POST("/task/accept", api.AcceptTask)
	Router.POST("/token", api.GetToken)
	Router.Run()
	// print(s)
	// time.Sleep(10 * time.Microsecond)
	// global.GVA_LOG.Info("server run success on ", zap.String("address", "127.0.0.1:8000"))
	// global.GVA_LOG.Error(s.ListenAndServe().Error())
}
