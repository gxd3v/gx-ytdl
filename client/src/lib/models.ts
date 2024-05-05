import * as pb from '$lib/protos/messages';

export type Message =
  | pb.DownloadResponse
  | pb.CreateSessionFolderResponse
  | pb.SendFileToClientResponse
  | pb.DeleteFileResponse
  | pb.CreateSessionResponse
  | pb.DeleteSessionResponse;
