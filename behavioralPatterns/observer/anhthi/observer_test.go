package anhthi

import "testing"

func TestObserver(t *testing.T) {
	var ytChannel Subject = &YoutubeChannel{}
	ui1 := NewUserInterface("Bob")
	ui2 := NewUserInterface("Peter")
	ytChannel.registerObserver(ui1)
	ytChannel.registerObserver(ui2)
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avatar 2 trailer"})
	ytChannel.(*YoutubeChannel).ReleaseNewVideo(&Video{title: "Avengers Endgame trailer"})
}
