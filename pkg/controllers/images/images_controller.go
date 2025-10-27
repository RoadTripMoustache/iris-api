// Package images contains all the methods of the Images controller.
package images

import (
	"fmt"
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/services/images"
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
func UploadImage(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	appConfig := config.GetConfigs()
	if err := ctx.Request.ParseMultipartForm(appConfig.Images.MaxSize); err != nil {
		return nil, errors.New(enum.ImageTooLarge, err)
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		return nil, errors.New(enum.ResourceNotFound, err)
	}
	defer file.Close()

	eerr := images.ExtensionValidation(header.Filename)
	if eerr != nil {
		return nil, eerr
	}

	cfg := config.GetConfigs().Server
	// Ensure destination directory exists
	if err := os.MkdirAll(cfg.ImagesDir, 0o755); err != nil {
		return nil, errors.New(enum.InternalServerError, err)
	}

	// Generate a unique filename
	ext := strings.ToLower(filepath.Ext(header.Filename))
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(cfg.ImagesDir, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		logging.Error(err, nil)
		return nil, errors.New(enum.InternalServerError, err)
	}
	defer dst.Close()

	written, err := io.Copy(dst, file)
	if err != nil {
		logging.Error(err, nil)
		return nil, errors.New(enum.InternalServerError, err)
	}
	if written > appConfig.Images.MaxSize {
		// Safety check if client bypassed MaxBytesReader somehow
		_ = os.Remove(dstPath)
		return nil, errors.New(enum.ImageTooLarge, err)
	}

	// Build public URL using ImagesBaseURL
	base := strings.TrimRight(cfg.ImagesBaseURL, "/")
	publicURL := fmt.Sprintf("%s/%s", base, filename)

	return utils.PrepareResponse(map[string]string{"url": publicURL})
}

// GetImage handles GET /images/{filename}
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
