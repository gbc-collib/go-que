package main

import (
	"net"
	"sync"
)

type Config struct {
	ListenAddr string
}

type Message struct {
	Offset int
	Value  string
}

type Topic struct {
	messages []Message
	mu       sync.Mutex
	offset   int
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = ":5001"
	}
	return &Server{
		Config: cfg,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	return nil
}

func (s *Server) loop() {
	for {
		select {}
	}
}
