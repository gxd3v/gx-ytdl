package message

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/protobuf/types/known/anypb"
)

/*
*pb.DeleteSessionResponse
 */

func NewDownloadResponse(_ context.Context, code int32, msg string, err error) *pb.DownloadResponse {
	var response pb.DownloadResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
	}

	return &response
}

func NewCreateSessionFolderResponse(_ context.Context, code int32, msg string, err error) *pb.CreateSessionFolderResponse {
	var response pb.CreateSessionFolderResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
		response.Created = false
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
		response.Created = true
	}

	return &response
}

func NewListFilesResponse(_ context.Context, code int32, msg string, err error, files []*pb.File) *pb.ListFilesResponse {
	var response pb.ListFilesResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
		response.Files = files
	}

	return &response
}

func NewSendFileToClientResponse(_ context.Context, code int32, msg string, err error, file *anypb.Any) *pb.SendFileToClientResponse {
	var response pb.SendFileToClientResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
		response.File = file
	}

	return &response
}

func NewDeleteFileResponse(_ context.Context, code int32, msg string, err error) *pb.DeleteFileResponse {
	var response pb.DeleteFileResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
	}

	return &response
}

func NewDeleteSessionResponse(_ context.Context, code int32, msg string, err error) *pb.DeleteSessionResponse {
	var response pb.DeleteSessionResponse

	response.Id = uuid.NewString()

	if err != nil {
		response.Successful = false
		response.Error = &pb.Error{
			Code:    pb.ErrorsEnum(code),
			Message: fmt.Sprintf("%s: %s", msg, err.Error()),
		}
	} else {
		response.Successful = true
		response.Success = &pb.Success{
			Code:   pb.SuccessEnum(code),
			Status: msg,
		}
	}

	return &response
}
