package gitlab

import (
	"context"

	"github.com/ovh/cds/sdk"
)

func (c *gitlabClient) CurrentUser(ctx context.Context) (sdk.VCSUser, error) {
	user, _, err := c.client.Users.CurrentUser()
	if err != nil {
		return sdk.VCSUser{}, err
	}

	return sdk.VCSUser{
		Login:     user.Username,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		MFA:       user.TwoFactorEnabled,
		Name:      user.Name,
	}, nil
}
