package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gx/youtubeDownloader/internal/logger"
	pb "github.com/gx/youtubeDownloader/protos"
)

func checkAction(ctx buffalo.Context, code int32, message []byte) error {
	switch code {
	case int32(pb.ActionsEnum_DOWNLOAD_AUDIO):
		logger.Info().Msg("starting audio download")
		request := &pb.DownloadRequest{}

		err := json.Unmarshal(message, &request)
		if ok := checkMessageError(ctx, err); ok {
			//download, _ := Download(ctx, request)
			//sendMessage(ctx, download)
		}

	case int32(pb.ActionsEnum_LIST_FILES):
		logger.Info().Msg("listing files")
		//files, _ := ListFiles(ctx, &emptypb.Empty{})
		//sendMessage(ctx, files)

	case int32(pb.ActionsEnum_DELETE_FILE):
		logger.Info().Msg("deleting a file")
		request := &pb.DeleteFileRequest{}

		err := json.Unmarshal(message, &request)
		if ok := checkMessageError(ctx, err); ok {
			//files, _ := DeleteFile(ctx, request)
			//sendMessage(ctx, files)
		}

	case int32(pb.ActionsEnum_RETRIEVE_FILE):
		logger.Info().Msg("sending a file to a client")
		request := &pb.SendFileToClientRequest{}

		err := json.Unmarshal(message, &request)
		if ok := checkMessageError(ctx, err); ok {
			//file, _ := SendFileToClient(ctx, request)
			//sendMessage(ctx, file)
		}

	default:
		logger.Info().Msg("message didn't have a known code")
		sendMessage(ctx, &pb.PanicResponse{
			Code:    pb.ErrorsEnum_NOT_RECOGNIZED,
			Message: fmt.Sprintf("The code %v sent was not recognized", code),
		})
	}

	return nil
}
