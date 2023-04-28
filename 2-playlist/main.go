package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		playlist_name string
		song          string
		playlist      []string
	)

	fmt.Print("Type 'exit' when you are done!\n\n")
	fmt.Print("Choose your playlist name: ")
	readName(&playlist_name)

	exit := playlist_name == "exit"

	for !exit {
		fmt.Printf("Choose a song for '%s' playlist: ", playlist_name)
		readName(&song)

		if song == "exit" {
			break
		}

		valid_playlist := len(strings.TrimSpace(playlist_name)) > 0
		valid_song := len(strings.TrimSpace(song)) > 0

		if !valid_playlist {
			fmt.Println("Playlist name mustn't be nothing.")
			return
		}

		if !valid_song {
			fmt.Println("Song name mustn't be nothing.")
			return
		}

		playlist = append(playlist, song)
	}

	fmt.Printf("Here are all the songs in '%s': %q\n", playlist_name, playlist)
	fmt.Println("Enjoy! Bye")
}

func readName(name *string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*name = scanner.Text()
}
