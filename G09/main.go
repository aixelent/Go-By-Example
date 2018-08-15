package main

import (
	"time"
	"fmt"
)

func main()  {
	t := time.Now()
	fmt.Println("Location", t.Location(), " Time : ", t)	// local time

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("location: ", location, "\nTime: ", t.In(location))

	locShan, _ := time.LoadLocation("Asia/Shanghai")
	nowShan := time.Now().In(locShan)
	fmt.Println("Location: ", locShan, "\nTime: ", nowShan)

	locSeoul, _ := time.LoadLocation("Asia/Seoul")
	nowSeoul := time.Now().In(locSeoul)
	fmt.Println("Location: ", locSeoul, "\nTime: ", nowSeoul)

	locTokyo, _ := time.LoadLocation("Asia/Tokyo")
	nowTokyo := time.Now().In(locTokyo)
	fmt.Println("Location: ", nowSeoul, "\nTime: ", nowTokyo)

	locPort, _ := time.LoadLocation("America/Port-au-Prince")
	nowPort := time.Now().In(locPort)
	fmt.Println("Location: ", locPort, "\nTime: ", nowPort)
}

// experiment, do not pay attention :-)
//func getTokyo() ()  {
//	getLocTokyo, _ := time.LoadLocation("Asia/Tokyo")
//	getNowTokyo := time.Now().In(getLocTokyo)
//	resLoc := getLocTokyo
//	resTime := getNowTokyo
//
//	return resLoc, resTime
//}