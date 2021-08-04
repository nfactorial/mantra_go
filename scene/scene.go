package scene

import (
	"errors"
	"fmt"
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/scene/camera"
	"strange-secrets.com/mantra/scene/geometry"
	"strange-secrets.com/mantra/scene/lighting"
	"strange-secrets.com/mantra/scene/shading"
)

type Scene struct {
	MaterialFactory	*shading.Factory
	nodes			map[string]*Node
	geometry		[]*Node
	lights			[]*Node
	cameras			[]*Node
}

func NewScene() *Scene {
	return &Scene{
		MaterialFactory:	shading.NewMaterialFactory(),
		nodes:				make(map[string]*Node),
	}
}

// Retrieves the number of nodes contained within the scene
func (s *Scene) Count() int {
	return len(s.nodes)
}

// Retrieves a specific node associated with the supplied identifier. Returns nil if the node could not be found.
func (s *Scene) GetNode(name string) *Node {
	if node, ok := s.nodes[name]; ok {
		return node
	}
	return nil
}

// Given a base name for a node, this method finds a suitable name that is not currently in use within the scene.
func (s *Scene) CreateUniqueName(baseName string) string {
	if _, ok := s.nodes[baseName]; !ok {
		return baseName
	}

	for counter := 0;; counter++ {
		result := fmt.Sprintf("%s_%03d", baseName, counter)
		if _, ok := s.nodes[result]; !ok {
			return result
		}
	}
}

// Creates a geometric node within the scene and associates it with the specified name
func (s *Scene) CreateGeometryNode(name string, material string, content geometry.Geometry) (*Node, error) {
	if _, ok := s.nodes[name]; ok {
		return nil, fmt.Errorf("cannot create node \"%s\", name already in use", name)
	}

	node := NewNode(name, nodeTypeGeometry)
	node.Geometry = content
	node.Surface.Material = s.MaterialFactory.Create(material)

	s.nodes[name] = node
	s.geometry = append(s.geometry, node)

	return node, nil
}

func (s *Scene) CreateCameraNode(name string, info *camera.Camera) (*Node, error) {
	if info == nil {
		return nil, errors.New("cannot create camera node without camera information")
	}

	if _, ok := s.nodes[name]; ok {
		return nil, fmt.Errorf("cannot create node \"%s\", name already in use", name)
	}

	node := NewNode(name, nodeTypeCamera)
	node.Camera = info

	s.nodes[name] = node
	s.cameras = append(s.cameras, node)

	return node, nil
}

func (s *Scene) CreateLightNode(name string, light lighting.Light) (*Node, error) {
	if _, ok := s.nodes[name]; ok {
		return nil, fmt.Errorf("cannot create node \"%s\", name already in use", name)
	}

	node := NewNode(name, nodeTypeLight)
	node.light = light

	s.nodes[name] = node
	s.lights = append(s.lights, node)

	return node, nil
}

func (s *Scene) Illuminate(info shading.ShadeInfo) algebra.Vector3 {
	//result := algebra.ZeroVector3
	result := algebra.Vector3{info.Surface.Ka, info.Surface.Ka, info.Surface.Ka}

	for _, node := range s.lights {
		shadow := node.light.Shadow(info, s.TraceShadow)
		if shadow < 1 {
			diffuse := node.light.Diffuse(info).MultiplyScalar(info.Surface.Kd)
			specular := node.light.Specular(info).MultiplyScalar(info.Surface.Ks)

			result = result.Add(node.light.Ambient(info).MultiplyScalar(info.Surface.Ka))
			result = result.Add(diffuse.Add(specular).MultiplyScalar(1 - shadow))
		}
	}

	return result
}

// Traces a ray through the scene and determines whether or not it hits an object
func (s *Scene) TraceShadow(ray algebra.Ray, valid algebra.Range) bool {
	for _, node := range s.nodes {
		if node.CastShadows {
			if hitResult := node.HitTest(ray, valid); hitResult != geometry.InvalidHitResult {
				return true
			}
		}
	}

	return false
}

// Traces a ray through the scene and computes the color of the surface
func (s *Scene) ShadeRay(info TraceInfo) algebra.Vector3 {
	traceResult := s.TraceRay(info.Ray, algebra.FrontRange)
	if traceResult.HitResult == geometry.InvalidHitResult {
		return algebra.Vector3{0.53, 0.8, 0.92}		// Sky blue
		//return algebra.ZeroVector3
	}

	shadeInfo := shading.ShadeInfo{
		Point:    traceResult.HitResult.Location,
		Normal:   traceResult.HitResult.Normal,
		Incident: info.Ray.Direction,
		View:     info.View,
		Surface:  traceResult.Node.Surface,
	}

	surfaceColor := shadeInfo.Surface.Material.Evaluate(shadeInfo)
	reflection := algebra.ZeroVector3

	if info.Depth < info.MaximumDepth {
		reflection = s.ShadeRay(info.CreateReflection(traceResult.HitResult.Location, traceResult.HitResult.Normal))
		// TODO: multiply by reflection coefficient
		reflection = surfaceColor.Multiply(reflection)
	}

	// Finally, apply illumination
	surfaceColor = surfaceColor.Multiply(s.Illuminate(shadeInfo)).MultiplyScalar(1.0 - shadeInfo.Surface.Kr)
	surfaceColor = surfaceColor.Add(reflection)
	//result.Color = result.Color.MultiplyScalar(1.0 / math.Pi).Saturate()

	return surfaceColor
}

// Traces a ray through the scene and finds the first place where it intersects a geometric object
func (s *Scene) TraceRay(ray algebra.Ray, validRange algebra.Range) TraceResult {
	result := TraceResult{
		Node:      nil,
		HitResult: geometry.InvalidHitResult,
	}

	for _, node := range s.geometry {
		hitResult := node.HitTest(ray, validRange)
		if hitResult != geometry.InvalidHitResult && hitResult.Distance < result.HitResult.Distance {
			result.Node = node
			result.HitResult = hitResult
			validRange.End = hitResult.Distance
		}
	}

	return result
}
