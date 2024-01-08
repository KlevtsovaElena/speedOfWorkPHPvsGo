package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	fmt.Println(messagesNumber, "сообщений в обычном режиме Golang отправил за", getSecond(strconv.FormatInt(finishTime-startTime, 10)), "секунд")

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
	fmt.Println(messagesNumber, "сообщений в многопоточном режиме Golang отправил менее, чем за", getSecond(strconv.FormatInt(finishTime-startTime, 10)), "секунд")

}

// отправка сообщений
func sendMessage(url string) {
	http.Get(url)
}

// представим наносекунды в секунды для удобства чтения
func getSecond(timeString string) string {
	// посчитаем сколько символов в строке
	countSymbols := len(timeString)
	seconds := ""
	switch {
	case countSymbols == 9:
		seconds = "0." + timeString
	case countSymbols < 9:
		nulls := ""
		for i := 0; i < 9-countSymbols; i++ {
			nulls += "0"
		}
		seconds = "0." + nulls + timeString
	case countSymbols > 9:
		intPart := strings.TrimSpace(timeString[0 : countSymbols-9])
		floatPart := strings.TrimSpace(timeString[countSymbols-9 : countSymbols])
		seconds = intPart + "." + floatPart
	}
	return seconds
}
