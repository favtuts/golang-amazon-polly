# golang-amazon-polly
Generating Text To Speech MP3 files with Amazon Polly: https://www.favtuts.com/golang-go-crash-course-10-generating-text-to-speech-mp3-files-with-amazon-polly/

# initialize the golang project

```
$ go mod init github.com/favtuts/golang-amazon-polly
```

# install AWS SDK for Golang

```
$ go get -u github.com/aws/aws-sdk-go
go: added github.com/aws/aws-sdk-go v1.44.235
go: added github.com/jmespath/go-jmespath v0.4.0
```

# run and test application

```
$ go run *.go
$ play kimberly.mp3
$ play yoey.mp3
```