package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

const (
	baseUrl       = "http://localhost:8081"
	createPostfix = "/notes"
	getPostfix    = "/notes/%d"
)

type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

type Note struct {
	ID        int64    `json:"id"`
	Info      NoteInfo `json:"info"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func createNote() (Note, error) {
	// Создаю новую заметку на удаленном сервере через HTTP POST-запрос. Функция случайным образом генерирует
	// данные для заметки, отправляет их в формате JSON на сервер и возвращает созданную заметку или ошибку.

	note := NoteInfo{
		Title:    gofakeit.BeerName(),
		Context:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	// Получаю байтовый срез с данными, кодированными в формат JSON, они могут быть отправлены через сеть
	data, err := json.Marshal(note)
	if err != nil {
		return Note{}, err
	}

	// Отправляется HTTP POST-запрос на сервер. Данные отправляются в формате JSON, и как тело запроса
	resp, err := http.Post(baseUrl+createPostfix, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return Note{}, err
	}
	defer resp.Body.Close()

	// Проверяется что вернул сервис, если 201, то все создано
	if resp.StatusCode != http.StatusCreated {
		return Note{}, err
	}

	// Возвращаю созданную заметку, декодирую ответ от сервера
	var createdNote Note
	// Ответ от сервера(resp.Body) при http-запросе передается байтами в json-формате, помещаю(декодирую) их в объект
	if err = json.NewDecoder(resp.Body).Decode(&createdNote); err != nil {
		return Note{}, err
	}

	return createdNote, nil
}

func getNote(id int64) (Note, error) {
	// Метод отправляет HTTP GET-запрос на сервер для получения заметки по её идентификатору

	resp, err := http.Get(fmt.Sprintf(baseUrl+getPostfix, id))
	if err != nil {
		log.Fatal("Failed to get note:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Note{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Note{}, errors.Errorf("failed to get note: %d", resp.StatusCode)
	}

	var note Note
	if err = json.NewDecoder(resp.Body).Decode(&note); err != nil {
		return Note{}, err
	}

	return note, nil
}

func main() {
	note, err := createNote() // создаю новую заметку на удаленном сервере
	if err != nil {
		log.Fatal("failed to create note:", err)
	}

	log.Printf(color.RedString("Note created:\n"), color.GreenString("%+v", note))

	note, err = getNote(note.ID) // получаю заметку с сервера по айди
	if err != nil {
		log.Fatal("failed to get note:", err)
	}

	log.Printf(color.RedString("Note info got:\n"), color.GreenString("%+v", note))
}