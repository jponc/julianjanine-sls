package invitationscsv

import (
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	csvData []byte
}

type Invitation struct {
	InvitationCode string `csv:"InvitationCode"`
	Name           string `csv:"Name"`
}

func NewClient(csvData []byte) *Client {
	return &Client{csvData: csvData}
}

func (c *Client) ToInvitations() *[]Invitation {
	invitations := []Invitation{}

	err := gocsv.UnmarshalString(string(c.csvData), &invitations)
	if err != nil {
		log.Fatalf("Failed to unmarshal CSV into invitations: %w", err)
	}

	return &invitations
}
