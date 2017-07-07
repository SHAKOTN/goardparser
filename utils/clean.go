package utils

import "os"

func Clean(folderPath string) {
	os.RemoveAll(folderPath)
	os.Remove(folderPath + ".zip")
}
