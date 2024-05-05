package ws

import (
	"errors"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	"github.com/gx/youtubeDownloader/internal/logger"
	pb "github.com/gx/youtubeDownloader/protos"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func sendMessage(ctx buffalo.Context, message proto.Message) {
	ws, ok := ctx.Value("ws").(*websocket.Conn)
	if !ok {
		logger.Err(errors.New("no websocket in context")).Msg("context had no websocket value defined")
	}

	marshaller := protojson.MarshalOptions{
		EmitDefaultValues: true,
	}

	out, err := marshaller.Marshal(message)
	if err != nil {
		logger.Err(err).Msg("failed to marshall message")
	}

	if err = ws.WriteMessage(websocket.TextMessage, out); err != nil {
		logger.Err(err).Msg("failed to send message to client")
	}
}

func checkMessageError(ctx buffalo.Context, err error) bool {
	if err != nil {
		logger.Err(err).Msg("message was malformed")
		sendMessage(ctx, &pb.PanicResponse{
			Code:    pb.ErrorsEnum_MALFORMED_MESSAGE,
			Message: "message was malformed",
		})

		return false
	}

	return true
}
