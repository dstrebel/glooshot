// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type ExperimentClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Experiment, error)
	Write(resource *Experiment, opts clients.WriteOpts) (*Experiment, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (ExperimentList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan ExperimentList, <-chan error, error)
}

type experimentClient struct {
	rc clients.ResourceClient
}

func NewExperimentClient(rcFactory factory.ResourceClientFactory) (ExperimentClient, error) {
	return NewExperimentClientWithToken(rcFactory, "")
}

func NewExperimentClientWithToken(rcFactory factory.ResourceClientFactory, token string) (ExperimentClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Experiment{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Experiment resource client")
	}
	return NewExperimentClientWithBase(rc), nil
}

func NewExperimentClientWithBase(rc clients.ResourceClient) ExperimentClient {
	return &experimentClient{
		rc: rc,
	}
}

func (client *experimentClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *experimentClient) Register() error {
	return client.rc.Register()
}

func (client *experimentClient) Read(namespace, name string, opts clients.ReadOpts) (*Experiment, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Experiment), nil
}

func (client *experimentClient) Write(experiment *Experiment, opts clients.WriteOpts) (*Experiment, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(experiment, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Experiment), nil
}

func (client *experimentClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *experimentClient) List(namespace string, opts clients.ListOpts) (ExperimentList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToExperiment(resourceList), nil
}

func (client *experimentClient) Watch(namespace string, opts clients.WatchOpts) (<-chan ExperimentList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	experimentsChan := make(chan ExperimentList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				experimentsChan <- convertToExperiment(resourceList)
			case <-opts.Ctx.Done():
				close(experimentsChan)
				return
			}
		}
	}()
	return experimentsChan, errs, nil
}

func convertToExperiment(resources resources.ResourceList) ExperimentList {
	var experimentList ExperimentList
	for _, resource := range resources {
		experimentList = append(experimentList, resource.(*Experiment))
	}
	return experimentList
}
