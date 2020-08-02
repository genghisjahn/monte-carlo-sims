package main

import (
	"fmt"
)

func createFighters(n int)(string,error){
	gs:=[]*gunfighter{}
	for f:=0;f<n;f++{
		gf:=gunfighter{}
		fmt.Println(f)
		//Get Name
		//Get 3 values for Confidence, Speed & Accuracy

		gs=append(gs,&gf)
	}
	return "",nil
}