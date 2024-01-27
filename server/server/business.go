package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gx/youtubeDownloader/protos"
	"os"
	"os/exec"
)

var _ Business = (*Server)(nil)

func (s *Server) Download(_ *gin.Context, request *protos.DownloadRequest) (*protos.DownloadResponse, error) {
	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
	cmd.Args = append(cmd.Args, "-u", request.Payload.GetUrl())
	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(s.Config.OutputPath, s.SessionID))
	if request.Payload.GetAudio() {
		cmd.Args = append(cmd.Args, "-a")
	}

	err := cmd.Run()
	if err != nil {
		return &protos.DownloadResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:  protos.ErrorsEnum_FAILED_DOWNLOAD,
				Error: err.Error(),
			},
		}, nil
	}

	_ = cmd.Wait()
	return &protos.DownloadResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code: protos.SuccessEnum_VIDEO_DOWNLOADABLE,
		},
		Data: "File is done downloading",
	}, nil

}

func (s *Server) CreateSessionFolder(_ *gin.Context, request *protos.CreateSessionFolderRequest) (*protos.CreateSessionFolderResponse, error) {
	s.Logger.Info(fmt.Sprintf("Creating folder %s to store downloads", request.Payload.GetSession()))
	s.Storage = fmt.Sprintf(s.Config.OutputPath, request.Payload.GetSession())
	err := os.Mkdir(s.Storage, os.ModeAppend)
	if err != nil {
		return &protos.CreateSessionFolderResponse{
			Id:         uuid.NewString(),
			Successful: false,
			Error: &protos.Error{
				Code:  protos.ErrorsEnum_FOLDER_ALREADY_EXISTS,
				Error: err.Error(),
			},
			Created: false,
		}, nil
	}

	return &protos.CreateSessionFolderResponse{
		Id:         uuid.NewString(),
		Successful: true,
		Success: &protos.Success{
			Code: protos.SuccessEnum_SESSION_FOLDER_CREATED,
		},
		Created: true,
	}, nil
}

func (s *Server) ListFiles(ctx *gin.Context, request *protos.ListFilesRequest) (*protos.ListFilesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) SendFileToClient(ctx *gin.Context, request *protos.SendFileToClientRequest) (*protos.SendFileToClientResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteFile(ctx *gin.Context, request *protos.DeleteFileRequest) (*protos.DeleteFileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteSession(ctx *gin.Context, request *protos.DeleteSessionRequest) (*protos.DeleteSessionResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (s *Server) Download(ctx *gin.Context, audio bool, url string) {
//	cmd := exec.Command(s.Config.PythonBinary, s.Config.DownloaderPath)
//	cmd.Args = append(cmd.Args, "-u", url)
//	cmd.Args = append(cmd.Args, "-op", fmt.Sprintf(s.Config.OutputPath, s.SessionID))
//	if audio {
//		cmd.Args = append(cmd.Args, "-a")
//	}
//
//	err := cmd.Run()
//	if err != nil {
//		s.SendMessage(ctx, c.CODE_ERROR_DOWNLOAD_FAILED, err.Error())
//		return
//	}
//
//	go func() {
//		_ = cmd.Wait()
//		s.SendMessage(ctx, c.CODE_SUCCESS_VIDEO_DOWNLOADABLE, "File is done downloading")
//	}()
//}
//
//func (s *Server) CreateSessionFolder() {
//	s.Logger.Info(fmt.Sprintf("Creating folder %s to store downloads", s.SessionID))
//	s.Storage = fmt.Sprintf(s.Config.OutputPath, s.SessionID)
//	_ = os.Mkdir(s.Storage, os.ModeAppend)
//}
//
//func (s *Server) ListFiles(ctx *gin.Context) {
//	if files, err := os.ReadDir(fmt.Sprintf(s.Config.OutputPath, s.SessionID)); err != nil {
//		s.Logger.Error(c.TEXT_ERROR_FAILED_LISTING_FILES, err.Error())
//		s.SendMessage(ctx, c.CODE_ERROR_FAILED_LISTING_FILES, c.TEXT_ERROR_FAILED_LISTING_FILES)
//	} else {
//		if len(files) == 0 {
//			s.Logger.Warning("No files in folder to show")
//			s.SendMessage(ctx, c.CODE_SUCCESS_LISTED_FILES, base64.StdEncoding.EncodeToString([]byte("{}")))
//		} else {
//			output := map[int]string{}
//
//			for index, file := range files {
//				output[index] = file.Name()
//			}
//
//			out, _ := json.Marshal(output)
//			s.Logger.Info("files in the session", base64.StdEncoding.EncodeToString(out))
//			s.SendMessage(ctx, c.CODE_SUCCESS_LISTED_FILES, base64.StdEncoding.EncodeToString(out))
//		}
//
//	}
//}
//
//func (s *Server) SendFileToClient(ctx *gin.Context, fileName string) {
//	s.Logger.Info("Client requested file", fileName)
//	s.SendMessage(ctx, c.CODE_SUCCESS_READY_TO_SEND, fileName)
//}
//
//func (s *Server) DeleteFile(ctx *gin.Context, fileName string) {
//	err := os.Remove(path.Join(fmt.Sprintf(s.Config.OutputPath, s.SessionID), fileName))
//	if err != nil {
//		s.Logger.Error(c.TEXT_ERROR_FAILED_DELETE_FILE)
//		s.SendMessage(ctx, c.CODE_ERROR_FAILED_DELETE_FILE, c.TEXT_ERROR_FAILED_DELETE_FILE)
//	} else {
//		s.Logger.Info("Client removed file", fileName)
//		s.SendMessage(ctx, c.CODE_SUCCESS_DELETE_FILE, "File deleted")
//		s.ListFiles(ctx)
//	}
//}
//
//func (s *Server) DeleteSession(ctx *gin.Context) {
//	err := os.Remove(fmt.Sprintf(s.Config.OutputPath, s.SessionID))
//	if err != nil {
//		s.Logger.Error(c.TEXT_ERROR_FAILED_DELETE_SESSION)
//		s.SendMessage(ctx, c.CODE_SUCCESS_SESSION_DELETE, c.TEXT_ERROR_FAILED_DELETE_SESSION)
//	} else {
//		s.Logger.Info("Client removed session")
//		s.SendMessage(ctx, c.CODE_SUCCESS_SESSION_DELETE, "Session deleted")
//
//		_ = s.Ws.Close()
//		s.SessionID = ""
//	}
//}
//
