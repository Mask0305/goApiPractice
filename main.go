package main 

import (
	userRouter"goApiPractice/routers"
)

func main(){
	userRouter := userRouter.userRouter()
	userRouter.Run(":3000")
}