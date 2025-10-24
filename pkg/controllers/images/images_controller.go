// Package images contains all the methods of the Images controller.
package images

import (
	"encoding/json"
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/constantes"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"github.com/gorilla/mux"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var filenameSafeRegexp = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

// UploadImage handles POST /v1/images
// Expects multipart/form-data with field name "file"
// TODO : Refacto
func UploadImage(w http.ResponseWriter, r *http.Request) {
	// Enforce max upload size: 2MB
	r.Body = http.MaxBytesReader(w, r.Body, constantes.MaxImageSizeBytes)

	if err := r.ParseMultipartForm(constantes.MaxImageSizeBytes); err != nil {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		_, _ = w.Write([]byte("file too large"))
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("missing file field"))
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".png" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid file type; only .jpg and .png are allowed"))
		return
	}

	cfg := config.GetConfigs().Server
	// Ensure destination directory exists
	if err := os.MkdirAll(cfg.ImagesDir, 0o755); err != nil {
		logging.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate a unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(cfg.ImagesDir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		logging.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	written, err := io.Copy(dst, file)
	if err != nil {
		logging.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if written > constantes.MaxImageSizeBytes {
		// Safety check if client bypassed MaxBytesReader somehow
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		_, _ = w.Write([]byte("file too large"))
		_ = os.Remove(dstPath)
		return
	}

	// Build public URL using ImagesBaseURL
	base := strings.TrimRight(cfg.ImagesBaseURL, "/")
	publicURL := fmt.Sprintf("%s/%s", base, filename)

	resp := map[string]string{"url": publicURL}
	b, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(b)
}

// GetImage handles GET /images/{filename}
// TODO : Refacto
func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["filename"]
	if name == "" || !filenameSafeRegexp.MatchString(name) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cfg := config.GetConfigs().Server
	path := filepath.Join(cfg.ImagesDir, name)
	f, err := os.Open(path)
	fmt.Println(path)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logging.Error(err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Determine content-type
	ext := strings.ToLower(filepath.Ext(name))
	ct := mime.TypeByExtension(ext)
	if ct == "" {
		if ext == ".png" {
			ct = "image/png"
		} else {
			ct = "image/jpeg"
		}
	}
	w.Header().Set("Content-Type", ct)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, f)
}
