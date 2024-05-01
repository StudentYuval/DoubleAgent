package logic

// main logic - will be listening on a port and will be able to handle requests
// all requests should include a token in the header to make sure they are authorized

// Path: logic/main.go

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

const (
	// token to be used for authorization
	token = "0xdeadbeef"
	port  = ":54321"
)

// ExecuteCommand executes the command
func ExecuteCommand(command string) string {
	// execute the command as bash command - safe since we authorize the requests

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	return fmt.Sprintf("Command %s executed.\noutput: %s\nerror: %s", command, cmd.Stdout, cmd.Stderr)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// check if the token is correct
	if r.Header.Get("Authorization") != token {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// read the command from the request and execute it
	command := r.URL.Query().Get("command")

	result := ExecuteCommand(command)

	// write the result to the response
	_, _ = w.Write([]byte(result))
}

func StartServer() {

	// start the server
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(port, nil)

}
