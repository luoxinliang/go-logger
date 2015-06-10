package logger

import (
//	"runtime"
	"strconv"
	"fmt"
	"testing"
	"sync"

	"github.com/luoxinliang/go-logger"
)

var wg sync.WaitGroup

func log(i int) {
	fmt.Println("log i=", i)
	logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	logger.Debugf("Debug>>>>>>>>>>>>>>>>>>>>>> %d",i)

	logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	logger.Debugf("Debug>>>>>>>>>>>>>>>>>>>>>> %d",i)

	logger.Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	logger.Warnf("Debug>>>>>>>>>>>>>>>>>>>>>> %d",i)

	logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	logger.Errorf("Debug>>>>>>>>>>>>>>>>>>>>>> %d",i)

	//	logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	//	logger.Fatalf("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>%d", i)

	wg.Done()
}

func Test(t *testing.T) {
	logger.Init(".", "test.log", logger.DEBUG)

	//指定是否控制台打印，默认为false
//	logger.SetConsole(true)

	for i := 10; i > 0; i-- {
		wg.Add(1)
		log(i)
	}

	wg.Wait()
}