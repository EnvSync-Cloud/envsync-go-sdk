// Code generated by Fern. DO NOT EDIT.

package roles

import (
	context "context"
	sdk "github.com/EnvSync-Cloud/envsync-go-sdk/sdk"
	core "github.com/EnvSync-Cloud/envsync-go-sdk/sdk/core"
	internal "github.com/EnvSync-Cloud/envsync-go-sdk/sdk/internal"
	option "github.com/EnvSync-Cloud/envsync-go-sdk/sdk/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Retrieve all roles in the organization
func (c *Client) GetAllRoles(
	ctx context.Context,
	opts ...option.RequestOption,
) (sdk.RolesResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := baseURL + "/api/role"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response sdk.RolesResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create a new role in the organization
func (c *Client) CreateRole(
	ctx context.Context,
	request *sdk.CreateRoleRequest,
	opts ...option.RequestOption,
) (*sdk.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := baseURL + "/api/role"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve statistics about roles in the organization
func (c *Client) GetRoleStats(
	ctx context.Context,
	opts ...option.RequestOption,
) (*sdk.RoleStatsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := baseURL + "/api/role/stats"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.RoleStatsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Retrieve a specific role by ID
func (c *Client) GetRole(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*sdk.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/role/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Delete an existing role (non-master roles only)
func (c *Client) DeleteRole(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*sdk.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/role/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Update an existing role
func (c *Client) UpdateRole(
	ctx context.Context,
	id string,
	request *sdk.UpdateRoleRequest,
	opts ...option.RequestOption,
) (*sdk.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"http://localhost:8600",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/role/%v",
		id,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &sdk.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
