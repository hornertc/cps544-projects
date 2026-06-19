// Package cas provides code for the Content Addressable Storage system
package cas

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var validAlgorithm = map[string]func() hash.Hash{
	"sha256":     sha256.New,
	"sha384":     sha512.New384,
	"sha512":     sha512.New,
	"sha512-224": sha512.New512_224,
	"sha512-256": sha512.New512_256,
}

// FileStats struct
type FileStats struct {
	Count  int     `json:"count"`
	Mean   float64 `json:"mean"`
	Stddev float64 `json:"stddev"`
}

// CAS struct
type CAS struct {
	mux        *http.ServeMux
	storageDir string
	algorithms map[string]func() hash.Hash
}

// ErrInvalidDigest error message
var ErrInvalidDigest = errors.New("invalid digest name")

// New function
func New(dir string, algorithms ...string) (*CAS, error) {
	if len(algorithms) == 0 {
		return nil, errors.New("at least one algorithm must be provided")
	}

	algMap := make(map[string]func() hash.Hash, len(algorithms))
	for _, alg := range algorithms {
		h := validAlgorithm[alg]
		if h == nil {
			return nil, ErrInvalidDigest
		}
		err := os.MkdirAll(filepath.Join(dir, alg), 0777)
		if err != nil {
			return nil, fmt.Errorf("creating digest directory: %w", err)
		}
		algMap[alg] = h
	}

	// if err := atomicMkdir(dir); err != nil {
	// 	return nil, fmt.Errorf("error creating storage directory: %s", err)
	// }

	cas := &CAS{
		storageDir: dir,
		algorithms: algMap,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/blob/", cas.handleBlob)
	mux.HandleFunc("/stats/", cas.handleStats)
	cas.mux = mux

	return cas, nil
}

// func validAlgorithm(algorithm string) (crypto.Hash, bool) {
// 	switch algorithm {
// 	case "sha256":
// 		return crypto.SHA256, true
// 	case "sha384":
// 		return crypto.SHA384, true
// 	case "sha512":
// 		return crypto.SHA512, true
// 	case "sha512-224":
// 		return crypto.SHA512_224, true
// 	case "sha512-256":
// 		return crypto.SHA512_256, true
// 	default:
// 		return 0, false
// 	}
// }

func (cas *CAS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cas.mux.ServeHTTP(w, r)
	// switch r.Method {
	// case http.MethodPost:
	// 	cas.handleBlobUpload(w, r)
	// case http.MethodGet:
	// 	if strings.HasPrefix(r.URL.Path, "/blob/") {
	// 		cas.handleBlobRetrieval(w, r)
	// 	} else if r.URL.Path == "/stats" {
	// 		cas.handleStats(w, r)
	// 	} else {
	// 		http.Error(w, "Not Found", http.StatusNotFound)
	// 	}
	// case http.MethodHead:
	// 	if strings.HasPrefix(r.URL.Path, "/blob/") {
	// 		cas.handleBlobRetrievalHead(w, r)
	// 	} else {
	// 		http.Error(w, "Not Found", http.StatusNotFound)
	// 	}
	// default:
	// 	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	// }
}

func (cas *CAS) handleBlob(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodHead:
		cas.handleBlobRetrieval(w, r)
	case http.MethodPost:
		cas.handleBlobUpload(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (cas *CAS) handleBlobUpload(w http.ResponseWriter, r *http.Request) {
	tmpFile, err := os.CreateTemp("", "cas-upload-")
	if err != nil {
		http.Error(w, fmt.Errorf("Error creating temporary file: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()
	// 	err := tmpFile.Close()
	// 	if err != nil {
	// 		http.Error(w, fmt.Errorf("Error closing temporary file: %w", err).Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	err = os.Remove(tmpFile.Name())
	// 	if err != nil {
	// 		http.Error(w, fmt.Errorf("Error removing temporary file %w", err).Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }()

	names := make([]string, 0, len(cas.algorithms))
	hashes := make([]hash.Hash, 0, len(cas.algorithms))
	writers := make([]io.Writer, 0, len(cas.algorithms)+1)
	for name, genHash := range cas.algorithms {
		h := genHash()
		hashes = append(hashes, h)
		writers = append(writers, h)
		names = append(names, name)
	}
	writers = append(writers, tmpFile)

	_, err = io.Copy(io.MultiWriter(writers...), r.Body)
	if err != nil {
		http.Error(w, fmt.Errorf("error copying request body: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	for i, h := range hashes {
		digest := hex.EncodeToString(h.Sum(nil))
		dest := filepath.Join(cas.storageDir, names[i], digest)
		w.Header().Set("X-Digest-"+names[i], digest)
		err := os.Link(tmpFile.Name(), dest)
		if err != nil {
			if !errors.Is(err, os.ErrExist) {
				http.Error(w, fmt.Errorf("error creating hard link: %w", err).Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	if err := errors.Join(tmpFile.Close(), os.Remove(tmpFile.Name())); err != nil {
		fmt.Print(fmt.Errorf("error closing and removing temporary file: %w", err))
	}

	// _, err = io.Copy(tmpFile, r.Body)
	// if err != nil {
	//     http.Error(w, fmt.Errorf("Error copying request body to temporary file: %w", err).Error(), http.StatusInternalServerError)
	//     return
	// }

	// for algorithmName, algorithm := range cas.algorithms {
	//     algorithmDir := filepath.Join(cas.storageDir, algorithmName)
	//     if err := atomicMkdir(algorithmDir); err != nil {
	//         http.Error(w, "Error creating algorithm directory", http.StatusInternalServerError)
	//         return
	//     }

	//     digest := calculateDigest(tmpFile.Name(), algorithm)
	//     if err != nil {
	//         http.Error(w, fmt.Sprintf("Error calculating digest: %s", err), http.StatusInternalServerError)
	//         return
	//     }

	//     linkPath := filepath.Join(algorithmDir, digest)

	//     if _, err := os.Stat(linkPath); os.IsNotExist(err) {
	//         size, err := atomicLink(tmpFile.Name(), linkPath)
	//         if err != nil {
	//             http.Error(w, fmt.Sprintf("Error creating file at %s: %s", linkPath, err), http.StatusInternalServerError)
	//             return
	//         }

	//         updateFileSize(linkPath, size)

	//         w.Header().Set("X-Digest-"+algorithmName, digest)
	//     } else {
	//         w.Header().Set("X-Digest-"+algorithmName, digest)
	//     }
	// }
	// w.WriteHeader(http.StatusOK)
}

func (cas *CAS) handleBlobRetrieval(w http.ResponseWriter, r *http.Request) {
	algorithm, digest, err := parseURLPath(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Errorf("error parsing URL: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println("Algorithm:", algorithm)
	// fmt.Println("Digest:", digest)

	algorithmDir := filepath.Join(cas.storageDir, algorithm)
	// alg, digest := parts[2], parts[3]
	filePath := filepath.Join(algorithmDir, digest)
	fmt.Println("File Path:", filePath)

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	http.Error(w, fmt.Errorf("File not found: %w", err).Error(), http.StatusNotFound)
	// 	return
	// }
	// defer file.Close()

	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)

	// _, err = io.Copy(w, file)
	// if err != nil {
	// 	http.Error(w, fmt.Errorf("Error copying file to response: %w", err).Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func parseURLPath(path string) (string, string, error) {
	parts := strings.Split(path, "/")

	if len(parts) != 4 || parts[0] != "" || parts[1] != "blob" {
		return "", "", errors.New("Expected format blob/DIGEST/HEX")
	}
	// if len(parts) < 4 {
	// 	return "", "", errors.New("invalid URL path")
	// }

	return parts[2], parts[3], nil
}

func (cas *CAS) handleStats(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cas.getStats(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (cas *CAS) getStats(w http.ResponseWriter, r *http.Request) {
	filePaths, err := listFiles(cas.storageDir)
	if err != nil {
		http.Error(w, fmt.Errorf("error listing files: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	count := len(filePaths) / 2
	mean, stddev := calculateStats(filePaths)

	stats := FileStats{
		Count:  count,
		Mean:   mean,
		Stddev: stddev,
	}

	statsJSON, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, fmt.Errorf("Error encoding statistics to JSON: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(statsJSON)
	if err != nil {
		http.Error(w, fmt.Errorf("Error writing statistics: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func calculateStats(filePaths []string) (float64, float64) {
	var totalSize int64
	for _, filePath := range filePaths {
		size, _ := getFileStats(filePath)
		totalSize += size
	}
	count := len(filePaths)
	mean := float64(totalSize) / float64(count)

	var sumSquaredDiff float64
	for _, filePath := range filePaths {
		size, _ := getFileStats(filePath)
		diff := float64(size) - mean
		sumSquaredDiff += diff * diff
	}

	stddev := math.Sqrt(sumSquaredDiff / float64(count))
	return mean, stddev
}

func getFileStats(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to get file stats: %w", err)
	}
	return fileInfo.Size(), nil
}

func listFiles(dir string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}
	return filePaths, nil
}
