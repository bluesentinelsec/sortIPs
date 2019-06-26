package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
)

func main() {

	// setup command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: sortIPs -i <input file>")
		os.Exit(1)
	}
	inputFile := flag.String("i", "", "Specifies the input file: sortIPs -i <input file>")
	flag.Parse()

	if len(*inputFile) == 0 {
		fmt.Println("Usage: sortIPs -i <input file>")
		os.Exit(1)
	}

	// open input file
	realIPs, err := openInputFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// run a function to compare IPs and sort
	sort.Slice(realIPs, func(i, j int) bool {
		return bytes.Compare(realIPs[i], realIPs[j]) < 0
	})

	// print sorted IPs to standard out
	for _, ip := range realIPs {
		fmt.Printf("%s\n", ip)
	}
}

// openInputFile opens the source file and appends each ip address to a list
func openInputFile(path string) ([]net.IP, error) {
	// open source file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// insert each IP address into a slice, "realIPs"
	realIPs := make([]net.IP, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text())
		realIPs = append(realIPs, ip)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return realIPs, err

}
