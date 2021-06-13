package main

import (
	"gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{
		"GOD BLESS THE RATCHETS",
		"a p p l y i n g  p r e s s u r e",
		"Handsome",
	}

	// After adding an interface and updating the playList() func,
	// can now pass both TapePlayer and TapeRecorder types to playList()
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
