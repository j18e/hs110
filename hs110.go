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

// Plug represents an HS110 plug, and exposes various methods for interacting
// with it.
type Plug struct {
	address string
	timeout time.Duration
	mtx     *sync.Mutex
	lastCmd time.Time
}

// NewPlug initializes and tests the connection to a plug.
func NewPlug(addr string) *Plug {
	return &Plug{
		address: addr + ":9999",
		timeout: time.Second * 5,
		mtx:     &sync.Mutex{},
	}
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
	var cmd Cmd
	if state {
		cmd = cmdOn
		log.Debug("turning on plug")
	} else {
		cmd = cmdOff
		log.Debug("turning off plug")
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
func (p *Plug) sendCmd(cmd Cmd) ([]byte, error) {
	// protect against sending too many commands at once
	p.mtx.Lock()
	defer func() {
		p.lastCmd = time.Now()
		p.mtx.Unlock()
	}()
	if time.Since(p.lastCmd) < time.Millisecond*500 {
		time.Sleep(time.Millisecond * 500)
	}

	res := make([]byte, 2048)

	// connect to plug
	conn, err := net.DialTimeout("tcp", p.address, p.timeout)
	if err != nil {
		return res, fmt.Errorf("connecting to plug: %w", err)
	}
	defer conn.Close()

	// set timeout
	if err := conn.SetDeadline(time.Now().Add(p.timeout)); err != nil {
		return res, fmt.Errorf("setting timeout: %w", err)
	}

	// encrypt payload, write data
	payload := encrypt([]byte(cmd))
	if _, err := conn.Write(payload); err != nil {
		return res, fmt.Errorf("writing payload: %w", err)
	}

	// receive, decrypt response
	i, err := conn.Read(res)
	if err != nil {
		return res, err
	}
	decrypted := decrypt(res[:i]) // only include the bytes that were read
	return decrypted, nil
}

// encrypt follows the autokey cipher used by the HS110 to encrypt commands.
func encrypt(bx []byte) []byte {
	key := 171
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(len(bx))) // equivalent in python: struct.pack('>I', len(cmd))

	for i := range bx {
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
