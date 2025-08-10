package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	http "net/http"
	"time"
)

// Структура сервера
type Serv struct {
	loggerServ *log.Logger
	HTTPServer http.Server
}

func Myserver(logger *log.Logger) *Serv {
	router := http.NewServeMux() // Маршрутизатор////////////////////////////////////////
	// Регистрируем хэндлеры для различных путей/////////////////////////////////////////
	router.HandleFunc("/", handlers.Handler1)       // вызов хэндлера по ....."/"
	router.HandleFunc("/upload", handlers.Handler2) // вызов хэндлера по ....."/upload"
	// Создаем экземпляр для настройки роутера///////////////////////////////////////////
	S := &Serv{
		loggerServ: logger,
		HTTPServer: http.Server{
			Addr:                         ":8080",
			Handler:                      router,
			DisableGeneralOptionsHandler: false,
			TLSConfig:                    nil,
			ReadTimeout:                  5 * time.Second,
			ReadHeaderTimeout:            0,
			WriteTimeout:                 10 * time.Second,
			IdleTimeout:                  15 * time.Second,
			MaxHeaderBytes:               0,
			TLSNextProto:                 nil,
			ConnState:                    nil,
			ErrorLog:                     logger,
			BaseContext:                  nil,
			ConnContext:                  nil,
			HTTP2:                        nil,
			Protocols:                    nil,
		},
	}
	return S
}
