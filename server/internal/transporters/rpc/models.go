package rpc

import (
	"context"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Downloader struct {
}

type downloader interface {
	Download(context.Context, *pb.DownloadRequest) (*pb.DownloadResponse, error)
	CreateSessionFolder(context.Context, *pb.CreateSessionFolderRequest) (*pb.CreateSessionFolderResponse, error)
	ListFiles(context.Context, *emptypb.Empty) (*pb.ListFilesResponse, error)
	SendFileToClient(context.Context, *pb.SendFileToClientRequest) (*pb.SendFileToClientResponse, error)
	DeleteFile(context.Context, *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error)
	DeleteSession(context.Context, *emptypb.Empty) (*pb.DeleteSessionResponse, error)
}
