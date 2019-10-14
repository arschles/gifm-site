package youtube

import (
	"io"

	"google.golang.org/api/youtube/v3"
)

// UploadToYoutube uses service to upload file to YouTube
//
// Get the service from GetService
func UploadToYoutube(service *youtube.Service, file io.Reader, channelID string, title string) (string, error) {
	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			CategoryId: "22",
			ChannelId:  channelID,z
			Title:      title,
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "private",
		},
	}
	call := service.Videos.Insert("snippet,status", video)

	response, err := call.Media(file).Do()
	if err != nil {
		return "", err
	}

	return response.Id, nil
}
