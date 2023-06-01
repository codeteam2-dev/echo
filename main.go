// a simple TCP echo server
// this code was lifted from the official Go documentation at https://pkg.go.dev/net#example-Listener
package main

import (
	"io"
	"net"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting echo server...")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Error().Err(err).Msgf("Failed to accept connection")
			continue
		}
		log.Info().Msgf("Accepted connection from %s", conn.RemoteAddr())
		go func(conn net.Conn) {
			// Echo all incoming data.
			io.Copy(conn, conn)
			conn.Close()

		}(conn)
	}
}
