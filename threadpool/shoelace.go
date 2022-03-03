package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2d struct {
	x, y int
}

/*
	Shoelace algofrithm is generally used to calculate the area of the polygon of arbitary size
	described by the points
*/

const numberOfThreads int = 8

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitgroup = sync.WaitGroup{}
)

func findArea(inputChannel chan string) {
	for pointStr := range inputChannel {
		var points []Point2d
		for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2d{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitgroup.Done()
}
func main() {
	absPath, _ := filepath.Abs("C:/Users/Acer/go/src/github.com/Concurrency-Go/threadpool/")
	dat, _ := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(dat)

	/*
		A particular line include points in format (1,100) (2,20) ...
	*/

	inputChannel := make(chan string, 1000)
	for i := 0; i < numberOfThreads; i++ {
		go findArea(inputChannel)
	}
	waitgroup.Add(numberOfThreads)
	// 8 threads are there so we need all the threads to come back before exiting the main
	start := time.Now()

	// Input channel receives the point line by line and @ max can store about 1000 points
	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)
	waitgroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s \n", elapsed)
}
