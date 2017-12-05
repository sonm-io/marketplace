package cli

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/stretchr/testify/suite"
	"gopkg.in/tomb.v2"

	"github.com/sonm-io/marketplace/cli"
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

// SetupSuite runs once on suit initialization.
func (s *AppTestSuite) SetupSuite() {
	s.Lock()
	defer s.Unlock()

	s.App = cli.NewApp(cli.WithListenAddr(ListenAddr))
	err := s.App.Init()
	s.NoError(err, "cannot initialize application")
}

// SetupTest initializes application before each test.
func (s *AppTestSuite) SetupTest() {
	s.StartApp()
}

// TearDownTest stops the application after each test.
func (s *AppTestSuite) TearDownTest() {
	s.StopApp() //nolint
}

// StartApp starts the application.
func (s *AppTestSuite) StartApp() {
	s.Require().NotNil(s.App, "application must be initialized")
	s.False(s.IsAppRunning(), "application must not be running before starting")

	s.t.Go(s.App.Run)

	s.True(s.t.Alive())
	s.True(s.IsAppRunning(), "application must be running after starting")
}

// StopApp gracefully stops the application.
func (s *AppTestSuite) StopApp() error {

	// Do nothing if StopApp is already been called.
	// if s.shutdownCounter == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.shutdownCounter, 0, 1) {
		return nil
	}

	// stop the server goroutine
	s.t.Kill(nil)

	// wait for application to stop
	s.waitShutdown()

	return s.t.Wait()
}

func (s *AppTestSuite) waitShutdown() {
	<-s.t.Dying()
	s.App.Stop()
}

// IsAppRunning Checks whether the application is listening to the port (serving).
func (s *AppTestSuite) IsAppRunning() bool {
	for i := 0; i < 100; i++ {
		conn, _ := net.DialTimeout("tcp", ListenAddr, 10*time.Millisecond) //nolint
		if conn != nil {
			conn.Close() //nolint
			return true
		}
	}

	return false
}
