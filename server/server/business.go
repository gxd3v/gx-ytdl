package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"os/exec"
	"path"
	"time"
)

var _ Business = (*Server)(nil)

func (s *Server) Download(_ *gin.Context, request *protos.DownloadRequest) (*protos.DownloadResponse, error) {
	url, err := util.ParseURL(request.Payload.GetUrl())
	if err != nil {
		return &protos.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DOWNLOAD,
				Message: "Invalid URL: " + err.Error(),
			},
		}, nil
	}

	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", url.String())
	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if request.Payload.GetAudio() {
		cmd.Args = append(cmd.Args, "-a")
	}

	err = cmd.Run()
	if err != nil {
		return &protos.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DOWNLOAD,
				Message: err.Error(),
			},
		}, nil
	}

	_ = cmd.Wait()
	return &protos.DownloadResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_VIDEO_DOWNLOADABLE,
			Status: "File is done downloading",
		},
	}, nil

}

func (s *Server) CreateSessionFolder(_ *gin.Context, request *protos.CreateSessionFolderRequest) (*protos.CreateSessionFolderResponse, error) {
	s.Logger.Info("Creating folder to store downloads")
	s.Storage = fmt.Sprintf(s.Config.OutputPath, request.Payload.GetSession())
	err := os.Mkdir(s.Storage, os.ModeAppend)
	if err != nil {
		return &protos.CreateSessionFolderResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FOLDER_ALREADY_EXISTS,
				Message: err.Error(),
			},
			Created: false,
		}, nil
	}

	return &protos.CreateSessionFolderResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_SESSION_FOLDER_CREATED,
			Status: "Session folder created",
		},
		Created: true,
	}, nil
}

func (s *Server) ListFiles(ctx *gin.Context) (*protos.ListFilesResponse, error) {
	if files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID)); err != nil {
		s.Logger.Error(c.TEXT_ERROR_FAILED_LISTING_FILES, err.Error())
		return &protos.ListFilesResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_LISTING_FILES,
				Message: "Error listing files: " + err.Error(),
			},
		}, nil
	} else {
		if len(files) == 0 {
			s.Logger.Warning("No files in folder to show")
			s.SendMessage(ctx, &protos.ListFilesResponse{
				Id:         uuid.NewString(),
				Successful: false,
				Error: &protos.Error{
					Code:    protos.ErrorsEnum_NO_ITEMS_PRESENT,
					Message: err.Error(),
				},
			})
		} else {
			returningFiles := make([]*protos.File, 0)

			for _, file := range files {
				info, _ := file.Info()
				returningFiles = append(returningFiles, &protos.File{
					Name:            file.Name(),
					Created:         timestamppb.New(info.ModTime()),
					TimesDownloaded: 0,
					Ttl:             timestamppb.New(info.ModTime().Add(time.Second * c.FILE_TTL)),
					Size:            info.Size(),
				})
			}

			return &protos.ListFilesResponse{
				Id:         uuid.NewString(),
				Successful: true,
				Success: &protos.Success{
					Code:   protos.SuccessEnum_LISTED_FILES,
					Status: "Files listed with success",
				},
				Files: returningFiles,
			}, nil

		}
	}

	return &protos.ListFilesResponse{
		Id:         uuid.NewString(),
		Successful: false,
		Error: &protos.Error{
			Code:    protos.ErrorsEnum_CATASTROPHIC_ERROR,
			Message: "Something went terribly wrong",
		},
	}, nil
}

func (s *Server) SendFileToClient(_ *gin.Context, request *protos.SendFileToClientRequest) (*protos.SendFileToClientResponse, error) {
	s.Logger.Info("Client requested file", request.Payload.File.Name)
	return &protos.SendFileToClientResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_READY_TO_SEND,
			Status: "Retrieved file: " + request.Payload.GetFile().GetName(),
		},
		File: nil,
	}, nil
}

func (s *Server) DeleteFile(ctx *gin.Context, request *protos.DeleteFileRequest) (*protos.DeleteFileResponse, error) {
	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), request.Payload.File.Name))
	if err != nil {
		return &protos.DeleteFileResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DELETE_FILE,
				Message: err.Error(),
			},
		}, nil
	} else {
		files, err := s.ListFiles(ctx)
		if err != nil {
			return &protos.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: false,
				Error: &protos.Error{
					Code:    protos.ErrorsEnum_FAILED_LISTING_FILES,
					Message: "Couldn't list files after deletion",
				},
			}, nil
		} else {
			return &protos.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: true,
				Success: &protos.Success{
					Code:   protos.SuccessEnum_DELETED_FILE,
					Status: "File was deleted successfully",
				},
				Files: files.Files,
			}, nil
		}
	}
}

func (s *Server) DeleteSession(_ *gin.Context) (*protos.DeleteSessionResponse, error) {
	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		return &protos.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DELETE_SESSION,
				Message: err.Error(),
			},
		}, nil
	} else {
		err = s.Ws.Close()
		s.SessionID = ""

		return &protos.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: true,
			Success: &protos.Success{
				Code:   protos.SuccessEnum_SESSION_DELETE,
				Status: "Your session was deleted, forever.",
			},
		}, nil
	}
}
