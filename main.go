package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"syscall"

	"github.com/mackerelio/go-osstat/memory"
)

func main() {
	displayMemory()
	displayCoreNumbers()
	displayDiskSize()
}

func displayMemory() {
	fmt.Print("\n============ MEMORY ============\n")
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	totalMemory := bToGb(memory.Total)
	fmt.Printf("Total memory: %.2f G\n", totalMemory)
	fmt.Printf("Reserved memory: %.2f G\n", getReservedMemory(totalMemory))
	fmt.Print("================================\n")
}

func bToGb(b uint64) float64 {
	return math.Round(float64(b) / 1024 / 1024 / 1024)
}

func getReservedMemory(total float64) float64 {
	return 0.25 + 0.75 + (total / 100 * 5) + 0.5
}

func displayCoreNumbers() {
	fmt.Print("\n============= CPU ==============\n")
	totalCPU := runtime.NumCPU()
	fmt.Printf("Total CPU (cores): %d\n", totalCPU)
	fmt.Printf("Reserved CPU (mi): %d\n", getReservedCPU(totalCPU))
	fmt.Print("================================\n")
}

func getReservedCPU(numCores int) int {
	return 50 + 50 + (numCores * 5)
}

func displayDiskSize() {
	fmt.Print("\n============= DISK =============\n")
	fs := syscall.Statfs_t{}
	syserr := syscall.Statfs("/", &fs)
	if syserr != nil {
		return
	}
	diskSize := bToGb(fs.Blocks * uint64(fs.Bsize))
	fmt.Printf("Root disk size %.2f G\n", diskSize)
	fmt.Printf("Reserved disk size %.2f G\n", getReservedDisk(diskSize))
	fmt.Print("================================\n")
}

func getReservedDisk(diskSize float64) float64 {
	eviction := diskSize / 10
	system := math.Log10(diskSize) * 10

	reserved := eviction + system
	if reserved > 100 {
		return 100
	}

	return reserved
}
