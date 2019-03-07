package api

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/ovh/cds/engine/api/repositoriesmanager"

	"github.com/gorilla/mux"
	"github.com/ovh/cds/engine/vcs/github"

	"github.com/ovh/cds/engine/api/accesstoken"
	"github.com/ovh/cds/engine/service"
	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/jws"
	"github.com/ovh/cds/sdk/log"
)

func (api *API) loginUserCallbackHandler() service.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// api.Cache.SetWithTTL("api:loginUserHandler:RequestToken:"+loginUserRequest.RequestToken, token, 30*60)
		var request sdk.UserLoginCallbackRequest
		if err := service.UnmarshalBody(r, &request); err != nil {
			return sdk.WithStack(err)
		}

		var accessToken sdk.AccessToken
		if !api.Cache.Get("api:loginUserHandler:RequestToken:"+request.RequestToken, &accessToken) {
			return sdk.ErrNotFound
		}

		pk, err := jws.NewPublicKeyFromPEM(request.PublicKey)
		if err != nil {
			log.Debug("unable to read public key: %v", string(request.PublicKey))
			return sdk.WithStack(err)
		}

		var x sdk.AccessTokenRequest
		if err := jws.Verify(pk, request.RequestToken, &x); err != nil {
			return sdk.WithStack(err)
		}

		jwt, err := accesstoken.Regen(&accessToken)
		if err != nil {
			return sdk.WithStack(err)
		}

		w.Header().Add("X-CDS-JWT", jwt)

		return service.WriteJSON(w, accessToken, http.StatusOK)
	}
}

const (
	AuthTypeLocal   = "Builtin Authentication"
	AuthTypeLDAP    = "LDAP"
	AuthTypeCorpSSO = "Corporate SSO"
	AuthTypeGithub  = "Github"
)

func (api *API) loginSSORedirectHandler() service.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var providers []sdk.AuthProvider
		var authCfg = &api.Config.Auth

		privKey, err := jws.NewPrivateKeyFromPEM([]byte(api.Config.Auth.RSAPrivateKey))
		if err != nil {
			return sdk.WithStack(err)
		}

		signer, err := jws.NewSigner(privKey)
		if err != nil {
			return sdk.WithStack(err)
		}

		if authCfg.Local.Enabled {
			providers = append(providers, sdk.AuthProvider{
				Name: "Local",
				Type: AuthTypeLocal,
				Icon: "database",
			})
		}

		if authCfg.LDAP.Enabled {
			providers = append(providers, sdk.AuthProvider{
				Name: authCfg.LDAP.Name,
				Type: AuthTypeLDAP,
				Icon: "address book",
			})
		}

		if authCfg.CorporateSSO.Enabled {
			providers = append(providers, sdk.AuthProvider{
				Name: authCfg.CorporateSSO.Name,
				Type: AuthTypeCorpSSO,
				Icon: "handshake outline",
			})
		}

		if authCfg.Github.Enabled {
			var authRequest = sdk.AuthRequest{
				AuthProvider: AuthTypeGithub,
				State:        sdk.UUID(),
			}

			state, err := jws.Sign(signer, authRequest)
			if err != nil {
				return sdk.WithStack(err)
			}

			callBackURL, err := url.Parse(api.Config.URL.API)
			if err != nil {
				return sdk.WithStack(err)
			}
			callBackURL.Path += "/repositories_manager/oauth2/callback"

			providers = append(providers, sdk.AuthProvider{
				Type:           AuthTypeGithub,
				Icon:           "github",
				RedirectURL:    "https://github.com/login/oauth/authorize",
				RedirectMethod: http.MethodGet,
				ContentType:    "application/x-www-form-urlencoded",
				Body: map[string]interface{}{
					"client_id":    authCfg.Github.ClientID,
					"scope":        strings.Join(github.RequestedScope, " "),
					"state":        state,
					"redirect_uri": callBackURL.String(),
				},
			})
		}

		return service.WriteJSON(w, providers, http.StatusOK)
	}
}

func (api *API) loginSSORedirectCallbackHandler() service.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		authProvider := vars["authProvider"]

		switch authProvider {
		case AuthTypeGithub:
			code := r.FormValue("code")
			state := r.FormValue("state")

			privKey, err := jws.NewPrivateKeyFromPEM([]byte(api.Config.Auth.RSAPrivateKey))
			if err != nil {
				return sdk.WithStack(err)
			}

			var authRequest sdk.AuthRequest
			if err := jws.Verify(&privKey.PublicKey, state, &authRequest); err != nil {
				return sdk.NewErrorFrom(sdk.ErrUnauthorized, "request verification failed: %v", err)
			}

			// Get the user information
			// TODO authProvider must be the same than the vcs server name
			vcsServer, err := repositoriesmanager.NewVCSServerConsumer(api.mustDB, api.Cache, authProvider)
			if err != nil {
				return sdk.WithStack(err)
			}

			token, secret, err := vcsServer.AuthorizeToken(ctx, state, code)
			if err != nil {
				return sdk.WithStack(err)
			}

			// TODO Create a project ???

			vcsClient, err := vcsServer.GetAuthorizedClient(ctx, token, secret)
			if err != nil {
				return sdk.WithStack(err)
			}

			currentUser, err := vcsClient.CurrentUser(ctx)
			if err != nil {
				return sdk.WithStack(err)
			}

			user.FindByEmail()

		default:
			return sdk.ErrUnauthorized
		}

		http.Redirect(w, r, "", http.StatusTemporaryRedirect) // ...

		return nil
	}
}
