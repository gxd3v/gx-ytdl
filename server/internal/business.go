package internal

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

func (s *Server) Download(_ *gin.Context, request *protos.DownloadRequest) *protos.DownloadResponse {
	s.Logger.Info("Trying to parse the URL")
	url, err := util.ParseURL(request.Payload.GetUrl())
	if err != nil {
		s.Logger.Error("URL was not valid", err.Error())
		return &protos.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DOWNLOAD,
				Message: "Invalid URL: " + err.Error(),
			},
		}
	}

	s.Logger.Info("Invoking external downloader")
	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", url.String())
	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if request.Payload.GetAudio() {
		cmd.Args = append(cmd.Args, "-a")
	}

	err = cmd.Run()
	if err != nil {
		s.Logger.Error("Failed to run external downloader", err.Error())
		return &protos.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DOWNLOAD,
				Message: err.Error(),
			},
		}
	}

	_ = cmd.Wait()
	s.Logger.Info("Download successfully finished")
	return &protos.DownloadResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_VIDEO_DOWNLOADABLE,
			Status: "File is done downloading",
		},
	}

}

func (s *Server) CreateSessionFolder(_ *gin.Context, request *protos.CreateSessionFolderRequest) *protos.CreateSessionFolderResponse {
	s.Logger.Info("Creating folder to store downloads")
	s.Storage = fmt.Sprintf(s.Config.OutputPath, request.Payload.GetSession())
	err := os.Mkdir(s.Storage, os.ModeAppend)
	if err != nil {
		s.Logger.Error("Failed to create a session folder")
		return &protos.CreateSessionFolderResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FOLDER_ALREADY_EXISTS,
				Message: err.Error(),
			},
			Created: false,
		}
	}

	return &protos.CreateSessionFolderResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_SESSION_FOLDER_CREATED,
			Status: "Session folder created",
		},
		Created: true,
	}
}

func (s *Server) ListFiles(ctx *gin.Context) *protos.ListFilesResponse {
	if files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID)); err != nil {
		s.Logger.Error("Failed to read the directory", err.Error())
		return &protos.ListFilesResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_LISTING_FILES,
				Message: "Error listing files: " + err.Error(),
			},
		}
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
			s.Logger.Info("Sending list of files to client")
			returningFiles := make([]*protos.File, 0)

			for _, file := range files {
				info, _ := file.Info()
				returningFiles = append(returningFiles, &protos.File{
					Name:            file.Name(),
					Created:         timestamppb.New(info.ModTime()),
					TimesDownloaded: 0,
					Ttl:             timestamppb.New(info.ModTime().Add(time.Second * c.FileTtl)),
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
			}

		}
	}

	return &protos.ListFilesResponse{
		Id:         uuid.NewString(),
		Successful: false,
		Error: &protos.Error{
			Code:    protos.ErrorsEnum_CATASTROPHIC_ERROR,
			Message: "Something went terribly wrong",
		},
	}
}

func (s *Server) SendFileToClient(_ *gin.Context, request *protos.SendFileToClientRequest) *protos.SendFileToClientResponse {
	s.Logger.Info("Client requested file", request.Payload.File.Name)
	return &protos.SendFileToClientResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code:   protos.SuccessEnum_READY_TO_SEND,
			Status: "Retrieved file: " + request.Payload.GetFile().GetName(),
		},
		File: nil,
	}
}

func (s *Server) DeleteFile(ctx *gin.Context, request *protos.DeleteFileRequest) *protos.DeleteFileResponse {
	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), request.Payload.File.Name))
	if err != nil {
		s.Logger.Error("Failed to delete a file", err.Error())
		return &protos.DeleteFileResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DELETE_FILE,
				Message: err.Error(),
			},
		}
	} else {
		s.Logger.Info("File was deleted successfully")
		files := s.ListFiles(ctx)
		if !files.Successful {
			s.Logger.Error("Couldn't list files")
			return &protos.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: false,
				Error: &protos.Error{
					Code:    protos.ErrorsEnum_FAILED_LISTING_FILES,
					Message: "Couldn't list files after deletion",
				},
			}
		} else {
			s.Logger.Info("Sending new list of files to client")
			return &protos.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: true,
				Success: &protos.Success{
					Code:   protos.SuccessEnum_DELETED_FILE,
					Status: "File was deleted successfully",
				},
				Files: files.Files,
			}
		}
	}
}

func (s *Server) DeleteSession(_ *gin.Context) *protos.DeleteSessionResponse {
	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		s.Logger.Error("Failed to delete session")
		return &protos.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:    protos.ErrorsEnum_FAILED_DELETE_SESSION,
				Message: err.Error(),
			},
		}
	} else {
		err = s.Ws.Close()
		s.SessionID = ""

		s.Logger.Info("Session deleted, disconnecting client from the internal")
		return &protos.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: true,
			Success: &protos.Success{
				Code:   protos.SuccessEnum_SESSION_DELETE,
				Status: "Your session was deleted, forever.",
			},
		}
	}
}
