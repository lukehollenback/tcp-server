package tcpserver

import (
	"net"
	"testing"
	"time"
)

const TestServerAddress = "localhost:9999"
const TestMessage = "This is a test message. Here is a number: 12345.67890!\n"

func buildTestServer() *Server {
	return New(TestServerAddress)
}

func TestBasicLifecycle(t *testing.T) {
	//
	// Define variables upon which we will state and assert proper functionality.
	//
	var messageReceived bool
	var messageText string
	var newClient bool
	var connectinClosed bool

	//
	// Create a new server and register event handlers that will set variables against which we can
	// run some assertions.
	//
	server := buildTestServer()

	server.OnNewClient(func(c *Client) {
		newClient = true
	})

	server.OnNewMessage(func(c *Client, message string) {
		messageReceived = true
		messageText = message
	})

	server.OnClientConnectionClosed(func(c *Client, err error) {
		connectinClosed = true
	})

	go server.Start()

	//
	// Give the server some time to bind.
	//
	// NOTE: Although server.Listen() is a blocking call, we are running it asynchronously in a
	//  goroutine for the sake of this test.
	//
	time.Sleep(10 * time.Millisecond)

	//
	// Connect to the server as a new client and sent it a test message.
	//
	conn, err := net.Dial("tcp", TestServerAddress)

	if err != nil {
		t.Fatal("Failed to connect to test server")
	}

	conn.Write([]byte(TestMessage))

	conn.Close()

	//
	// Give the server a chance to recieve the new connection and the new message.
	//
	time.Sleep(10 * time.Millisecond)

	//
	// Assert that the server's handlers fired and that the expected messages came back.
	//
	if newClient != true {
		t.Error("The \"OnNewClient\" event handler never fired.")
	}

	if messageReceived != true {
		t.Error("The \"OnNewMessage\" event handler never fired.")
	}

	if messageText != TestMessage {
		t.Error("A message was recieved, but it was not equal to what was expected.")
	}

	if connectinClosed != true {
		t.Error("The \"OnClientConnectionClosed\" event handler never fired.")
	}

	server.Stop()
}

func TestBasicLifecycleAgain(t *testing.T) {
	//
	// Define variables upon which we will state and assert proper functionality.
	//
	var messageReceived bool
	var messageText string
	var newClient bool
	var connectinClosed bool

	//
	// Create a new server and register event handlers that will set variables against which we can
	// run some assertions.
	//
	server := buildTestServer()

	server.OnNewClient(func(c *Client) {
		newClient = true
	})

	server.OnNewMessage(func(c *Client, message string) {
		messageReceived = true
		messageText = message
	})

	server.OnClientConnectionClosed(func(c *Client, err error) {
		connectinClosed = true
	})

	go server.Start()

	//
	// Give the server some time to bind.
	//
	// NOTE: Although server.Listen() is a blocking call, we are running it asynchronously in a
	//  goroutine for the sake of this test.
	//
	time.Sleep(10 * time.Millisecond)

	//
	// Connect to the server as a new client and sent it a test message.
	//
	conn, err := net.Dial("tcp", TestServerAddress)

	if err != nil {
		t.Fatal("Failed to connect to test server")
	}

	conn.Write([]byte(TestMessage))

	conn.Close()

	//
	// Give the server a chance to recieve the new connection and the new message.
	//
	time.Sleep(10 * time.Millisecond)

	//
	// Assert that the server's handlers fired and that the expected messages came back.
	//
	if newClient != true {
		t.Error("The \"OnNewClient\" event handler never fired.")
	}

	if messageReceived != true {
		t.Error("The \"OnNewMessage\" event handler never fired.")
	}

	if messageText != TestMessage {
		t.Error("A message was recieved, but it was not equal to what was expected.")
	}

	if connectinClosed != true {
		t.Error("The \"OnClientConnectionClosed\" event handler never fired.")
	}

	server.Stop()
}
