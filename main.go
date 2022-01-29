package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool // 行番号出力オプション
var filepaths string // ファイルパスオプション

func init() {
	// フラッグ名付き行番号出力オプション入力
	flag.BoolVar(&n, "n", false, "whether line number is displayed")
	// フラッグ名なしファイルパス入力
	flag.StringVar(&filepaths, "", "", "whether line number is displayed")
}

func main() {
	flag.Parse()
	ln := 0
	for _, fn :=  range flag.Args() {
		readFile(fn, &ln)
	}
}

func readFile(fn string, ln *int) error {
	// ファイルの処理
	f, err := os.Open(fn)
	if err != nil { return err }
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if n {
			fmt.Fprintf(os.Stdout, "%d: %s\n", *ln, scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
		*ln++
	}
	return err
}