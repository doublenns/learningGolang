package main

import "gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{
		"GOD BLESS THE RATCHETS",
		"a p p l y i n g  p r e s s u r e",
		"Handsome",
	}
	playList(player, mixtape)
}
