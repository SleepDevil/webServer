package core

import (
	"yasi_audio/api"
	"yasi_audio/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.POST("/register", api.Register)
	Router.POST("/login", api.Login)
	Router.GET("/invite", api.Invite)
	Router.POST("/invite", api.Check_Invited)
	Router.Run()
	// print(s)
	// time.Sleep(10 * time.Microsecond)
	// global.GVA_LOG.Info("server run success on ", zap.String("address", "127.0.0.1:8000"))
	// global.GVA_LOG.Error(s.ListenAndServe().Error())
}
