package main

import (
	"fmt"
	"sync"
	"time"
)

var icons map[string]string
var mu sync.Mutex

func loadIcons() {

	//icons = map[string]string{
	//	"spades.png":   loadIcon("spades.png"),
	//	"hearts.png":   loadIcon("hearts.png"),
	//	"diamonds.png": loadIcon("diamonds.png"),
	//	"clubs.png":    loadIcon("clubs.png"),
	//}
	icons = make(map[string]string)
	time.Sleep(1 * time.Second)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}
func Icon(name string) string {
	mu.Lock()
	loadIcons()
	defer mu.Unlock()
	return icons[name]
}
func loadIcon(name string) string {
	return "image_" + name
}
func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	beginTestTime := time.Now()
	for i := 0; i < 10; i++ {
		go func(index int) {
			defer wg.Done()
			fmt.Printf("gorountines %d:%s", index, Icon("clubs.png")+"\n")
		}(i)
	}
	wg.Wait()
	duration := time.Since(beginTestTime)
	fmt.Println("complete:" + duration.String())

}
