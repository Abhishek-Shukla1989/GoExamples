package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ##########################
// # SIMPLE CHANNEL COMM    #
// ##########################
func SimpleChannelCommunication() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	value := <-ch
	fmt.Println("SimpleChannelCommunication - Received value:", value)
}

// ##########################
// # BUFFERED CHANNELS      #
// ##########################
func BufferedChannels() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	fmt.Println("BufferedChannels - Received values:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// ##########################
// # CHANNEL DIRECTION      #
// ##########################
func sendOnly(ch chan<- int, value int) {
	ch <- value
}

func receiveOnly(ch <-chan int) {
	value := <-ch
	fmt.Println("ChannelDirection - Received from receiveOnly:", value)
}

func ChannelDirection() {
	ch := make(chan int)
	go sendOnly(ch, 100)
	receiveOnly(ch)
}

// ##########################
// # CHANNEL CLOSE          #
// ##########################
func ChannelClose() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	fmt.Println("ChannelClose - Iterating over closed channel:")
	for v := range ch {
		fmt.Println(v)
	}
}

// ##########################
// # FAN-OUT, FAN-IN PATTERN#
// ##########################
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("FanOutFanIn - Worker %d processing job %d\n", id, job)
		results <- job * job // e.g., squaring the job number
	}
}

func FanOutFanIn() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	fmt.Println("FanOutFanIn - Results:")
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println(result)
	}
}

// ##########################
// # MULTIPLEXING CHANNELS  #
// ##########################
func MultiplexingChannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("MultiplexingChannels - Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("MultiplexingChannels - Received:", msg2)
	}
}

// ##########################
// # CHANNEL TIMEOUTS       #
// ##########################
func ChannelTimeouts() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 10
	}()

	select {
	case value := <-ch:
		fmt.Println("ChannelTimeouts - Received:", value)
	case <-time.After(1 * time.Second):
		fmt.Println("ChannelTimeouts - Timeout reached!")
	}
}

// ##########################
// # DEADLOCK DETECTION     #
// ##########################
func DeadlockExample() {
	ch := make(chan int)
	// The following line would cause a deadlock if uncommented:
	// fmt.Println(<-ch)
	// Prevent deadlock by sending a value in a goroutine:
	go func() {
		ch <- 1
	}()
	fmt.Println("DeadlockExample - Received:", <-ch)
}

// ##########################
// # CHANNEL OF CHANNELS    #
// ##########################
func ChannelOfChannels() {
	chOfCh := make(chan chan int)

	go func() {
		innerCh := make(chan int)
		go func() {
			innerCh <- 100
		}()
		chOfCh <- innerCh
	}()

	innerCh := <-chOfCh
	fmt.Println("ChannelOfChannels - Value from inner channel:", <-innerCh)
}

// ##########################
// # RATE LIMITING          #
// ##########################
func RateLimiting() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	fmt.Println("RateLimiting - Operations:")
	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("Operation", i)
	}
}

// ##########################
// # PIPELINE PROCESSING    #
// ##########################
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func PipelineProcessing() {
	in := generator(1, 2, 3, 4, 5)
	out := square(in)
	fmt.Println("PipelineProcessing - Squared results:")
	for result := range out {
		fmt.Println(result)
	}
}

// ##########################
// # SELECT WITH DEFAULT    #
// ##########################
func SelectWithDefault() {
	ch := make(chan int)
	select {
	case value := <-ch:
		fmt.Println("SelectWithDefault - Received:", value)
	default:
		fmt.Println("SelectWithDefault - No value received, moving on")
	}
}

// ##########################
// # WAIT FOR GOROUTINES    #
// ##########################
func WaitForGoroutines() {
	done := make(chan bool)

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("WaitForGoroutines - Goroutine 1 done")
		done <- true
	}()
	go func() {
		time.Sleep(700 * time.Millisecond)
		fmt.Println("WaitForGoroutines - Goroutine 2 done")
		done <- true
	}()

	<-done
	<-done
	fmt.Println("WaitForGoroutines - All goroutines finished")
}

// ##########################
// # PRODUCER-CONSUMER      #
// ##########################
func ProducerConsumer() {
	data := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("ProducerConsumer - Producing", i)
			data <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(data)
	}()

	// Consumer
	for d := range data {
		fmt.Println("ProducerConsumer - Consuming", d)
	}
}

// ##########################
// # CHANNEL CLOSING DETECT #
// ##########################
func ChannelClosingDetection() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if v, ok := <-ch; ok {
			fmt.Println("ChannelClosingDetection - Received:", v)
		} else {
			fmt.Println("ChannelClosingDetection - Channel closed")
			break
		}
	}
}

// ##########################
// # CHANNEL CANCELLATION   #
// ##########################
func ChannelCancellation() {
	cancel := make(chan struct{})

	go func() {
		for {
			select {
			case <-cancel:
				fmt.Println("ChannelCancellation - Goroutine cancelled")
				return
			default:
				fmt.Println("ChannelCancellation - Working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(cancel)
	time.Sleep(500 * time.Millisecond)
}

// ##########################
// # CHANNEL ERROR HANDLING #
// ##########################
func ChannelErrorHandling() {
	errCh := make(chan error)

	go func() {
		time.Sleep(200 * time.Millisecond)
		errCh <- fmt.Errorf("ChannelErrorHandling - an error occurred")
	}()

	select {
	case err := <-errCh:
		fmt.Println("ChannelErrorHandling - Error received:", err)
	case <-time.After(1 * time.Second):
		fmt.Println("ChannelErrorHandling - No error received within timeout")
	}
}

func main() {
	fmt.Println("=== Simple Channel Communication ===")
	SimpleChannelCommunication()

	fmt.Println("\n=== Buffered Channels ===")
	BufferedChannels()

	fmt.Println("\n=== Channel Direction ===")
	ChannelDirection()

	fmt.Println("\n=== Channel Close ===")
	ChannelClose()

	fmt.Println("\n=== Fan-out, Fan-in Pattern ===")
	FanOutFanIn()

	fmt.Println("\n=== Multiplexing Channels ===")
	MultiplexingChannels()

	fmt.Println("\n=== Channel Timeouts ===")
	ChannelTimeouts()

	fmt.Println("\n=== Deadlock Detection ===")
	DeadlockExample()

	fmt.Println("\n=== Channel of Channels ===")
	ChannelOfChannels()

	fmt.Println("\n=== Rate Limiting ===")
	RateLimiting()

	fmt.Println("\n=== Pipeline Processing ===")
	PipelineProcessing()

	fmt.Println("\n=== Select with Default ===")
	SelectWithDefault()

	fmt.Println("\n=== Wait for Multiple Goroutines ===")
	WaitForGoroutines()

	fmt.Println("\n=== Producer-Consumer ===")
	ProducerConsumer()

	fmt.Println("\n=== Channel Closing Detection ===")
	ChannelClosingDetection()

	fmt.Println("\n=== Channel Cancellation ===")
	ChannelCancellation()

	fmt.Println("\n=== Channel Error Handling ===")
	ChannelErrorHandling()
}
