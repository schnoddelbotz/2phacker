package inputs

type AudioProcessingQueue struct {
	Extensions []string
}

func NewAudioProcessingQueue() *AudioProcessingQueue {
	return &AudioProcessingQueue{
		Extensions: []string{"mp3", "m4a"},
	}
}
