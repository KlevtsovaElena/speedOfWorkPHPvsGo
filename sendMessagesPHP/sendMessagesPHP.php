<?php

	// основные переменные для url
	$host = "https://api.telegram.org/";
	$token = "bot6131123688:AAGV7bDvX4aX4_n-ShaiKjXlpUvlnfXsQFY";
	$chatId = 1752911328;
	$text = "testTextFromPHP";
	$getString = "/sendMessage?chat_id=" . $chatId . "&text=" . $text;

	$url = $host . $token . $getString;

	// количество отправляемых сообщений
	$messagesNumber = 10;

	// ****ОТПРАВКА СООБЩЕНИЙ В ОБЫЧНОМ РЕЖИМЕ***
	// время начала проги в наносекундах
	$startTime = system('date +%s%N');
	// отправка сообщений в бота в могопоточном режиме
	for ($i = 0; $i < $messagesNumber; $i++) {
		file_get_contents($url);
	}

	// время конца проги в наносекундах
	$finishTime = system('date +%s%N');
	echo $messagesNumber . " сообщений PHP отправил за " . ($finishTime - $startTime) / 1000000000 . " секунд";

?>