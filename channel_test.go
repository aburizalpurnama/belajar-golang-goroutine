package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// membuat sebuah function yang mengembalikan sebuah nilai
func getName() string {
	return "Moh. Aburizal Purnama"
}

func TestChannelTest(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		// mengirim data ke channel
		channel <- getName()
	}()

	// // menerima data dari channel (menyimpan kedalam variable)
	// data := <-channel

	// fmt.Println(data)

	// menerima data dari channel (langsung diparsing sebagai parameter)
	fmt.Println(<-channel)

	// close channel
	defer close(channel)
}

func SendDataToChannel(channel chan<- string, data string) {
	time.Sleep(1 * time.Second)
	channel <- data
}

func GetDataFromChannel(channel <-chan string) string {
	return <-channel
}

func TestChannelAsParameter(t *testing.T) {

	// create channel
	channel := make(chan string, 1)

	// parsing channel as parameter, and send value to channel
	go SendDataToChannel(channel, getName())

	// parsing channel as parameter, and receive value from channel
	go GetDataFromChannel(channel)

	// close channel
	defer close(channel)

	time.Sleep(5 * time.Second)
}

func TestBuufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Abu"
		channel <- "Rizal"
		channel <- "Purnama"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println("Channel capacity : ", cap(channel))
		fmt.Println("Channel Size : ", len(channel))
	}()

	fmt.Println("Ups")

	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulanan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	// buat 2 channel
	channel1 := make(chan string)
	channel2 := make(chan int)
	defer close(channel1)
	defer close(channel2)

	// masukkan data pada channel
	go SendDataToChannel(channel1, "rizal")

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- 98
	}()

	// Lakukan infinite loop dan select channel
	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel1 :", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data from channel2 :", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestSelectChannelWithDefault(t *testing.T) {
	// buat 2 channel
	channel1 := make(chan string)
	channel2 := make(chan int)
	defer close(channel1)
	defer close(channel2)

	// masukkan data pada channel
	go SendDataToChannel(channel1, "rizal")

	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- 98
	}()

	// Lakukan infinite loop dan select channel
	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel1 :", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data from channel2 :", data)
			counter++

		default:
			fmt.Println("Waiting for data..")
		}

		if counter == 2 {
			break
		}
	}

}
