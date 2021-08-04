package main

import (
	"fmt"
	"strange-secrets.com/mantra/algebra"
	"strange-secrets.com/mantra/render"
	"strange-secrets.com/mantra/scene"
	"strange-secrets.com/mantra/scene/camera"
	"strange-secrets.com/mantra/scene/geometry"
	"strange-secrets.com/mantra/scene/lighting"
	"strange-secrets.com/mantra/scene/shading"
)

const (
	TestOutputFileName = "output.png"
	TestImageName = "test"
)

var (
	SphereRadius = 4.0
	SphereSeparation = 9.0
	SphereColor = algebra.Vector3{0.8, 0.8, 0.8}
	FloorColor = algebra.Vector3{0.2, 0.1, 0.6}
	LightColor = algebra.Vector3{1.0, 1.0, 1.0}
	//LightColor = algebra.Vector3{0.53, 0.8, 0.92}		// Sky blue
)

func CreateSphereFlake() (*scene.Scene, error) {
	generator := scene.SphereFlake{
		Origin: algebra.Vector3{0, 4.0, -20.0},
		Radius: 5,
		Depth:  4,
		Scene:  scene.NewScene(),
	}

	world := generator.Generate()

	if cameraNode, err := world.CreateCameraNode(world.CreateUniqueName("camera"), &camera.Camera{
		Type:        camera.PerspectiveCamera,
		Near:        0.01,
		Far:         100.0,
		FieldOfView: algebra.ToRadians(70.0),
	}); err != nil {
		return nil, err
	} else {
		cameraNode.Location.Y = 10.0
		cameraNode.LookAt(generator.Origin)
	}

	if _, err := world.CreateLightNode(world.CreateUniqueName("light"), &lighting.Directional{
		Direction: algebra.UpVector3.Negate(),
		Color:     LightColor,
	}); err != nil {
		return nil, err
	}

	return world, nil
}

func CreateScene() (*scene.Scene, error) {
	world := scene.NewScene()	// TODO: Load from file

	if floorNode, err := world.CreateGeometryNode(world.CreateUniqueName("floor"), shading.CheckerboardMaterialName, &geometry.PlaneGeometry{
		Normal:   algebra.UpVector3,
		Distance: -0.5,
	}); err != nil {
		return nil, err
	} else {
//		floorNode.Surface.Material = &shading.GridColorMaterial{
//			Color:     algebra.Vector3{1.0, 1.0, 1.0},
//			LineColor: algebra.Vector3{0.4, 0.4, 0.4},
//			LineWidth: 0.05,
//			CellWidth: 2.5,
//		}
		checkerboard := floorNode.Surface.Material.(*shading.CheckerboardMaterial)
		checkerboard.ColorA = algebra.Vector3{0.6, 0.4, 0.2}
		checkerboard.ColorB = algebra.Vector3{0.4, 0.2, 0.6}
	}

	for loop := 0; loop < 5; loop++ {
		x := -(SphereSeparation * 2.0) + SphereSeparation * float64(loop)
		if sphereNode, err := world.CreateGeometryNode(world.CreateUniqueName("sphere"), shading.SolidColorMaterialName, &geometry.SphereGeometry{
			Radius: SphereRadius,
		}); err != nil {
			return nil, err
		} else {
			sphereNode.Location.X = x
			sphereNode.Location.Y = 4.0 + float64(loop)
			sphereNode.Location.Z = -20.0

			material := sphereNode.Surface.Material.(*shading.SolidColorMaterial)
			material.Color = SphereColor
		}
	}

	if cameraNode, err := world.CreateCameraNode("camera", &camera.Camera{
		Type:        camera.PerspectiveCamera,
		Near:        0.01,
		Far:         100.0,
		FieldOfView: algebra.ToRadians(70.0),
	}); err != nil {
		return nil, err
	} else {
		cameraNode.Location.X = 8
		cameraNode.Location.Y = 8.0
		cameraNode.Location.Z = -6.0
		cameraNode.LookAt(algebra.Vector3{0.0, 4.0, -20.0})
	}
/*
	if pointLight, err := world.CreateLightNode(world.CreateUniqueName("light"), &lighting.Point{
		Radius:	   100,
		Color:     LightColor,
	}); err != nil {
		return nil, err
	} else {
		pointLight.Location = algebra.Vector3{0, 4, -20.0}
	}
*/
	if _, err := world.CreateLightNode(world.CreateUniqueName("light"), &lighting.Directional{
		Direction: algebra.UpVector3.Negate(),
		Color:     LightColor,
	}); err != nil {
		return nil, err
	}

/*
	if _, err := world.CreateNode(world.CreateUniqueName("light"), &lighting.Hemisphere{
		Up: algebra.UpVector3,
		UpperColor: LightColor,
		LowerColor: algebra.Vector3{0.1, 0.2, 0.4},
	}); err != nil {
		return nil, err
	}
*/
	return world, nil
}

func main() {
	if world, err := CreateScene(); err == nil {
		renderer := render.NewRenderer()

		if target, err := renderer.CreateRenderTarget(TestImageName, render.TestImageWidth, render.TestImageHeight); err == nil {
			fmt.Println("Rendering scene...")
			if err := renderer.Render(TestImageName, "camera", world); err == nil {
				fmt.Println(fmt.Sprintf("Saving rendered image to \"%s\"", TestOutputFileName))
				_ = target.SavePngImage(TestOutputFileName)
				fmt.Println("Complete")
			}
		}
	}
}
