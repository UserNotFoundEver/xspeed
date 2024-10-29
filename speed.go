package main

import (
	"fmt"
	"sync"
	"time"
)

// Constants for the specific time thresholds
const (
	timeForNoteLive = 14*time.Minute + 33*time.Second // 14m33s after written
	timeForPostLive = 18*time.Minute + 20*time.Second // 18m20s after post
)

// Task simulation function with timing
func evaluatePost(post string, wg *sync.WaitGroup, startTime time.Time) {
	defer wg.Done()
	processingTime := 3 * time.Second // Simulate processing time
	time.Sleep(processingTime)
	elapsed := time.Since(startTime)
	fmt.Printf("Post evaluated for accuracy (%s): %v\n", post, elapsed)

	if elapsed <= timeForNoteLive {
		fmt.Println("✅ Post evaluation completed within the 14m33s timeframe")
	} else {
		fmt.Println("❌ Post evaluation exceeded the 14m33s timeframe")
	}
}

func checkMediaAuthenticity(media string, wg *sync.WaitGroup, startTime time.Time) {
	defer wg.Done()
	processingTime := 3 * time.Second // Simulate processing time
	time.Sleep(processingTime)
	elapsed := time.Since(startTime)
	fmt.Printf("Media authenticity checked (%s): %v\n", media, elapsed)

	if elapsed <= timeForPostLive {
		fmt.Println("✅ Media check completed within the 18m20s timeframe")
	} else {
		fmt.Println("❌ Media check exceeded the 18m20s timeframe")
	}
}

func scoreNote(note string, wg *sync.WaitGroup, startTime time.Time) {
	defer wg.Done()
	processingTime := 3 * time.Second // Simulate processing time
	time.Sleep(processingTime)
	elapsed := time.Since(startTime)
	fmt.Printf("Note scored (%s): %v\n", note, elapsed)

	if elapsed <= timeForNoteLive {
		fmt.Println("✅ Note scoring completed within the 14m33s timeframe")
	} else {
		fmt.Println("❌ Note scoring exceeded the 14m33s timeframe")
	}
}

func main() {
	// Simulate post, media, and note
	post := "Sample Post"
	media := "Sample Image/Video"
	note := "Proposed Note"

	// Track the start time
	startTime := time.Now()

	// WaitGroup to handle concurrency
	var wg sync.WaitGroup
	wg.Add(3)

	// Concurrently process tasks
	go evaluatePost(post, &wg, startTime)
	go checkMediaAuthenticity(media, &wg, startTime)
	go scoreNote(note, &wg, startTime)

	// Wait for all tasks to complete
	wg.Wait()

	// Display the total processing time
	fmt.Printf("All tasks completed in %v\n", time.Since(startTime))
}
