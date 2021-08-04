package shading

import (
	"fmt"
)

type CreateMaterial func() Material

type Factory struct {
	registry	map[string] CreateMaterial
}

func NewMaterialFactory() *Factory {
	factory := &Factory{
		registry: make(map[string] CreateMaterial),
	}

	_ = factory.Register(CheckerboardMaterialName, NewCheckerboardMaterial)
	_ = factory.Register(SolidColorMaterialName, NewSolidColorMaterial)
	_ = factory.Register(StripeMaterialName, NewStripeMaterial)
	_ = factory.Register(GridMaterialName, NewGridMaterial)

	return factory
}

func (f *Factory) Create(name string) Material {
	if creator, ok := f.registry[name]; ok {
		return creator()
	}

	return nil
}

func (f *Factory) Register(name string, creator CreateMaterial) error {
	if creator == nil {
		return fmt.Errorf("cannot register material \"%s\" without creation method", name)
	}

	if _, ok := f.registry[name]; ok {
		return fmt.Errorf("cannot register material \"%s\" name already in use", name)
	}

	f.registry[name] = creator
	return nil
}

func (f *Factory) Unregister(name string) error {
	if _, ok := f.registry[name]; !ok {
		return fmt.Errorf("cannot unregister material \"%s\", name could not be found", name)
	}

	delete(f.registry, name)
	return nil
}
