package common

import (
	"fmt"
	"os"
)

func ErrorHandle(err error,s string){
	if err!=nil{
		fmt.Println(s," err: ",err)
		os.Exit(2)
	}
}
