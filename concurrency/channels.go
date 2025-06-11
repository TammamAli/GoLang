package concurrency

import (
	"fmt"
	"reflect"
	"time"
)

func Channels() {
	/*
	   What are channels
	   Channels can be thought of as pipes using which Goroutines communicate.
	   Similar to how water flows from one end to another in a pipe, data can be sent from one end and received
	   from the other end using channels.


	   Declaring channels
	   Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport.
	   No other type is allowed to be transported using the channel.

	   chan T is a channel of type T

	   The zero value of a channel is nil.
	   nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.

	   As usual, the short hand declaration is also a valid and concise way to define a channel.
	   b := make(chan int)

	   Sending and receiving from a channel
	   The syntax to send and receive data from a channel is given below,
	   The direction of the arrow with respect to the channel specifies whether the data is sent or received.
	   In the first line, the arrow points outwards from a and hence we are reading from channel a and storing
	   the value to the variable data. In the second line, the arrow points towards a and hence we are writing to channel a.

	   data := <- a // read from channel a
	   a <- data // write to channel a


	   Sends and receives are blocking by default
	   Sends and receives to a channel are blocking by default What does this mean?
	   When data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel.
	   Similarly, when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.

	   This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or
	   conditional variables that are quite common in other programming languages.
	   The upcoming sections will add more clarity on how channels are blocking by default.


	   In the above program, we create a done bool channel in line no. 12 and pass it as a parameter to the hello Goroutine.
	   In line no. 14 we are receiving data from the done channel. This line of code is blocking which means that until
	   some Goroutine writes data to the done channel, the control will not move to the next line of code.
	   Hence this eliminates the need for the time.Sleep which was present in the original program to prevent
	   the main Goroutine from exiting.

	   The line of code <-done receives data from the done channel but does not use or store that data in any variable.
	   This is perfectly legal.

	   Now we have our main Goroutine blocked waiting for data on the done channel. The hello Goroutine receives this channel
	   as a parameter, prints Hello world goroutine and then writes to the done channel. When this write is complete,
	   the main Goroutine receives the data from the done channel, it is unblocked and then the text main function is printed.

	   Deadlock
	   One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel,
	   then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the
	   program will panic at runtime with Deadlock.

	   Similarly, if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write
	   data on that channel, else the program will panic.



	   Unidirectional channels
	   All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them.
	   It is also possible to create unidirectional channels, that is channels that only send or receive data.

	   All is well but what is the point of writing to a send only channel if it cannot be read from!

	   This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send
	   only or receive only channel but not the vice versa.

	   Closing channels and for range loops on channels
	   . Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
	   . Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
	   				v, ok := <- ch
	   The for range form of the for loop can be used to receive values from a channel until it is closed.






	*/

	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	b := make(chan int)
	fmt.Printf("b= %v : %v ,type is %T,  %v", b, &b, b, reflect.ValueOf(b).Kind())

}

func channel_s0(done chan bool) {
	fmt.Println("Hello world goroutine")
	fmt.Println(time.Now().Format(time.DateTime))
	done <- true            //write to channel a  // blocking the code
	time.Sleep(time.Second) // iqnore from here
	fmt.Println("End Write")

}

func channelS1(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true //write to channel & blocking the code
}

func CallChannelS1() {
	done := make(chan bool)
	go channelS1(done)
	b := <-done //read from channel & blocking the code
	fmt.Printf("chanel value is %v", b)
	fmt.Printf("\nEnd of main function")
}

func channelS2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func CallChannelS2() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go channelS2(done)
	<-done
	fmt.Println("Main received data")

	/*
		<- chan name  read
		chan name <-  value  wirte
	*/
}

func CalcSquares(number int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}

	return sum
}

func CalCubes(number int) int {

	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	return sum
}
func Chan_CalcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum

}
func Chan_CalCubes(number int, cubeeop chan int) {

	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeeop <- sum
}

func sendOnlyChan(sendch chan<- int) {
	sendch <- 10
}

func receiveOnlyChan(receiveOnly <-chan int) {
	ch := <-receiveOnly
	fmt.Println(ch)
}

func CallSendOnlyChan() {
	sendch := make(chan<- int)
	go sendOnlyChan(sendch)

	//r := <-sendch     //|_ invalid operation: cannot receive from send-only channel sendch (variable of type chan<- int)

	r := 0
	fmt.Println(r)

	chnl := make(chan int)
	go sendOnlyChan(chnl) //undefined: sendData
	fmt.Println(<-chnl)
}

func CallReceiveOnlyChan() {
	sendch := make(<-chan int)
	go receiveOnlyChan(sendch)

	//r := <-sendch     //|_ invalid operation: cannot receive from send-only channel sendch (variable of type chan<- int)

	r := 0
	fmt.Println(r)

	chnl := make(chan int)
	go sendOnlyChan(chnl) //undefined: sendData
	fmt.Println(<-chnl)
}

func producers(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func Receiver() {
	ch := make(chan int)
	go producers(ch)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Receive: %v\n", v)
	}
}

func ReceiverForRange() {
	fmt.Println("ReceiverForRange")
	ch := make(chan int)
	go producers(ch)

	for v := range ch {
		fmt.Printf("Receive: %v\n", v)
	}
}

func digits(num int, ch chan int) {

	for num != 0 {
		ch <- num % 10
		num /= 10
	}
	close(ch)
}

func CalcSquares2(num int, squreChan chan int) {

	result := 0
	ch := make(chan int)
	go digits(num, ch)
	for v := range ch {
		result += v * v
	}
	squreChan <- result
}

func CalCubes2(num int, cubeChan chan int) {

	result := 0
	ch := make(chan int)
	go digits(num, ch)
	for v := range ch {
		result += v * v * v
	}
	cubeChan <- result

}

func TestCall(number int) {
	sqrch := make(chan int)
	cubech := make(chan int)
	go CalcSquares2(number, sqrch)
	go CalCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("squares: ", squares)
	fmt.Println("cubes: ", cubes)

	fmt.Println("Final output", squares+cubes)
}
