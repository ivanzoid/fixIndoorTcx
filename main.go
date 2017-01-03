package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/philhofer/tcx"
)

var (
	stdErr = log.New(os.Stderr, "", 0)
)

func main() {

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("Usage: %s <tcxFile> <distance>\n", os.Args[0])
		return
	}

	filePath := args[0]
	dist, _ := strconv.ParseFloat(args[1], 64)

	db, err := tcx.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Can't read file: %v\n", err)
	}

	for _, activity := range db.Acts.Act {

		for _, lap := range activity.Laps {

			ptCount := len(lap.Trk.Pt)
			if ptCount <= 2 {
				continue
			}

			distInc := float64(dist) / float64(ptCount-1)
			currentDist := 0.0

			for i := range lap.Trk.Pt {
				pt := &lap.Trk.Pt[i]
				pt.Dist = currentDist
				currentDist += distInc
			}
		}
	}

	bytes, err := tcx.ToBytes(*db)

	if err != nil {
		log.Fatalf("Can't convert result data to bytes: %v\n", err)
	}

	bytesString := string(bytes[:])

	fmt.Println("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	fmt.Println(bytesString)
}
