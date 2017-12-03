package cli

import (
	"github.com/stretchr/testify/suite"

	"github.com/sonm-io/marketplace/cli"
	"gopkg.in/tomb.v2"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// ListenAddr network address and port.
	ListenAddr = "127.0.0.1:9095"
)

// AppTestSuite defines a test suite for the application.
type AppTestSuite struct {
	suite.Suite

	sync.RWMutex
	App *cli.App

	shutdownCounter uint32
	t               tomb.Tomb
}

// SetupTest initializes application.
func (s *AppTestSuite) SetupTest() {
	s.Lock()
	defer s.Unlock()

	s.App = cli.NewApp(cli.WithListenAddr(ListenAddr))
	err := s.App.Init()
	s.NoError(err, "cannot initialize application")
}

// StartApp starts the application.
func (s *AppTestSuite) StartApp() {
	if s.App == nil {
		s.FailNow("application must be initialized")
	}

	s.False(s.IsAppRunning(), "application must not be running before starting")

	s.T().Log("Starting application")
	s.t.Go(s.App.Run)
	s.T().Log("Application started")

	s.True(s.t.Alive())
	s.True(s.IsAppRunning(), "application must be running after starting")
}

// StopApp gracefully stops the application.
func (s *AppTestSuite) StopApp() error {

	s.T().Log("Stopping application")
	// Do nothing if StopApp is already been called.
	// if s.shutdownCounter == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.shutdownCounter, 0, 1) {
		return nil
	}

	s.t.Kill(nil)

	// wait for application to stop
	s.waitShutdown()

	s.T().Log("Application stopped")
	return s.t.Wait()
}

func (s *AppTestSuite) waitShutdown() {
	s.T().Log("Waiting for graceful stop")
	<-s.t.Dying()
	s.App.Stop()
}

// IsAppRunning Checks whether the application is listening to the port (serving).
func (s *AppTestSuite) IsAppRunning() bool {
	for i := 0; i < 100; i++ {
		conn, _ := net.DialTimeout("tcp", ListenAddr, 10*time.Millisecond)
		if conn != nil {
			conn.Close()
			return true
		}
	}

	return false
}
