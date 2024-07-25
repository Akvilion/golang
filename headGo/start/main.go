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
	var player gadget.Player = gadget.TapePlayer{} // отут робимо змінну player типом gadget.Player, тоді в неї можна буде перезберігати інші обєкти інтерфейсу
	PlayList(player, mixtape)
	player = gadget.TapeRecorder{} // якщо вище player з типом інтерфейс, то сюди можна переасайнити інший обєкт
	PlayList(player, mixtape)
}
