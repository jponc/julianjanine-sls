package tools

import (
	"context"

	"github.com/google/uuid"
	"github.com/jponc/julianjanine/internal/invitationscsv"
	"github.com/jponc/julianjanine/internal/types"
)

type Repository interface {
	SaveGuest(ctx context.Context, invitationCode string, guestId string, guest types.Guest) error
}

type InvitationsCSV interface {
	ToInvitations() *[]invitationscsv.Invitation
}

type Client struct {
	repository     Repository
	invitationsCSV InvitationsCSV
}

func NewClient(repository Repository, invitationsCSV InvitationsCSV) *Client {
	return &Client{
		repository:     repository,
		invitationsCSV: invitationsCSV,
	}
}

func (c *Client) RebuildInvitations(ctx context.Context) error {
	// NOTE Assume that the entire dynamodb table has already been truncated
	invitations := c.invitationsCSV.ToInvitations()

	for _, invitation := range *invitations {
		guestId := uuid.New().String()
		guest := types.Guest{
			Id:         guestId,
			Name:       invitation.Name,
			Attendance: types.AttendancePending,
		}

		c.repository.SaveGuest(ctx, invitation.InvitationCode, guestId, guest)
	}

	return nil
}
