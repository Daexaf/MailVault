package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func SaveAttachment(
	basePath string,
	mailAccountID uint,
	emailID uint,
	originalName string,
	data []byte,
) (storedName string, storagePath string, err error) {
	ext := filepath.Ext(originalName)

	storedName = fmt.Sprintf(
		"%d%s",
		time.Now().UnixNano(),
		ext,
	)

	dir := filepath.Join(
		basePath,
		"mail-accounts",
		fmt.Sprintf("%d", mailAccountID),
		"emails",
		fmt.Sprintf("%d", emailID),
	)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}

	fullPath := filepath.Join(dir, storedName)

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return "", "", err
	}

	return storedName, fullPath, nil
}
