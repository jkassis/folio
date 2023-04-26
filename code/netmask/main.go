// Goal: parse the nginx log file and answer the following questions:
//   1. How many lines contain IPs within each of the first four subnets?
//   2. How many lines contain IPs within the final subnet?
//   3. How many lines contain IPs within two of the subnets?

// really just get a count for log records in each subnet
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Feel free to make the code specific to the following subnets (i.e. you don't need to parse these).
var subnets = []string{
	"192.0.0.0/8",
	"93.180.0.0/16",
	"80.91.0.0/16",
	"192.235.0.0/16",
	"217.168.17.2/29",
}

var subnetsParsed = make([]*Net, len(subnets))

type Net struct {
	ip    int64
	mask  int
	count int64
}

func (a *Net) Parse(str string) {
	parts := strings.Split(str, "/")
	ipStr := parts[0]

	if len(parts) > 1 {
		maskStr := parts[1]
		a.mask, _ = strconv.Atoi(maskStr)
	}

	quadParts := strings.Split(ipStr, ".")
	p0, _ := strconv.Atoi(quadParts[0])
	a.ip += int64(p0) << 24
	p1, _ := strconv.Atoi(quadParts[1])
	a.ip += int64(p1) << 16
	p2, _ := strconv.Atoi(quadParts[2])
	a.ip += int64(p2) << 8
	p3, _ := strconv.Atoi(quadParts[3])
	a.ip += int64(p3)
}

func (a1 *Net) Contains(a2 *Net) bool {
	c1 := a1.ip & (0xFFFFFFFF << (32 - a1.mask))
	c2 := a2.ip & (0xFFFFFFFF << (32 - a1.mask))
	// fmt.Printf("comparing %d to %d\n", c1, c2)
	return c1 == c2
}

func parseLogLine(line string) {
	defer func() {
		err := recover()
		if err != nil {
			// probably some parse err
			//  fmt.Println(fmt.Sprintf("line: %s", line))
		}
	}()
	parts := strings.Split(line, " ")
	ipStr := parts[0]
	a2 := &Net{}
	a2.Parse(ipStr)

	for _, a1 := range subnetsParsed {
		if a1.Contains(a2) {
			a1.count++
			// fmt.Println(fmt.Sprintf("line: %s", line))
		}
	}
}

func main() {
	for i, subnetStr := range subnets {
		a := &Net{}
		a.Parse(subnetStr)
		subnetsParsed[i] = a
	}

	logFile, err := os.Open("./golang/netmask/nginx_logs.txt")
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logScanner := bufio.NewScanner(logFile)
	for logScanner.Scan() {
		parseLogLine(logScanner.Text())
	}

	for _, subnet := range subnetsParsed {
		fmt.Printf("%d : %d\n", subnet.ip, subnet.count)
	}
}
