// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the webhook client
type API interface {
	/*
	   CreateWebhookPolicyOfProject creates project webhook policy

	   This endpoint create a webhook policy if the project does not have one.
	*/
	CreateWebhookPolicyOfProject(ctx context.Context, params *CreateWebhookPolicyOfProjectParams) (*CreateWebhookPolicyOfProjectCreated, error)
	/*
	   DeleteWebhookPolicyOfProject deletes webhook policy of a project

	   This endpoint is aimed to delete webhookpolicy of a project.
	*/
	DeleteWebhookPolicyOfProject(ctx context.Context, params *DeleteWebhookPolicyOfProjectParams) (*DeleteWebhookPolicyOfProjectOK, error)
	/*
	   GetSupportedEventTypes gets supported event types and notify types

	   Get supportted event types and notify types.*/
	GetSupportedEventTypes(ctx context.Context, params *GetSupportedEventTypesParams) (*GetSupportedEventTypesOK, error)
	/*
	   GetWebhookPolicyOfProject gets project webhook policy

	   This endpoint returns specified webhook policy of a project.
	*/
	GetWebhookPolicyOfProject(ctx context.Context, params *GetWebhookPolicyOfProjectParams) (*GetWebhookPolicyOfProjectOK, error)
	/*
	   LastTrigger gets project webhook policy last trigger info

	   This endpoint returns last trigger information of project webhook policy.
	*/
	LastTrigger(ctx context.Context, params *LastTriggerParams) (*LastTriggerOK, error)
	/*
	   ListWebhookPoliciesOfProject lists project webhook policies

	   This endpoint returns webhook policies of a project.
	*/
	ListWebhookPoliciesOfProject(ctx context.Context, params *ListWebhookPoliciesOfProjectParams) (*ListWebhookPoliciesOfProjectOK, error)
	/*
	   UpdateWebhookPolicyOfProject updates webhook policy of a project

	   This endpoint is aimed to update the webhook policy of a project.
	*/
	UpdateWebhookPolicyOfProject(ctx context.Context, params *UpdateWebhookPolicyOfProjectParams) (*UpdateWebhookPolicyOfProjectOK, error)
}

// New creates a new webhook API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for webhook API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateWebhookPolicyOfProject creates project webhook policy

This endpoint create a webhook policy if the project does not have one.
*/
func (a *Client) CreateWebhookPolicyOfProject(ctx context.Context, params *CreateWebhookPolicyOfProjectParams) (*CreateWebhookPolicyOfProjectCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateWebhookPolicyOfProject",
		Method:             "POST",
		PathPattern:        "/projects/{project_name_or_id}/webhook/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateWebhookPolicyOfProjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateWebhookPolicyOfProjectCreated), nil

}

/*
DeleteWebhookPolicyOfProject deletes webhook policy of a project

This endpoint is aimed to delete webhookpolicy of a project.
*/
func (a *Client) DeleteWebhookPolicyOfProject(ctx context.Context, params *DeleteWebhookPolicyOfProjectParams) (*DeleteWebhookPolicyOfProjectOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteWebhookPolicyOfProject",
		Method:             "DELETE",
		PathPattern:        "/projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteWebhookPolicyOfProjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWebhookPolicyOfProjectOK), nil

}

/*
GetSupportedEventTypes gets supported event types and notify types

Get supportted event types and notify types.
*/
func (a *Client) GetSupportedEventTypes(ctx context.Context, params *GetSupportedEventTypesParams) (*GetSupportedEventTypesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSupportedEventTypes",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/webhook/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSupportedEventTypesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSupportedEventTypesOK), nil

}

/*
GetWebhookPolicyOfProject gets project webhook policy

This endpoint returns specified webhook policy of a project.
*/
func (a *Client) GetWebhookPolicyOfProject(ctx context.Context, params *GetWebhookPolicyOfProjectParams) (*GetWebhookPolicyOfProjectOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWebhookPolicyOfProject",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWebhookPolicyOfProjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebhookPolicyOfProjectOK), nil

}

/*
LastTrigger gets project webhook policy last trigger info

This endpoint returns last trigger information of project webhook policy.
*/
func (a *Client) LastTrigger(ctx context.Context, params *LastTriggerParams) (*LastTriggerOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LastTrigger",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/webhook/lasttrigger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LastTriggerReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LastTriggerOK), nil

}

/*
ListWebhookPoliciesOfProject lists project webhook policies

This endpoint returns webhook policies of a project.
*/
func (a *Client) ListWebhookPoliciesOfProject(ctx context.Context, params *ListWebhookPoliciesOfProjectParams) (*ListWebhookPoliciesOfProjectOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListWebhookPoliciesOfProject",
		Method:             "GET",
		PathPattern:        "/projects/{project_name_or_id}/webhook/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListWebhookPoliciesOfProjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListWebhookPoliciesOfProjectOK), nil

}

/*
UpdateWebhookPolicyOfProject updates webhook policy of a project

This endpoint is aimed to update the webhook policy of a project.
*/
func (a *Client) UpdateWebhookPolicyOfProject(ctx context.Context, params *UpdateWebhookPolicyOfProjectParams) (*UpdateWebhookPolicyOfProjectOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateWebhookPolicyOfProject",
		Method:             "PUT",
		PathPattern:        "/projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateWebhookPolicyOfProjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateWebhookPolicyOfProjectOK), nil

}