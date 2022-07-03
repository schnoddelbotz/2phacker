package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"schnoddelbotz/mm-webpage/inputs"
)

type texasRanger struct {
	rootDirectory string
	audioQueue    inputs.AudioProcessingQueue
}

func NewWalker(rootDirectory string) *texasRanger {
	return &texasRanger{
		rootDirectory: rootDirectory,
		audioQueue:    *inputs.NewAudioProcessingQueue(),
	}
}

func (r *texasRanger) RunForest() {
	println("Yeah ...")
	println(r.rootDirectory)
	filepath.Walk(r.rootDirectory, r.myWalker)
}

func (r *texasRanger) myWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fileExtension := strings.ToLower(filepath.Ext(path))

	// todo generic(s), hicks?

	for _, ext := range r.audioQueue.Extensions {
		if ext == fileExtension {
			println(fileExtension)
			// add to audio queue
			return nil
		}
	}

	fmt.Println(path, info.Size())
	return nil
}
