package main

import "xxx/src/github.com/headfirstgo/gadget"

func PlayList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Sting", "Yello", "Depech mode"}
	PlayList(player, mixtape)
}
