package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/s1ovac/todobot/pkg/storage"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultPerm = 0774
)

var ErrorNoSavedFiles = fmt.Errorf("no files in derectory")

type FileStorage struct {
	basePath string
}

func NewFileStorage(basePath string) *FileStorage {
	return &FileStorage{
		basePath: basePath,
	}
}

func (s *FileStorage) Save(p *storage.Page) error {
	fPath := filepath.Join(s.basePath, p.UserName)
	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}
	fName, err := fileName(p)
	if err != nil {
		return err
	}
	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(p); err != nil {
		return err
	}

	return nil
}
func (s *FileStorage) PickRandom(userName string) (*storage.Page, error) {
	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, ErrorNoSavedFiles
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(files))
	file := files[randIndex]

}
func (s *FileStorage) Remove(p *storage.Page) error {
	fileName, err := fileName(p)
	if err != nil {
		return err
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)

	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
func (s *FileStorage) IsExists(p *storage.Page) (bool, error) {
	fileName, err := fileName(p)
	if err != nil {
		return false, err
	}
	path := filepath.Join(s.basePath, p.UserName, fileName)
	switch _, err := os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, err
	}

	return true, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}

func (s *FileStorage) decodePath(filepath string) (*storage.Page, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()
	var p storage.Page

	if err := gob.NewEncoder(file).Encode(p); err != nil {
		return nil, err
	}
	return &p, nil
}
