package meshblu

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/octoblu/go-meshblu/config"
)

// Meshblu interfaces with a remote meshblu server
type Meshblu interface {
	SetAuth(uuid, token string)
	GetDevice(uuid string) ([]byte, error)
}

// Client interfaces with a remote meshblu server
type Client struct {
	uri, uuid, token string
}

// Dail constructs a new Meshblu instance and creates a connection
func Dail(uri string) (Meshblu, error) {
	return &Client{
		uri: uri,
	}, nil
}

// SetAuth sets the authentication
func (client *Client) SetAuth(uuid, token string) {
	client.uuid = uuid
	client.token = token
}

// GetDevice returns a byte response of the meshblu device
func (client *Client) GetDevice(uuid string) ([]byte, error) {
	return client.getRequest(fmt.Sprintf("/v2/devices/%s", uuid))
}

func (client *Client) getRequest(path string) ([]byte, error) {
	meshbluURL, err := config.ParseURL(client.uri)
	if err != nil {
		return nil, err
	}
	meshbluURL.SetPath(path)

	response, err := http.Get(meshbluURL.String())
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("Meshblu returned invalid response code: %v", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
