package main

import "github.com/favtuts/golang-amazon-polly/service"

var (
	kimberly service.PollyService = service.NewKimberlyPollyService()
	joey     service.PollyService = service.NewJoeyPollyService()
)

func main() {
	err := kimberly.Synthesize("Hi, I am Kimberly, how are you?", "kimberly.mp3")
	if err != nil {
		panic(err)
	}

	err = joey.Synthesize("Hi, I am Joey. Nice to meet you.", "yoey.mp3")
	if err != nil {
		panic(err)
	}
}
