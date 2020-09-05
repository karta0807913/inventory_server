package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/karta0807913/inventory_server/model"
)

func main() {
	force := flag.Bool("f", false, "force clear database")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("usage %s: <csv file> <out sqlite>\n", os.Args[0])
		os.Exit(1)
	}
	csvFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println("open csv file error", err)
		os.Exit(1)
	}
	db, err := model.SqliteDB(args[1])
	if err != nil {
		fmt.Println("create db error", err)
		os.Exit(1)
	}
	err = model.InitDB(db)
	if err != nil {
		fmt.Println("create db error", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(csvFile)
	dropCR := func(data []byte) []byte {
		if len(data) > 0 && data[len(data)-1] == '\r' {
			return data[0 : len(data)-1]
		}
		return data
	}
	ScanCRLF := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
			// We have a full newline-terminated line.
			return i + 2, dropCR(data[0:i]), nil
		}
		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return len(data), dropCR(data), nil
		}
		// Request more data.
		return 0, nil, nil
	}
	scanner.Split(ScanCRLF)
	var first bool = true
	for scanner.Scan() {
		raw_text := strings.ReplaceAll(scanner.Text(), " ", "")
		if first {
			first = false
			continue
		}
		values := strings.Split(raw_text, ",")
		if len(values) < 3 {
			fmt.Printf("can't convert %s\n", raw_text)
			continue
		}
		item := model.ItemTable{
			ItemID: values[0] + "-" + values[1],
			Name:   values[2],
		}
		if len(values) > 8 {
			item.Date = values[8]
		}
		if len(values) > 12 {
			item.Location = values[12]
		}

		if len(values) > 14 {
			item.Note = values[14]
		}
		ageLimit := 0
		if len(values) > 9 {
			ageLimit, err = strconv.Atoi(values[9])
		}
		if err != nil {
			fmt.Printf("%s age limit convert failed, set to 0 \n", item.Name)
		}
		item.AgeLimit = uint(ageLimit)

		err = db.Where(model.ItemTable{ItemID: item.ItemID}).FirstOrCreate(&item).Error
		if err != nil {
			fmt.Println("insert", item.Name, "error", err)
		} else {
			if *force {
				db.Where(model.ItemTable{ItemID: item.ItemID}).Updates(&item)
			}
		}
	}
}
