package hs110

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Plug represents an HS110 plug, and exposes various functions for interacting
// with it.
type Plug struct {
	address string
	timeout time.Duration
	conn    net.Conn
	mutex   *sync.Mutex
	lastCmd time.Time
}

// NewPlug initializes and tests the connection to a plug.
func NewPlug(addr string) (*Plug, error) {
	plug := &Plug{
		address: addr + ":9999",
		timeout: time.Second * 5,
		mutex:   &sync.Mutex{},
	}

	// connect to plug with 5 second timeout
	conn, err := net.DialTimeout("tcp", plug.address, time.Second*5)
	if err != nil {
		return plug, err
	}
	plug.conn = conn

	if _, err := plug.Status(); err != nil {
		return plug, fmt.Errorf("talking to plug: %w", err)
	}
	return plug, nil
}

// Close should be defered whenever a plug is initialized.
func (p *Plug) Close() {
	p.conn.Close()
}

// On turns on the plug.
func (p *Plug) On() error {
	return p.setState(true)
}

// Off turns off the plug.
func (p *Plug) Off() error {
	return p.setState(false)
}

func (p *Plug) setState(state bool) error {
	var cmd = `{"system": {"set_relay_state": {"state": %d}}}`
	if state {
		cmd = fmt.Sprintf(cmd, 1)
	} else {
		cmd = fmt.Sprintf(cmd, 0)
	}

	// send the command, get the response
	data, err := p.sendCmd(cmd)
	if err != nil {
		return fmt.Errorf("sending state: %w", err)
	}

	// decode response data
	var resp struct {
		System struct {
			SetRelayState struct {
				ErrorCode int `json:"err_code"`
			} `json:"set_relay_state"`
		} `json:"system"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("unmarshalling the response: %w", err)
	}

	// check for error code
	if code := resp.System.SetRelayState.ErrorCode; code != 0 {
		return fmt.Errorf("plug gave error code: %d", code)
	}
	return nil
}

// sendCmd handles the communication with the plug.
func (p *Plug) sendCmd(cmd string) ([]byte, error) {
	// protect against sending too many commands at once
	p.mutex.Lock()
	defer func() {
		p.mutex.Unlock()
		p.lastCmd = time.Now()
	}()
	if time.Since(p.lastCmd) < time.Second {
		time.Sleep(time.Second)
	}

	res := make([]byte, 2048)
	// set timeout
	if err := p.conn.SetDeadline(time.Now().Add(p.timeout)); err != nil {
		return res, fmt.Errorf("setting timeout: %w", err)
	}

	// encrypt payload, write data
	log.Debugf("sendCmd sending: %s", cmd)
	payload := encrypt([]byte(cmd))
	if _, err := p.conn.Write(payload); err != nil {
		return res, fmt.Errorf("writing payload: %w", err)
	}

	// receive, decrypt response
	i, err := p.conn.Read(res)
	if err != nil {
		return res, err
	}
	decrypted := decrypt(res[:i]) // only include the bytes that were read
	log.Debugf("sendCmd got response: %s", decrypted)
	return decrypted, nil
}

// encrypt follows the autokey cipher used by the HS110 to encrypt commands.
func encrypt(bx []byte) []byte {
	key := 171
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(len(bx))) // equivalent in python: struct.pack('>I', len(cmd))

	for i, _ := range bx {
		b := key ^ int(bx[i])
		key = b
		res = append(res, byte(b))
	}
	return res
}

// decrypt follows the autokey cipher used by the HS110 to decrypt commands.
func decrypt(bx []byte) []byte {
	key := 171
	var res []byte

	for i := 4; i < len(bx); i++ { // first 4 bytes are padding
		b := key ^ int(bx[i])
		key = int(bx[i])
		res = append(res, byte(b))
	}
	return res
}
