package services

import (
	"fmt"
	"time"
)

func LastUpdated() string {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t := time.Now().In(loc)
	return t.Format(time.RFC822)
}
