package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
