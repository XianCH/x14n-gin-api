package initliza

func InitServer() {
	InitViperConfig()
	InitLoger2("./config", "debug")
	InitGorm()
}
