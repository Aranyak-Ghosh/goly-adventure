package providers

type Dependency struct {
	ServiceConstructor func() interface{}
	Dependency []func() interface{}
}

type DependencyProvider struct {
	configContainer *ConfigContainer
	provider        map[string]Dependency
}

func (provider *DependencyProvider) GetService(serviceIdentifier string) (interface{}, error) {
	if provider==nil {		
		return nil, error.new("Dependency Provider is not initialized")
	}	
	serviceDependency, ok := provider.provider[serviceIdentifier]

	// array of interfaces
	var serviceDependencies []interface{}

	for _, dependency := range dependencies {




	if !ok {
		return nil, error.new("Service not found")
	return service, nil
}

func (provider *DependencyProvider) getService(serviceIdentifier string) (interface{}, error){
	if provider==nil {		
		return nil, error.new("Dependency Provider is not initialized")
	}	
	serviceDependency, ok := provider.provider[serviceIdentifier]
	
	if !ok {
		return nil, error.new("Service not found")
	}

	
	

}

func (provider *DependencyProvider) RegisterService(configContainer *ConfigContainer) {

func NewDependencyProvider(environment string) (*DependencyProvider, error) {
	configContainer, err := NewConfigContainer(environment)
	if err != nil {
		return nil, err
	}
	provider := map[string]interface{}{}
	return &DependencyProvider{configContainer, provider}, nil
}
