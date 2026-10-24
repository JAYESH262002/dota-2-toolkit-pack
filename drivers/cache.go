package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.4" )

func V59pPPQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggAq3mdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33dFcNlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9fiEdNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJdDRhDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKUlSK71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ky30kAiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xvi9VGaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHskvFA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HIrZy9OwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvdVsOnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLNp0ppxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFvSo771Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60pHVbHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YOinPLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjzvCpTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12ZHc1QgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufR59PC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nh2BTGfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GmA1cCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8maNdbeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsGfWGDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyLKcdX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWW5XAcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DLPN9ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsGhjjSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Wrs6N3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEe2FdeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQ5dEIhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zztvZjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHSFhc4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWdujRy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cm1VHP31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2xBZoFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tdjRedIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ul47qRWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70t7LdI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RC3USL4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrqeiP6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tq2AuvAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajPi6daYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PmiWtn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lNILAzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyLfTQofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtT8MZ1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRfXzSerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4duEsdv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMNEnxMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2ZRVJ80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3A9hiUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmhN2xdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdZq2PGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhV2iMwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWrYMc1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p76fsa0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tK6AAIzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOHQGtvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBCGXs8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b68gN2E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WEiGDMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvElSZf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyQaOCisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUOjngqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9Bm2q9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3fELXC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISYLbNzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhhBFLWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxKNzzceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEw1YRz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3peJNo3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDgkO8FFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVgK5x9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQaM7hdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1Ej77cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utqEEyqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIqgn7MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDuioatnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4nHoowHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqQX8YWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJLABicDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoU3yzNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfHeE0C8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdS2lo0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXrqCFdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJjTjU3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycCYWlx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpq9YTEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkjZKyvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlRDtDI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gesEL5mkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOUF4JTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bM4F2gKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXwOhxevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBPFtPzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdmw7VdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDkZmVCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQFvfaDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivC2ezz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idhO7g8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZrGsHhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2pR4QAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRzZHsPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMJQ5lAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKdvlIAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E51Twva1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyCAwkLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tX8zwIVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMuMxy16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3LHNroPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKP5o0RNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHjcR5p9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yza3wQFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPIEefKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6AMUtXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpGHdUZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xye5URs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEFlx9v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtoTmhlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4eHqqDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXXwzESvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oazj58HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5G3egLvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTUOSW5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqn04PP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7lWIFXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A154wCQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCthTDR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgUNF3XOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMDkTRlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXOBcgKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUplYGBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7IgTkmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHJV8U5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsiD7HTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkgrP39eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdxSMCU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTarCfX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7y4NxpDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXKOrpJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UUVlCPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GxCDwk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESxuzS6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5HjvSWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBO2D9qaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vkkc32ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACvHFtNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRpKCgrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

