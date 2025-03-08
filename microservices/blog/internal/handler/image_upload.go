// /home/nurlashko/work/blog/microservices/blog/internal/handler/image.go
package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	auth "nurlashko.dev/auth/client"
	"nurlashko.dev/blog/internal/middleware"
)

const (
	UploadPath = "/www/data/images/"
)

func ImageUpload(auth *auth.AuthClient) http.HandlerFunc {
	return middleware.AuthenticationMiddleware(auth, func(w http.ResponseWriter, r *http.Request) {
		// Get file from form
		fmt.Println(r)
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Invalid file", http.StatusBadRequest)
			slog.Error("upload: invalid file", "err", err)
			return
		}
		defer file.Close()

		// Check file type (basic check, consider a more robust solution)
		contentType := fileHeader.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "image/") {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			return
		}

		// Create uploads directory if it doesn't exist
		if err := os.MkdirAll(UploadPath, os.ModePerm); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			slog.Error("upload: cannot create directory", "err", err)
			return
		}

		// Generate a unique filename
		ext := filepath.Ext(fileHeader.Filename)
		newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		fullPath := filepath.Join(UploadPath, newFilename)

		// Create the file on the server
		dst, err := os.Create(fullPath)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			slog.Error("upload: cannot create file", "err", err)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the server file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			slog.Error("upload: cannot save file", "err", err)
			return
		}

		remoteImageAddr := strings.Replace(fullPath, UploadPath, "https://static.nurlashko.dev/images/", 1)
		response := map[string]string{"path": remoteImageAddr}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
			slog.Error("upload: cannot write JSON response", "err", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
