package youtube

import (
	"io"

	"google.golang.org/api/youtube/v3"
	ytcore "google.golang.org/api/youtube/v3"
)

// UploadToYoutube uses service to upload file to YouTube and returns the new
// YouTube video
//
// Get the service from GetService
func UploadToYoutube(
	service *ytcore.Service,
	file io.Reader,
	channelID string,
	title string,
) (*ytcore.Video, error) {
	video := &ytcore.Video{
		Snippet: &youtube.VideoSnippet{
			CategoryId: "22",
			ChannelId:  channelID,
			Title:      title,
		},
		Status: &ytcore.VideoStatus{
			PrivacyStatus: "private",
		},
	}
	call := service.Videos.Insert("snippet,status", video)

	response, err := call.Media(file).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
