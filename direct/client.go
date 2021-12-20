package direct

import (
	"context"
	"sync"

	"github.com/skycoin/dmsg/cipher"
	"github.com/skycoin/dmsg/disc"

	"github.com/skycoin/skycoin/src/util/logging"
)

// directClient represents a client that doesnot communicates with a dmsg-discovery, instead directly gets the dmsg-server info via the user or is hardcoded, it
// implements APIClient
type directClient struct {
	entries map[cipher.PubKey]*disc.Entry
	mx      sync.RWMutex
}

// NewClient constructs a new APIClient that communicates with discovery via http.
func NewClient(entries []*disc.Entry, log *logging.Logger) disc.APIClient {
	entriesMap := make(map[cipher.PubKey]*disc.Entry)
	for _, entry := range entries {
		entriesMap[entry.Static] = entry
	}
	log.WithField("func", "direct.NewClient").
		Debug("Created Direct client.")
	return &directClient{
		entries: entriesMap,
	}
}

// Entry retrieves an entry associated with the given public key.
func (c *directClient) Entry(ctx context.Context, pubKey cipher.PubKey) (*disc.Entry, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, entry := range c.entries {
		if entry.Static == pubKey {
			return entry, nil
		}
	}
	return &disc.Entry{}, nil
}

// PostEntry adds a new Entry.
func (c *directClient) PostEntry(ctx context.Context, entry *disc.Entry) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[entry.Static] = entry
	return nil
}

// DelEntry deletes an Entry.
func (c *directClient) DelEntry(ctx context.Context, entry *disc.Entry) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.entries, entry.Static)
	return nil
}

// PutEntry updates Entry.
func (c *directClient) PutEntry(ctx context.Context, _ cipher.SecKey, entry *disc.Entry) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.entries[entry.Static] = entry
	return nil
}

// AvailableServers returns list of available servers.
func (c *directClient) AvailableServers(ctx context.Context) (entries []*disc.Entry, err error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, entry := range c.entries {
		if entry.Server != nil {
			entries = append(entries, entry)
		}
	}
	return entries, nil
}
