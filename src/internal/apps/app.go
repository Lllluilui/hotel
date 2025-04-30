package apps

import (
	"log"

	"hotel/internal/apps/apis"
)

// App 用用
type App struct {
}

// Start
//
//	@desc: 启动
//	@receiver app *App
//	@return error
func (app *App) Start() error {
	if err := new(apis.Register).RegisterAndRun(); err != nil {
		log.Fatal("run server error:" + err.Error())
	}

	return nil
}
