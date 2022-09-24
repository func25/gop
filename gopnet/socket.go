package gopnet

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type stream struct {
	Conn *websocket.Conn

	onError func(error) bool // continue?
	onMsg   func([]byte)
}

func Stream(url string, header http.Header) (*stream, error) {
	c, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		return nil, err
	}

	s := stream{Conn: c, onError: onError, onMsg: nil}

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		<-interrupt
		s.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}()

	return &s, nil
}

func (s *stream) OnMsg(f func([]byte)) *stream {
	p := s.onMsg
	s.onMsg = f

	if p == nil {
		go func() {
			for {
				_, message, err := s.Conn.ReadMessage()
				if err != nil {
					if ok := s.onError(err); !ok {
						return
					}

					continue
				}

				s.onMsg(message)
			}
		}()
	}

	return s
}

func (s *stream) OnError(f func(error) bool) *stream {
	s.onError = f

	return s
}

func onError(error) bool {
	return true
}
