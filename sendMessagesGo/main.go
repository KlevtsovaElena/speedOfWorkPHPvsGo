package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

	// основные переменные для url
	var host string = "https://api.telegram.org/"
	var token string = "bot6131123688:AAGV7bDvX4aX4_n-ShaiKjXlpUvlnfXsQFY"
	var chatId int = 1752911328
	var text string = "testTextFromGo"
	var getString string = "/sendMessage?chat_id=" + strconv.Itoa(chatId) + "&text=" + text
	var getString2 string = "/sendMessage?chat_id=" + strconv.Itoa(chatId) + "&text=" + "testTextFromGoGooooo"

	url := host + token + getString
	url2 := host + token + getString2

	// количество отправляемых сообщений
	messagesNumber := 10

	// ****ОТПРАВКА СООБЩЕНИЙ В ОБЫЧНОМ РЕЖИМЕ***
	// время начала проги в наносекундах
	startTime := time.Now().UnixNano()

	// отправка сообщений в бота в обычном режиме
	for i := 0; i < messagesNumber; i++ {
		http.Get(url)
	}

	// время конца проги в наносекундах
	finishTime := time.Now().UnixNano()
	fmt.Printf("%v сообщений в обычном режиме Golang отправил за %.9f секунд \n", messagesNumber, float64((float64(finishTime)-float64(startTime))/1000000000))

	// ****ОТПРАВКА СООБЩЕНИЙ В МНОГОПОТОЧНОМ РЕЖИМЕ***
	// время начала проги в наносекундах
	startTime = time.Now().UnixNano()

	// отправка сообщений в бота в могопоточном режиме
	for i := 0; i < messagesNumber; i++ {
		go sendMessage(url2)
	}

	// ждём немного, чтобы все потоки закончились, здесь прикинутое время, примерное на 10-20 сообщений
	time.Sleep(time.Millisecond * 400)

	// время конца проги в наносекундах
	finishTime = time.Now().UnixNano()
	fmt.Printf("%v сообщений в многопоточном режиме Golang отправил менее, чем за %.9f секунд", messagesNumber, float64((float64(finishTime)-float64(startTime))/1000000000))
}

// отправка сообщений
func sendMessage(url string) {
	http.Get(url)
}
