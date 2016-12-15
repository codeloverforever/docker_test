package main

import (
    "fmt"
)

func split (sum int ) (x,y int) {
    x = sum * 4 / 9
    y = sum - x

    return 
}

func main(){
    sum := 1
    for ;sum < 50; {
        sum += sum
    }
    fmt.Println(sum)

}
