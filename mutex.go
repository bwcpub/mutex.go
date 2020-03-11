package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "time"
)

func main() {

    var mtx sync.Mutex

    rand.Seed(time.Now().Unix())    // .UnixNano() doesn't seem to work well under Windows.

    mtx.Lock()
    for i := 0 ; i < 5 ; i++ {
        go waitRnd(&mtx, i)
    }
    fmt.Println("main() waiting for all threads to finish...")
    mtx.Unlock()

    for runtime.NumGoroutine() > 1 {                    // main is also a goroutine
        time.Sleep(time.Duration(1) * time.Second)      // No reason to chew up CPU time while waiting.
    }
    fmt.Println("main() exiting")
}

func waitRnd(mtx *sync.Mutex, threadId int) {

    defer mtx.Unlock()                // Guarantees mutex unlock no matter how this func returns

    waitSecs := rand.Intn(9) + 1
    mtx.Lock()
    fmt.Printf("Thread %v sleeping %v secs... ", threadId, waitSecs)
    time.Sleep(time.Duration(waitSecs) * time.Second)
    fmt.Println("done");
}
