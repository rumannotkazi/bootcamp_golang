package concur

import (
	// "sync"
	"time"
)

func TheWork() string {

	time.Sleep(5 * time.Second)
	return "Done!"
}
