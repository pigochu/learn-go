// Windows 下捕獲 CTRL+C
// 這個範例是參考以網址修改的
// 如果 10 秒內按下 CTRL + C , 則會印出 Signal 的名稱
// 如果超過 10 秒沒有按下，則印出 Timeout
// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe

package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			println("Signal : " + sig.String())
			println("bye")
			os.Exit(0)
		}
	}()

	time.Sleep(time.Duration(10) * time.Second)
	println("Timeout , bye")
}
