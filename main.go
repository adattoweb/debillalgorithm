package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generate(){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	speed := 2000 + rand.Intn(500);
	fmt.Printf("[%s] GPU 0: %d H/s\n", currentTime, speed)

	letters := []string{"a", "b", "c", "e", "d", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	hash := "0x";
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(letters))
		hash += letters[index]
	}
	hash += "..."
	fmt.Printf("[%s] Submitting share: %s\n", currentTime, hash)

}

func solution(){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	speed := 2000 + rand.Intn(500);
	fmt.Printf("[%s] GPU 0: %d H/s\n", currentTime, speed)

	letters := []string{"a", "b", "c", "e", "d", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	hash := "0x";
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(letters))
		hash += letters[index]
	}
	hash += "..."
	fmt.Printf("[%s] Found solution for share: %s\n", currentTime, hash)
	fmt.Printf("[%s] Submiting solution: %s\n", currentTime, hash)
}

func block(){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	speed := 2000 + rand.Intn(500);
	fmt.Printf("[%s] GPU 0: %d H/s\n", currentTime, speed)

	letters := []string{"a", "b", "c", "e", "d", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	hash := "0x";
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(letters))
		hash += letters[index]
	}
	hash += "..."
	fmt.Printf("[%s] Found block: %s\n", currentTime, hash)
	fmt.Printf("[%s] Submitted block to pool, waiting for confirmation\n", currentTime)
}

func main(){
	fmt.Println("De Bill algorithm activated!")
	time.Sleep(1 * time.Second)
	var bitcoin float64;
	bitcoin += float64(rand.Intn(100)) / float64((rand.Intn(10000) + 1000))
	count := 0;
	for count < 5000{
		count++
		generate()
		solution()
		block()
	}
	fmt.Printf("Earned: %.6f BTC \n", bitcoin)
}