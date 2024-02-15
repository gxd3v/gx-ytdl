package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	c "github.com/gx/youtubeDownloader/constants"
	pb "github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"os/exec"
	"path"
	"time"
)

var _ Business = (*Server)(nil)

func (s *Server) Download(_ *gin.Context, request *pb.DownloadRequest) *pb.DownloadResponse {
	transaction := s.Database.Transactional()
	defer func() { _ = transaction.Commit() }()

	s.Logger.Info("Trying to parse the URL")
	url, err := util.ParseURL(request.Payload.GetUrl())
	if err != nil {
		s.Logger.Error("URL was not valid", err.Error())
		return &pb.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FAILED_DOWNLOAD,
				Message: "Invalid URL: " + err.Error(),
			},
		}
	}

	outputPath := fmt.Sprintf(s.Config.OutputPath, s.SessionID)

	s.Logger.Info("Invoking external downloader")
	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", url.String())
	cmd.Args = append(cmd.Args, "-op", outputPath)
	if request.Payload.GetAudio() {
		cmd.Args = append(cmd.Args, "-a")
	}

	err = cmd.Run()
	if err != nil {
		s.Logger.Error("Failed to run external downloader", err.Error())
		return &pb.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FAILED_DOWNLOAD,
				Message: err.Error(),
			},
		}
	}

	_ = cmd.Wait()

	s.Logger.Info("Download successfully finished")

	//files := s.ListFiles(ctx)
	ytdl := s.Database.NewYTDL(url.String(), s.Storage, s.SessionID, 0)

	err = transaction.Insert(ytdl)
	if err != nil {
		_ = transaction.Rollback()
	}

	return &pb.DownloadResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &pb.Success{
			Code:   pb.SuccessEnum_VIDEO_DOWNLOADABLE,
			Status: "File is done downloading",
		},
	}
}

func (s *Server) CreateSessionFolder(_ *gin.Context, request *pb.CreateSessionFolderRequest) *pb.CreateSessionFolderResponse {
	s.Logger.Info("Creating folder to store downloads")
	s.Storage = fmt.Sprintf(s.Config.OutputPath, request.Payload.GetSession())
	err := os.Mkdir(s.Storage, os.ModeAppend)
	if err != nil {
		s.Logger.Error("Failed to create a session folder")
		return &pb.CreateSessionFolderResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FOLDER_ALREADY_EXISTS,
				Message: err.Error(),
			},
			Created: false,
		}
	}

	return &pb.CreateSessionFolderResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &pb.Success{
			Code:   pb.SuccessEnum_SESSION_FOLDER_CREATED,
			Status: "Session folder created",
		},
		Created: true,
	}
}

func (s *Server) ListFiles(ctx *gin.Context) *pb.ListFilesResponse {
	if files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID)); err != nil {
		s.Logger.Error("Failed to read the directory", err.Error())
		return &pb.ListFilesResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FAILED_LISTING_FILES,
				Message: "Error listing files: " + err.Error(),
			},
		}
	} else {
		if len(files) == 0 {
			s.Logger.Warning("No files in folder to show")
			s.SendMessage(ctx, &pb.ListFilesResponse{
				Id:         uuid.NewString(),
				Successful: false,
				Error: &pb.Error{
					Code:    pb.ErrorsEnum_NO_ITEMS_PRESENT,
					Message: err.Error(),
				},
			})
		} else {
			s.Logger.Info("Sending list of files to client")
			returningFiles := make([]*pb.File, 0)

			for _, file := range files {
				info, _ := file.Info()
				returningFiles = append(returningFiles, &pb.File{
					Name:            file.Name(),
					Created:         timestamppb.New(info.ModTime()),
					TimesDownloaded: 0,
					Ttl:             timestamppb.New(info.ModTime().Add(time.Second * c.FileTtl)),
					Size:            info.Size(),
				})
			}

			return &pb.ListFilesResponse{
				Id:         uuid.NewString(),
				Successful: true,
				Success: &pb.Success{
					Code:   pb.SuccessEnum_LISTED_FILES,
					Status: "Files listed with success",
				},
				Files: returningFiles,
			}

		}
	}

	return &pb.ListFilesResponse{
		Id:         uuid.NewString(),
		Successful: false,
		Error: &pb.Error{
			Code:    pb.ErrorsEnum_CATASTROPHIC_ERROR,
			Message: "Something went terribly wrong",
		},
	}
}

func (s *Server) SendFileToClient(_ *gin.Context, request *pb.SendFileToClientRequest) *pb.SendFileToClientResponse {
	s.Logger.Info("Client requested file", request.Payload.File.Name)
	return &pb.SendFileToClientResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &pb.Success{
			Code:   pb.SuccessEnum_READY_TO_SEND,
			Status: "Retrieved file: " + request.Payload.GetFile().GetName(),
		},
		File: nil,
	}
}

func (s *Server) DeleteFile(ctx *gin.Context, request *pb.DeleteFileRequest) *pb.DeleteFileResponse {
	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), request.Payload.File.Name))
	if err != nil {
		s.Logger.Error("Failed to delete a file", err.Error())
		return &pb.DeleteFileResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FAILED_DELETE_FILE,
				Message: err.Error(),
			},
		}
	} else {
		s.Logger.Info("File was deleted successfully")
		files := s.ListFiles(ctx)
		if !files.Successful {
			s.Logger.Error("Couldn't list files")
			return &pb.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: false,
				Error: &pb.Error{
					Code:    pb.ErrorsEnum_FAILED_LISTING_FILES,
					Message: "Couldn't list files after deletion",
				},
			}
		} else {
			s.Logger.Info("Sending new list of files to client")
			return &pb.DeleteFileResponse{
				Id:         uuid.NewString(),
				Successful: true,
				Success: &pb.Success{
					Code:   pb.SuccessEnum_DELETED_FILE,
					Status: "File was deleted successfully",
				},
				Files: files.Files,
			}
		}
	}
}

func (s *Server) DeleteSession(_ *gin.Context) *pb.DeleteSessionResponse {
	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		s.Logger.Error("Failed to delete session")
		return &pb.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &pb.Error{
				Code:    pb.ErrorsEnum_FAILED_DELETE_SESSION,
				Message: err.Error(),
			},
		}
	} else {
		err = s.Ws.Close()
		s.SessionID = ""

		s.Logger.Info("Session deleted, disconnecting client from the server")
		return &pb.DeleteSessionResponse{
			Id:         uuid.NewString(),
			Successful: true,
			Success: &pb.Success{
				Code:   pb.SuccessEnum_SESSION_DELETE,
				Status: "Your session was deleted, forever.",
			},
		}
	}
}
