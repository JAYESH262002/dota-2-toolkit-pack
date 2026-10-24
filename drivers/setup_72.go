package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.1" )

func aThUoLzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pg6Q05rOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDYklWKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJIeJVPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxgOWL4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UfaXw5iFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvgKpAZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func El2dMjTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShGGKhF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RH0Da1SIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syoI50YLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTzDP808Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iKCnavVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNc6Xm0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMNRRlWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecBUAdaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzIouDFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6Sn8J6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZyst9ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIKinnIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhDxoBrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fLOVgIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYhXPHHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxeiWYjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHHzRvHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOHFXNzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yR01HMouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDXF4DtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2JX92TCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jq6lCRG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CmeOE55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1FwEMXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPVGt3dpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCSsDp4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HedVerbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsF4yWjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3DfMJTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5N9ztfkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UL06qNpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUxOtWZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DQpBLrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sh6JxjrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tq1LsxAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01VfE3cQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1nuoFzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6JzAY3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTf6dA65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ug6Xto6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlxEgRriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syFXLSxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj1lk3BfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFnbCLUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VexGitFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fp789mfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vJOIVYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZ3NC7OKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSI7jeUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yA1x42s1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEHDUE8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8y9QkX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjJ3QBytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQ97ua5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVTkVEVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w1K3tIkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFV3dLZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUMaDBd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

