package configurationscansummary

import (
	"github.com/Aryaman6492/storage/pkg/apis/softwarecomposition"
	"github.com/Aryaman6492/storage/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, storageImpl storage.Interface, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	dryRunnableStorage := genericregistry.DryRunnableStorage{Codec: nil, Storage: storageImpl}

	store := &genericregistry.Store{
		NewFunc:                   func() runtime.Object { return &softwarecomposition.ConfigurationScanSummary{} },
		NewListFunc:               func() runtime.Object { return &softwarecomposition.ConfigurationScanSummaryList{} },
		PredicateFunc:             MatchConfigurationScanSummary,
		DefaultQualifiedResource:  softwarecomposition.Resource("configurationscansummaries"),
		SingularQualifiedResource: softwarecomposition.Resource("configurationscansummary"),

		Storage: dryRunnableStorage,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		// TODO: define table converter that exposes more than name/creation timestamp
		TableConvertor: rest.NewDefaultTableConvertor(softwarecomposition.Resource("configurationscansummaries")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}
