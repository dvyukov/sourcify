package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`^([a-zA-Z0-9/_\-\.]+):([0-9]+): `)
	files := make(map[string][]string)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Printf("%v\n", s.Text())
		match := re.FindStringSubmatch(s.Text())
		if match == nil {
			continue
		}
		ln, err := strconv.ParseUint(match[2], 10, 32)
		if err != nil {
			continue
		}
		f := loadFile(files, match[1])
		if f == nil || ln >= uint64(len(f)) {
			continue
		}
		fmt.Printf("%v\n", f[ln])
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read stdin: %v\n", err)
		os.Exit(1)
	}
}

func loadFile(files map[string][]string, file string) []string {
	if f, ok := files[file]; ok {
		return f
	}
	fd, err := os.Open(file)
	if err != nil {
		files[file] = nil
		return nil
	}
	defer fd.Close()
	var data []string
	data = append(data, "") // lines are 1-based
	s := bufio.NewScanner(fd)
	for s.Scan() {
		data = append(data, s.Text())
	}
	files[file] = data
	return data
}
