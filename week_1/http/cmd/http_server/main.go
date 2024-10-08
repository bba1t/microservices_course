package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi"
)

const (
	baseUrl       = "localhost:8081"
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
	ID        int64     `json:"id"`
	Info      NoteInfo  `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SyncMap struct {
	elems map[int64]*Note
	m     sync.RWMutex
}

var notes = &SyncMap{
	elems: make(map[int64]*Note),
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	info := &NoteInfo{}
	// данные в теле http-запроса передаются байтами в json-формате, помещаю(декодирую) их в объект
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		http.Error(w, "Failed to decode note data", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	note := &Note{ //полностью заполняю дынные
		ID:        rand.Int63(),
		Info:      *info,
		CreatedAt: now,
		UpdatedAt: now,
	}

	w.Header().Set("Content-Type", "application/json") // какой тип данных я отправляю
	w.WriteHeader(http.StatusCreated)

	// структуру note превращаю в json, помещаю ее в ответ и отправляю пользователю
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}

	// помещаю данные в основной список
	notes.m.Lock() //  блокирую доступ к elems для других горутин
	notes.elems[note.ID] = note
	defer notes.m.Unlock()

}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	// В отличие от post запроса, get имеет request параметры, но не имеет body
	// И в таком параметре будет приходить id пользователя

	noteID := chi.URLParam(r, "id")
	id, err := parseNoteID(noteID)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// RLock запрещает изменять объект, но дает его читать
	notes.m.RLock()

	note, ok := notes.elems[id]
	if !ok {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	defer notes.m.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}
}

func parseNoteID(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func main() {
	// Можно использовать базовый net/http, а можно какой усовершенствованный роутер
	r := chi.NewRouter()

	// Post/Get - эндпоинты
	// первый аргумент - паттерн, как достучаться до ручки
	// второй - хендлер, у которого 1 параметр куда записать, 2 откуда записать
	r.Post(createPostfix, createNoteHandler)
	r.Get(getPostfix, getNoteHandler)

	err := http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}
