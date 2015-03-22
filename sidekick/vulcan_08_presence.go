package sidekick

import (
	"time"
	// "errors"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type Vulcan08Presence struct {
	EtcdEndpoint string
}

type VulcanSettingsKeepAlive struct {
	Period              time.Time
	MaxIdleConnsPerHost int
}

type VulcanSettingsTimeouts struct {
	Read         time.Time
	Dial         time.Time
	TLSHandshake time.Time
}

type VulcanBackendSettings struct {
	Timeouts  VulcanSettingsTimeouts
	KeepAlive VulcanSettingsKeepAlive
}

type VulcanBackend struct {
	Type     string
	Settings VulcanBackendSettings
}

func (v Vulcan08Presence) WasAdded() (err error) {
	client, err := NewEtcdClient(v.EtcdEndpoint)
	if err != nil {
		return err
	}

	kAPI := client.NewKeysAPI(client)
	// create a new key /foo with the value "bar"
	_, err = kAPI.Create(context.Background(), "/vulcan", "bar")
	if err != nil {
		return err
	}

	return
}

func (v Vulcan08Presence) WasRemoved() (err error) {
	return nil
}

// func NewKeysClient(c *client.Client) (*client.KeysAPI, error) {
//   kAPI := client.NewKeysAPI(c)
//   // create a new key /foo with the value "bar"
//   _, err = kAPI.Create(context.Background(), "/foo", "bar")
//   if err != nil {
//     // handle error
//   }
//   // delete the newly created key only if the value is still "bar"
//   _, err = kAPI.Delete(context.Background(), "/foo", &DeleteOptions{PrevValue: "bar"})
//   if err != nil {
//     // handle error
//   }
// }

func NewEtcdClient(etcdEndpoint string) (*client.Client, error) {
	cfg := client.Config{
		Endpoints: []string{etcdEndpoint},
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
