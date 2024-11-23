package main

var (
	clientStore = OAuthClientStore{
		OAuthClient{
			DisplayName: "Paranormal Dashboard",
			ID:          "haus.hai.paranormal.dashboard",
		},
	}
)

// OAuthClientStore stores known clients
type OAuthClientStore []OAuthClient

// GetClientByID gets a client by it's ID.
func (store OAuthClientStore) GetClientByID(id string) *OAuthClient {
	for _, client := range store {
		if client.ID == id {
			return &client
		}
	}

	return nil
}

// OAuthClient is a OAuth 2.0 client that is known by the server.
type OAuthClient struct {
	DisplayName string

	ID     string
	Secret string
}
