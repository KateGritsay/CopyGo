package main

import (
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

func main() {
	from := flag.String("from", "", "path to source")
	to := flag.String("to", "", "path to dest")
	offset :=flag.Int("offset", 0, "offset point for start copy")
	limit := flag.Int("limit", 0, "limit bytes to copy")
	flag.Parse()
	fromFile, err := os.Open(*from)
	if err != nil {
		fmt.Println("Can't open file for coping", err)
		os.Exit(1)
	}
	defer fromFile.Close()

	toFile, err := os.Create(*to)
	if err != nil {
		fmt.Println("Can't create copy of file", err)
		os.Exit(1)
	}
	defer toFile.Close()

	if *limit == 0 {
		stat, err := fromFile.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		*limit = int(stat.Size())

	}

	if *offset == 0 {
		written, err := io.CopyN(toFile, fromFile, int64(*limit))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(written)
	} else {

		buffer := make([]byte, *offset)
		for {
			written, err := fromFile.ReadAt(buffer, int64(*offset))

			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				fmt.Println(written)
			}
		}

		// create and start new bar
		bar := pb.StartNew(*limit)
		defer bar.Finish()

	}
}