package server

import (
	"context"
	"errors"
	"fmt"
	c "github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/log"
	pb "github.com/gx/youtubeDownloader/protos"
	"github.com/gx/youtubeDownloader/server/message"
	"github.com/gx/youtubeDownloader/util"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"os/exec"
	"path"
	"time"
)

func (s *Server) Download(ctx context.Context, request *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	database := s.Database.Transactional()
	defer database.Commit()

	log.Info("Trying to parse the URL")
	url, err := util.ParseURL(request.Payload.GetUrl())
	if err != nil {
		database.Rollback()
		log.Error(err, "URL was not valid")
		return message.NewDownloadResponse(ctx, int32(pb.ErrorsEnum_FAILED_DOWNLOAD), "Invalid URL", err), nil
	}

	outputPath := fmt.Sprintf(s.Config.OutputPath, s.SessionID)

	log.Info("Invoking external downloader")
	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", url.String())
	cmd.Args = append(cmd.Args, "-op", outputPath)
	if request.Payload.GetAudio() {
		cmd.Args = append(cmd.Args, "-a")
	}

	err = cmd.Run()
	if err != nil {
		database.Rollback()
		log.Error(err, "Failed to run external downloader")
		return message.NewDownloadResponse(ctx, int32(pb.ErrorsEnum_FAILED_DOWNLOAD), "", err), nil
	}

	_ = cmd.Wait()

	log.Info("Download successfully finished")

	//FIXME should not be sending messages here
	files, err := s.ListFiles(ctx, &emptypb.Empty{})
	s.sendMessage(ctx, files)

	ytdl := s.Database.NewYTDL(url.String(), s.Storage, s.SessionID, 0)

	_, err = database.Insert(ytdl)
	if err != nil {
		database.Rollback()
		log.Error(err, "Failed to insert data")
		return message.NewDownloadResponse(ctx, int32(pb.ErrorsEnum_FAILED_DOWNLOAD), "Download failed", errors.New("")), nil
	}

	return message.NewDownloadResponse(ctx, int32(pb.SuccessEnum_VIDEO_DOWNLOADABLE), "File is done downloading", nil), nil
}

func (s *Server) CreateSessionFolder(ctx context.Context, request *pb.CreateSessionFolderRequest) (*pb.CreateSessionFolderResponse, error) {
	log.Info("Creating folder to store downloads")
	s.Storage = fmt.Sprintf(s.Config.OutputPath, request.Payload.GetSession())
	err := os.Mkdir(s.Storage, os.ModeAppend)
	if err != nil {
		log.Error(err, "Failed to create a session folder")
		return message.NewCreateSessionFolderResponse(ctx, int32(pb.ErrorsEnum_FOLDER_ALREADY_EXISTS), "Failed to create the storage", err), nil
	}

	return message.NewCreateSessionFolderResponse(ctx, int32(pb.SuccessEnum_SESSION_FOLDER_CREATED), "Session folder created", nil), nil
}

func (s *Server) ListFiles(ctx context.Context, _ *emptypb.Empty) (*pb.ListFilesResponse, error) {
	files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		log.Error(err, "Failed to read the directory")
		return message.NewListFilesResponse(ctx, int32(pb.ErrorsEnum_FAILED_LISTING_FILES), "Error listing files", err, nil), nil
	}

	log.Info("Sending list of files to client")
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

	return message.NewListFilesResponse(ctx, int32(pb.SuccessEnum_LISTED_FILES), "Available files", nil, returningFiles), nil
}

func (s *Server) SendFileToClient(ctx context.Context, request *pb.SendFileToClientRequest) (*pb.SendFileToClientResponse, error) {
	log.Info("Client requested file", request.Payload.File.Name)
	return message.NewSendFileToClientResponse(ctx, int32(pb.SuccessEnum_READY_TO_SEND), "Retrieved file", nil, nil), nil
}

func (s *Server) DeleteFile(ctx context.Context, request *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), request.Payload.File.Name))
	if err != nil {
		log.Error(err, "Failed to delete a file")
		return message.NewDeleteFileResponse(ctx, int32(pb.ErrorsEnum_FAILED_DELETE_FILE), "Failed to delete file", err), nil
	}

	log.Info("File was deleted successfully")

	files, _ := s.ListFiles(ctx, &emptypb.Empty{})
	if !files.Successful {
		log.Error(err, "Couldn't list files")
		return message.NewDeleteFileResponse(ctx, int32(pb.ErrorsEnum_FAILED_LISTING_FILES), "Couldn't list files", nil), nil
	}

	//FIXME should not be sending messages here
	s.sendMessage(ctx, files)

	return message.NewDeleteFileResponse(ctx, int32(pb.SuccessEnum_DELETED_FILE), "File deleted", nil), nil
}

func (s *Server) DeleteSession(ctx context.Context, _ *emptypb.Empty) (*pb.DeleteSessionResponse, error) {
	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if err != nil {
		log.Error(err, "Failed to delete session folder")
		return message.NewDeleteSessionResponse(ctx, int32(pb.ErrorsEnum_FAILED_DELETE_SESSION), "Failed to close the session", err), nil
	}

	err = s.Ws.Close()
	if err != nil {
		log.Error(err, "Failed to close websocket connection")
		return message.NewDeleteSessionResponse(ctx, int32(pb.ErrorsEnum_FAILED_DELETE_SESSION), "Failed to close the session", err), nil
	}

	log.Info("Session deleted, disconnecting client from the server")
	s.SessionID = ""

	return message.NewDeleteSessionResponse(ctx, int32(pb.SuccessEnum_SESSION_DELETE), "Your session was deleted, forever.", nil), nil
}
