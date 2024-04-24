package main

func main() {
	LoadENV()
	initSQL()
	dao := NewUserDaoAdapter(DB, NewUserCoreAdapter())
	userHandler = NewUserHandlerAdapter(dao)
	InitializeRouter()
}
