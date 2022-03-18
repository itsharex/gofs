package fs

import (
	"fmt"
	"github.com/no-src/log"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var deletedPathRegexp *regexp.Regexp

// LogicallyDelete delete the path logically
func LogicallyDelete(path string) error {
	if IsDeleted(path) {
		return nil
	}
	deletedFile := toDeletedPath(path)
	err := rename(path, deletedFile)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// IsDeleted is deleted path
func IsDeleted(path string) bool {
	return isDeleted(path)
}

func isDeletedCore(path string) bool {
	return deletedPathRegexp.MatchString(path)
}

// ClearDeletedFile remove all the deleted files in the path
func ClearDeletedFile(clearPath string) error {
	return filepath.WalkDir(clearPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil && isNotExist(err) {
			return nil
		}
		if err != nil {
			return err
		}
		if IsDeleted(path) {
			err = removeAll(path)
			if err != nil {
				log.Error(err, "remove the deleted files error => [%s]", path)
			} else {
				log.Debug("remove the deleted files success => [%s]", path)
			}
		}
		return err
	})
}

// toDeletedPath convert to the logically deleted file name
func toDeletedPath(path string) string {
	return fmt.Sprintf("%s.%d.deleted", path, time.Now().Unix())
}

var (
	removeAll = os.RemoveAll
	rename    = os.Rename
	isDeleted = isDeletedCore
)

func init() {
	deletedPathRegexp = regexp.MustCompile(`^[\s\S]+\.[0-9]{10,}\.(?i)deleted$`)
}
