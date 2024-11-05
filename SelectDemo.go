package main


import (
    "fmt"
    "time"
)


func main1() {


    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)


    go func(){
        time.Sleep(3* time.Millisecond)
        ch1 <- "Messsage from first one"
    }()


    go func(){
        time.Sleep(1* time.Millisecond)
        ch1 <- "Messsage from second one"
    }()
    go func(){
        time.Sleep(2* time.Millisecond)
        ch1 <- "Messsage from third one"
    }()


    count := 0


    loop:
    for {
        select{
        case msg1:= <-ch1:{
            fmt.Println(msg1)
            count++
            if count == 3{
                break loop
            }
        }
        case msg2:= <-ch2:{
            fmt.Println(msg2)
            count++
            if count == 3{
                break loop
            }
        }
        case msg3:= <-ch3:{
            fmt.Println(msg3)
            count++
            if count == 3{
                break loop
            }
        }
    default:
        fmt.Println("nothing to read")
        }
    }


}
