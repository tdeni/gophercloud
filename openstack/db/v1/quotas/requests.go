package quotas

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retrieves the details of quotas for a specified tenant.
func Get(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	resp, err := client.Get(ctx, baseURL(client, projectID), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToRecordSetUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the base attributes that may be updated on Quotas.
type UpdateOpts struct {
	Instances *int `json:"instances,omitempty"`
	Ram       *int `json:"ram,omitempty"`
	Backups   *int `json:"backups,omitempty"`
	Volumes   *int `json:"volumes,omitempty"`
}

// ToRecordSetUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToRecordSetUpdateMap() (map[string]any, error) {
	b, err := gophercloud.BuildRequestBody(opts, "quotas")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Update updates quotas
func Update(ctx context.Context, client *gophercloud.ServiceClient, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRecordSetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := client.Put(ctx, baseURL(client, projectID), &b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
