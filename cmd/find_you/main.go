package main

import (
	"flag"
	"fmt"
	"gihub.com/Amore14rn/will_find_you/internal"
	"github.com/fatih/color"
	"github.com/gammazero/workerpool"
	"os"
	"strings"
)

var (
	queue int
)

func Action(url string) {
	sl := internal.Visitor(url, 10)
	internal.CheckTakeOver(internals.RemoveDuplicateStrings(sl))
	color.Magenta("Finished Checking: " + url)
	queue--
	fmt.Println("Remaining URLs:", queue)

}

func Run() {
	internal.LOGO()
	urlFile := flag.String("f", "", "Path of the URL file")
	numWorker := flag.Int("w", 5, "Number of worker.")
	flag.Parse()
	if *urlFile == "" {
		fmt.Println("Please specify all arguments!")
		flag.PrintDefaults()
		os.Exit(1)
	}
	file, err := os.ReadFile(*urlFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	urls := strings.Split(string(file), "\n")
	queue = len(urls)
	fmt.Println("Total URLs:", queue)
	wp := workerpool.New(*numWorker)

	for _, url := range urls {
		url := url
		wp.Submit(func() {
			fmt.Println("Checking:", url)
			Action(url)
		})

	}
	wp.StopWait()

	color.Cyan("Scan Completed")

}

func main() {
	Run()
}
