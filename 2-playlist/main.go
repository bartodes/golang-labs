package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var playlist_name string
	var song string
	var playlist []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Choose your playlist name: ")
	scanner.Scan()
	playlist_name = scanner.Text()

	fmt.Print("Choose your song: ")
	scanner.Scan()
	song = scanner.Text()

	valid_playlist := len(strings.TrimSpace(playlist_name)) > 0
	valid_song := len(strings.TrimSpace(song)) > 0

	if !valid_playlist {
		fmt.Println("Playlist name mustn't be nothing.")
		return
	}

	if !valid_song {
		fmt.Println("song name mustn't be nothing.")
		return
	}

	playlist = append(playlist, song)
	fmt.Printf("%#v has been added to the playlist '%v'\n", playlist, playlist_name)

}
