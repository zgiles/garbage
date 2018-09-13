package main

import (
	"flag"
	"github.com/dustin/go-humanize"
	"github.com/zgiles/garbage"
	"io"
	"log"
	"os"
	"sync"
)

func garbagegen(out io.Writer, size int64, wg *sync.WaitGroup) {
	r := garbage.NewLimitedReader(size)
	buf := make([]byte, 65536)
	_, _ = io.CopyBuffer(out, r, buf)
	wg.Done()
}

func zerogen(out io.Writer, size int64, wg *sync.WaitGroup) {
	r := garbage.NewLimitedZeroReader(size)
	buf := make([]byte, 65536)
	_, _ = io.CopyBuffer(out, r, buf)
	wg.Done()
}

func main() {
	output := flag.String("output", "", "output file or '-' for stdout (default '-')")
	sure := flag.Bool("stdout", false, "confirm output to stdout")
	threads := flag.Int("threads", 2, "number of threads for generating garbage")
	size := flag.String("size", "8GiB", "amount of data to generate")
	source := flag.String("source", "aes", "[ aes, zero ]")
	flag.Parse()

	if *source != "aes" && *source != "zero" {
		log.Fatal("Not a valid source")
	}

	sizenum, sizeerr := humanize.ParseBytes(*size)
	if sizeerr != nil {
		log.Fatal(sizeerr)
	}

	sizeperthread := int64(sizenum) / int64(*threads)

	var o io.Writer
	if *output == "" || *output == "-" {
		so, _ := os.Stdout.Stat()
		if ((so.Mode() & os.ModeCharDevice) == 0) || *sure {
			o = os.Stdout
		} else {
			log.Fatal("You must specify -stdout to output to stdout to prevent from console flooding")
		}
	} else {
		f, ferr := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0644)
		if ferr != nil {
			log.Fatal(ferr)
		}
		o = f
	}

	wg := &sync.WaitGroup{}
	for x := 0; x < *threads; x++ {
		wg.Add(1)
		if *source == "aes" {
			go garbagegen(o, int64(sizeperthread), wg)
		} else if *source == "zero" {
			go zerogen(o, int64(sizeperthread), wg)
		} else {
			log.Fatal("Not a valid source")
		}
		//go garbagegen(w, wg)
	}
	wg.Wait()
}
