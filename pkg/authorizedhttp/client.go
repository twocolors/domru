package authorizedhttp

import (
	myhttp "github.com/ad/domru/pkg/domru/http"
	"log/slog"
	"net/http"
	"strconv"
)

type TokenRefreshError struct {
	Err error
}

func (e TokenRefreshError) Error() string {
	return e.Err.Error()
}

func NewTokenRefreshError(err error) TokenRefreshError {
	return TokenRefreshError{Err: err}
}

type TokenProvider interface {
	GetToken() (string, error)
}

type TokenRefresher interface {
	RefreshToken() error
}

type Client struct {
	DefaultClient  myhttp.HTTPClient
	TokenProvider  TokenProvider
	TokenRefresher TokenRefresher
	Logger         *slog.Logger

	loginUrl   string
	operatorId int
}

func NewClient(operatorId int, tokenProvider TokenProvider, tokenRefresher TokenRefresher) *Client {
	return &Client{
		operatorId:     operatorId,
		TokenProvider:  tokenProvider,
		TokenRefresher: tokenRefresher,
		DefaultClient:  http.DefaultClient,
		Logger:         slog.Default(),
		loginUrl:       "/pages/login.html",
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	token, err := c.TokenProvider.GetToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := c.DefaultClient.Do(req)
	if err != nil {
		c.Logger.With("error", err).With("url", req.URL).With("method", req.Method).With("headers", req.Header).Warn("Failed to send request")
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		// Refresh the token
		c.Logger.Debug("Token expired. Refreshing token...")
		err = c.TokenRefresher.RefreshToken()
		if err != nil {
			c.Logger.Warn("Failed to refresh token. Redirecting to login page", err.Error())
			return nil, NewTokenRefreshError(err)
		}

		// Retry the request with the new token
		newToken, err := c.TokenProvider.GetToken()
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+newToken)
		req.Header.Set("Operator", strconv.Itoa(c.operatorId))
		resp, err = c.DefaultClient.Do(req)
	}

	return resp, err
}
