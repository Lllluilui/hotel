package main

import (
	"hotel/internal/apps"
	"hotel/internal/pkg/di"
)

// main
//
//	@desc: 主函数
func main() {
	di.InitMysql()

	_ = new(apps.App).Start()
}
