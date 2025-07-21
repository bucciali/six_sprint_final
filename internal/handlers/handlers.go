package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, "..", "index.html")
	http.ServeFile(w, r, path)

}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "mistake with getting file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "mistake with reading file", http.StatusInternalServerError)
	}

	datastr := service.Morze(string(data))

	safeTime := time.Now().UTC().Format("20060102_150405")
	filename := safeTime + filepath.Ext(header.Filename)

	err = os.WriteFile(filename, []byte(datastr), 0755)
	if err != nil {
		http.Error(w, "Mistake with writing file", http.StatusInternalServerError)

		return
	}

	w.Write([]byte(datastr))

}
