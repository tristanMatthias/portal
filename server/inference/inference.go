package inference

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"portal/server/lib"
	"sync"
	"syscall"
	"time"
)

func Init() {
	checkEnvironment()
	setupFiles()

	wg := sync.WaitGroup{}
	runInferenceServer(&wg)
	wg.Wait()
}

func checkEnvironment() {
	// Check if Python/pip are installed
	if _, err := exec.LookPath("python"); err != nil {
		panic("Python is not installed")
	}

	if _, err := exec.LookPath("pip"); err != nil {
		panic("pip is not installed")
	}

	// Install Python packages
	packages := []string{"torch", "transformers", "fastapi", "uvicorn", "pydantic"}
	cmd := exec.Command("pip", append([]string{"install"}, packages...)...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic("Error installing Python packages")
	} else {
		fmt.Println("Installed Python packages:", packages)
	}

	fmt.Println("Python environment is ready")
}

//go:embed server.py
var dataServer embed.FS
var fnServer = "server.py"
var fpServer = lib.ConfigPath("server", fnServer)

//go:embed Chat.py
var dataChat embed.FS
var fnChat = "Chat.py"
var fpChat = lib.ConfigPath("server", fnChat)

func setupFiles() {
	// Create the server directory if it doesn't exist
	os.MkdirAll(lib.ConfigPath("server"), 0755)

	if _, err := os.Stat(fpServer); os.IsNotExist(err) {
		// Write the file to disk if it doesn't exist
		data, _ := dataServer.ReadFile(fnServer)
		if err := ioutil.WriteFile(fpServer, data, 0644); err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(fpChat); os.IsNotExist(err) {
		// Write the file to disk if it doesn't exist
		data, _ := dataChat.ReadFile(fnChat)
		if err := ioutil.WriteFile(fpChat, data, 0644); err != nil {
			panic(err)
		}
	}

	fmt.Println("Python files are ready")
}

func runInferenceServer(wg *sync.WaitGroup) {
	// Check if port 9997 is available
	if _, err := net.Dial("tcp", "localhost:9997"); err == nil {
		fmt.Println("Inference service is already running")
		return
	}
	defer wg.Done()

	fmt.Println("Starting inference service...")
	// create a channel to receive signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// create a channel to stop the service
	stop := make(chan struct{})

	// create a channel to receive the command result
	done := make(chan error, 1)

	// start the service in a Goroutine
	go func() {
		// create a command to start the Python service
		cmd := exec.Command("python", fpServer)

		// set the output to Stdout so we can see the output
		cmd.Stdout = os.Stdout

		// start the command
		err := cmd.Start()
		if err != nil {
			fmt.Println("Failed to start service:", err)
			done <- err
			return
		}

		// wait for the service to stop or the stop signal to be received
		err = cmd.Wait()
		done <- err
	}()

	// wait for a signal or for the command to finish
	select {
	case <-sigs:
		// stop the service
		close(stop)
		fmt.Println("Waiting for service to stop...")
		time.Sleep(time.Second)
	case err := <-done:
		// the service has stopped
		if err != nil {
			fmt.Println("Inference service has stopped with an error:", err)
		} else {
			fmt.Println("Inference service has stopped.")
		}
	}
}
