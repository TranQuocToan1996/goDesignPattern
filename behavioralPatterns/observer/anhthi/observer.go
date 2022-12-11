package anhthi

import (
	"fmt"
	"reflect"
)

type Observer interface {
	update(interface{})
}

type Subject interface {
	registerObserver(obs Observer)
	removeObserver(obs Observer)
	notifyObservers()
}

type Video struct {
	title string
}

type View struct {
	video []Video
}

func (v *View) Append(vid Video) {
	v.video = append(v.video, vid)
}

// YoutubeChannel is a concrete implementation of Subject interface
type YoutubeChannel struct {
	Observers []Observer
	NewVideo  *Video
}

func (yt *YoutubeChannel) registerObserver(obs Observer) {
	yt.Observers = append(yt.Observers, obs)
}

func (yt *YoutubeChannel) removeObserver(obs Observer) {
	for i, o := range yt.Observers {
		if reflect.DeepEqual(o, obs) {
			yt.Observers = append(yt.Observers[:i], yt.Observers[i+1:]...)
		}
	}
	// TODO: Handle not match case
	// TODO: Noti to observer the result
}

// notify to all observers when a new video is released
func (yt *YoutubeChannel) notifyObservers() {
	for _, obs := range yt.Observers {
		obs.update(yt.NewVideo)
	}
}

func (yt *YoutubeChannel) ReleaseNewVideo(video *Video) {
	yt.NewVideo = video
	yt.notifyObservers()
}

// UserInterface is a concrete implementation of Observer interface
type UserInterface struct {
	UserName string
	Videos   []*Video
}

func (ui *UserInterface) update(video interface{}) {
	ui.Videos = append(ui.Videos, video.(*Video))
	for video := range ui.Videos {
		fmt.Println(video)
		//TODO: add view components
		// View.AddChildNode(NewVideoComponent(video))
	}
	fmt.Printf("UI %s - Video: '%s' has just been released\n", ui.UserName, video.(*Video).title)
}

func NewUserInterface(username string) Observer {
	return &UserInterface{UserName: username, Videos: make([]*Video, 0)}
}
