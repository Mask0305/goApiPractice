package main

import (
	userRouter "goApiPractice/routers"
)

func main() {
	userRouter := userRouter.UserRouter()
	userRouter.Run(":3000")
}
