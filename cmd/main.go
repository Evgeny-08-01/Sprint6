package main

import (
	"log"
	"os"

	//"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	
	// создаём файл info.log и обрабатываем ошибку, если что-то пошло не так
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// откладываем закрытие файла
	defer file.Close()
	//1й аргумент: устанавливаем вывод в file, 2й аргумент: префикс-""(.....), 3й аргумент: формат вывода в файл
	logger := log.New(file, ".....", log.Ldate|log.Ltime|log.Lshortfile)
	//handlers.Logger = logger
	//  Вызываем сервер Myserver , передаем созданный логгер, получаем экземпляр сервера и настраиваем его
	Srv := server.Myserver(logger)
    // Запуск сервера///////////////////////////////////////////////////////////////////////////////
    err=Srv.HTTPServer.ListenAndServe() 
    if err != nil {
		// Если при запуске сервера возникают ошибки, выводим их с помощью логгера на уровне Fatal
		logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
