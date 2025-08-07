package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
    //"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

/////////////////////////////////////////////////////////////////////////////////////////////////
//Реализация хэндлера по запросу "/"/////
func Handler1(res http.ResponseWriter, req *http.Request)  {
//открыть файл 
data, err := os.ReadFile("../index.html")
    if err != nil {
        log.Fatal(err)
    }
//res.Header().Set("Content-Type", "multipart/form-data")
//возвращается строка в res.Write([]byte(s))
res.Write([]byte(data))
//fmt.Println(string(data))
}
///////////////////////////////////////////////////////////////////////////////////////////////////
// Реализация хэндлера по запросу "/upload"
func Handler2(res http.ResponseWriter, req *http.Request) {
// Парсинг формы////
    if err := req.ParseForm(); err != nil {
        http.Error(res, "Ошибка при парсинге формы", http.StatusInternalServerError)
        return
    }

// Получение файла из формы////////////////////////////////////////////////////////////////////////////
    file,H, err := req.FormFile("myFile") // "myFile" - название поля файла в форме
    fmt.Println(file,H,err)
    log.Fatal(err)
    if err != nil { 
        http.Error(res, "Не удалось получить файл", http.StatusInternalServerError)
        return
    }
fmt.Println("222")
    defer file.Close() //Закрываем файл/////

// Чтение данных из файла////////////////////
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(res, "Ошибка при чтении файла", http.StatusInternalServerError)
        return
    }
//  конвертация строки//////////////////
fmt.Println("333")
   convString:=service.Texttomorse_and_revers(string(data))
// Создаем новый файл///////////////////
	urlString := req.URL.String()
    file2, err := os.Create(fmt.Sprintln(time.Now().UTC().String()+filepath.Ext(urlString)))
    if err != nil {
        log.Fatal(err)
    }
    defer file2.Close()

//Записываем реультат конвертации строки в файл file2////
    _, err = file2.WriteString(convString)
    if err != nil {
        log.Fatal(err)
    }
    _,err=res.Write([]byte(convString))
    if err != nil {
        log.Fatal(err)}
    }