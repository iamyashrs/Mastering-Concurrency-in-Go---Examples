package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

const TESTLENGTH = 100000

type CPUHog struct {
	hog []byte
}

func makeLongByte() []byte {
	longByte := make([]byte, TESTLENGTH)
	for i := 0; i < TESTLENGTH; i++ {
		longByte[i] = byte(i)
	}
	return longByte
}

var profile = flag.String("cpuprofile", "", "output pprof data to file")

func main() {
	var cpuhog []CPUHog

	flag.Parse()
	if *profile != "" {
		flag, err := os.Create(*profile)
		if err != nil {
			fmt.Println("Could not create profile", err)
		}
		pprof.StartCPUProfile(flag)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < TESTLENGTH; i++ {
		cpuh := CPUHog{}
		cpuh.hog = makeLongByte()
		cpuhog = append(cpuhog, cpuh)
	}
}
