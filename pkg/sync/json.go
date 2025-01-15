package sync

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/yznts/zen/v3/conv"
	"github.com/yznts/zen/v3/errx"
)

// Json is a json file repository.
// Instead of using a database, we will use a json file to store the data.
// This is a simple start implementation, and we will improve/replace it later.
type Json struct {
	// Path to the json file
	Path string
	// Lock to prevent concurrent access
	Lock *sync.Mutex
	// State of the json file
	State map[string]map[string]string
}

// read reads the json file.
func (j *Json) read() error {
	return json.NewDecoder(errx.Must(os.Open(j.Path))).Decode(&j.State)
}

// write writes the json file.
func (j *Json) write() error {
	return json.NewEncoder(errx.Must(os.Create(j.Path))).Encode(j.State)
}

// SetSrc sets the source video for a player.
func (j *Json) SetSrc(id string, src string) error {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return err
	}
	// Create the player object if it does not exist
	if j.State[id] == nil {
		j.State[id] = make(map[string]string)
	}
	// Set the source video
	j.State[id]["src"] = src
	// Write the json file
	if err := j.write(); err != nil {
		return err
	}
	// Return nil
	return nil
}

// GetSrc gets the source video for a player.
func (j *Json) GetSrc(id string) (string, error) {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return "", err
	}
	// Return an error if the player does not exist
	if j.State[id] == nil {
		return "", fmt.Errorf("player %s does not exist", id)
	}
	// If no source video is set, return an error
	if j.State[id]["src"] == "" {
		return "", fmt.Errorf("player %s does not have a source", id)
	}
	// Return the source video
	return j.State[id]["src"], nil
}

// SetSec sets the video seek state.
func (j *Json) SetSec(id string, sec int) error {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return err
	}
	// Create the player object if it does not exist
	if j.State[id] == nil {
		j.State[id] = make(map[string]string)
	}
	// Set the video seek state
	j.State[id]["sec"] = fmt.Sprintf("%d", sec)
	// Write the json file
	if err := j.write(); err != nil {
		return err
	}
	// Return nil
	return nil
}

// GetSec gets the video seek state.
func (j *Json) GetSec(id string) (int, error) {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return 0, err
	}
	// Return an error if the player does not exist
	if j.State[id] == nil {
		return 0, fmt.Errorf("player %s does not exist", id)
	}
	// If no video seek state is set, return 0
	if sec, exists := j.State[id]["sec"]; !exists {
		return 0, nil
	} else {
		return conv.Int(sec), nil
	}
}

// SetCmd sets the command for a player.
func (j *Json) SetCmd(id string, cmd string) error {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return err
	}
	// Create the player object if it does not exist
	if j.State[id] == nil {
		j.State[id] = make(map[string]string)
	}
	// Set the command
	j.State[id]["cmd"] = cmd
	// Write the json file
	if err := j.write(); err != nil {
		return err
	}
	// Return nil
	return nil
}

// GetCmd gets the command for a player.
func (j *Json) GetCmd(id string) (string, error) {
	// Acquire the lock
	j.Lock.Lock()
	defer j.Lock.Unlock()
	// Read the json file
	if err := j.read(); err != nil {
		return "", err
	}
	// Return an error if the player does not exist
	if j.State[id] == nil {
		return "", fmt.Errorf("player %s does not exist", id)
	}
	// If no command is set, return an error
	if j.State[id]["cmd"] == "" {
		return "", fmt.Errorf("player %s does not have a command", id)
	}
	// Return the command
	return j.State[id]["cmd"], nil
}

// NewJson creates a new json file repository.
func NewJson(path string) *Json {
	return &Json{
		Path: path,
	}
}
