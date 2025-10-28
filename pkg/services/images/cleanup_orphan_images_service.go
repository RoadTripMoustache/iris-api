package images

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	nosql "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// CleanupOrphanImages scans the images folder and removes files not referenced in the DB
func CleanupOrphanImages() {
	defer func() {
		if r := recover(); r != nil {
			logging.Info("panic recovered in CleanupOrphanImages", map[string]interface{}{"service": "images", "method": "CleanupOrphanImages", "panic": r})
		}
	}()

	cfg := config.GetConfigs().Server
	imagesDir := cfg.ImagesDir
	if imagesDir == "" {
		return
	}

	// Build the set of referenced filenames from DB (ideas + comments)
	referenced := make(map[string]struct{}, 256)
	baseURL := strings.TrimRight(cfg.ImagesBaseURL, "/")
	// Fetch all ideas (no pagination)
	var nilLimit *int = nil
	var nilOffset *int = nil
	docs := nosql.GetInstance().GetDocumentsOrderBy("ideas", bson.D{{Key: "created_at", Value: -1}}, nilLimit, nilOffset, []nosqlUtils.Filter{})
	for _, d := range docs {
		b, _ := bson.Marshal(d)
		var idea struct {
			Images   []string `bson:"images"`
			Comments []struct {
				Images []string `bson:"images"`
			} `bson:"comments"`
		}
		_ = bson.Unmarshal(b, &idea)
		// idea images
		for _, u := range idea.Images {
			if name, ok := normalizeToFilename(u, baseURL); ok {
				referenced[name] = struct{}{}
			}
		}
		// comments images
		for _, c := range idea.Comments {
			for _, u := range c.Images {
				if name, ok := normalizeToFilename(u, baseURL); ok {
					referenced[name] = struct{}{}
				}
			}
		}
	}

	// Walk through the images directory and remove files not referenced
	_ = filepath.WalkDir(imagesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip entries that cause errors
			return nil
		}
		if d.IsDir() {
			return nil
		}
		name := d.Name()
		if name == "" || !filenameSafeRegexpSvc.MatchString(name) {
			return nil
		}
		if _, exists := referenced[name]; !exists {
			if removeErr := os.Remove(path); removeErr != nil {
				logging.Error(removeErr, map[string]interface{}{"service": "images", "method": "CleanupOrphanImages", "file": name})
			} else {
				logging.Info("deleted orphan image", map[string]interface{}{"file": name})
			}
		}
		return nil
	})
}

// normalizeToFilename converts a stored image reference (URL or filename) to a server filename.
// It returns (filename, true) if the filename is considered safe; otherwise ("", false).
func normalizeToFilename(ref string, baseURL string) (string, bool) {
	if ref == "" {
		return "", false
	}
	// Trim query or fragment if any
	if i := strings.IndexAny(ref, "?#"); i >= 0 {
		ref = ref[:i]
	}
	// If it starts with baseURL, strip it
	if baseURL != "" {
		if strings.HasPrefix(ref, baseURL+"/") {
			ref = strings.TrimPrefix(ref, baseURL+"/")
		}
	}
	// If it's a full URL, keep the last segment
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") || strings.HasPrefix(ref, "/") {
		parts := strings.Split(ref, "/")
		ref = parts[len(parts)-1]
	}
	// Final safety check
	if filenameSafeRegexpSvc.MatchString(ref) {
		return ref, true
	}
	return "", false
}
