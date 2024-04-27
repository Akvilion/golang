package main

import "xxx/src/github.com/headfirstgo/gadget"

func PlayList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	// var player gadget.Player
	mixtape := []string{"Sting", "Yello", "Depech mode"}
	var player gadget.Player = gadget.TapePlayer{}
	PlayList(player, mixtape)
	player = gadget.TapeRecorder{}
	PlayList(player, mixtape)
}
