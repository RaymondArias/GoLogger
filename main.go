package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	sleepDur, found := os.LookupEnv("SLEEP_DURATION")
	if !found {
		sleepDur = "1s"
	}

	appID, found := os.LookupEnv("APP_ID")
	if !found {
		appID = "A"
	}

	threads, found := os.LookupEnv("THREADS")
	if !found {
		threads = "1"
	}
	numThread, err := strconv.Atoi(threads)
	if err != nil {
		panic(fmt.Sprintf("cannot parse %s error: %s", threads, err.Error()))
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for i := 0; i < numThread; i++ {
		go log(ctx, sleepDur, i, appID)
	}
	log(ctx, sleepDur, -1, appID)

}

// Output simply outputs
type Output struct {
	Message   string
	Timestamp string
	Thread    string
	SeqNum    int
	AppID     string
}

func log(ctx context.Context, sleepDur string, thread int, appID string) {
	duration, err := time.ParseDuration(sleepDur)
	if err != nil {
		panic(fmt.Sprintf("cannot parse %s error: %s", sleepDur, err.Error()))
	}
	threadStr := fmt.Sprintf("%d", thread)
	seqNum := 0
	for {
		select {
		case <-ctx.Done():
			return // returning not to leak the goroutine
		default:
			for {
				data := Output{
					Message:   fmt.Sprintf("data %d", rand.Int()),
					Timestamp: time.Now().String(),
					Thread:    threadStr,
					SeqNum:    seqNum,
					AppID:     appID,
				}
				jsonPayload, _ := json.Marshal(data)
				fmt.Printf("%s\n", string(jsonPayload))
				time.Sleep(duration)
				seqNum++
			}

		}
	}
}
