package main

import (
	"aoc2022/util"
	"fmt"
)

const inputFile = "../input.txt"

const ChamberWidth = 7

var SpawnOffset util.Coordinate = util.Coordinate{X: 2, Y: 3}

type Mesh []util.Coordinate

// min bound is always 0,0
type MeshTemplate struct {
	Mesh     Mesh
	MaxBound util.Coordinate
	Simple   bool
}

func CreateRockShapes() []MeshTemplate {
	const numRockShapes = 5
	meshes := make([]Mesh, numRockShapes)
	meshes[0] = Mesh{
		util.Coordinate{X: 0, Y: 0},
		util.Coordinate{X: 1, Y: 0},
		util.Coordinate{X: 2, Y: 0},
		util.Coordinate{X: 3, Y: 0},
	}
	meshes[1] = Mesh{
		util.Coordinate{X: 1, Y: 0},
		util.Coordinate{X: 0, Y: 1},
		util.Coordinate{X: 1, Y: 1},
		util.Coordinate{X: 2, Y: 1},
		util.Coordinate{X: 1, Y: 2},
	}
	meshes[2] = Mesh{
		util.Coordinate{X: 0, Y: 0},
		util.Coordinate{X: 1, Y: 0},
		util.Coordinate{X: 2, Y: 0},
		util.Coordinate{X: 2, Y: 1},
		util.Coordinate{X: 2, Y: 2},
	}
	meshes[3] = Mesh{
		util.Coordinate{X: 0, Y: 0},
		util.Coordinate{X: 0, Y: 1},
		util.Coordinate{X: 0, Y: 2},
		util.Coordinate{X: 0, Y: 3},
	}
	meshes[4] = Mesh{
		util.Coordinate{X: 0, Y: 0},
		util.Coordinate{X: 1, Y: 0},
		util.Coordinate{X: 0, Y: 1},
		util.Coordinate{X: 1, Y: 1},
	}
	meshTemplates := make([]MeshTemplate, numRockShapes)
	meshTemplates[0] = createMeshTemplate(meshes[0], true)
	meshTemplates[1] = createMeshTemplate(meshes[1], false)
	meshTemplates[2] = createMeshTemplate(meshes[2], false)
	meshTemplates[3] = createMeshTemplate(meshes[3], true)
	meshTemplates[4] = createMeshTemplate(meshes[4], true)
	return meshTemplates
}

func createMeshTemplate(mesh Mesh, simple bool) MeshTemplate {
	_, maxB := mesh.boundingBox()
	return MeshTemplate{Mesh: mesh, MaxBound: maxB, Simple: simple}
}

func (m Mesh) boundingBox() (util.Coordinate, util.Coordinate) {
	minX := m[0].X
	maxX := m[0].X
	minY := m[0].Y
	maxY := m[0].Y
	for i := 1; i < len(m); i++ {
		coord := m[i]
		if coord.X < minX {
			minX = coord.X
		}
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y < minY {
			minY = coord.Y
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	return util.Coordinate{X: minX, Y: minY}, util.Coordinate{X: maxX, Y: maxY}
}

type Object2d struct {
	MeshTemplate MeshTemplate
	Position     util.Coordinate
	AbsMesh      Mesh
	AbsMinBound  util.Coordinate
	AbsMaxBound  util.Coordinate
}

func (o *Object2d) PersistAbsoluteCoords() {
	o.AbsMinBound = o.Position
	o.AbsMaxBound = o.Position.Sum(o.MeshTemplate.MaxBound)
	o.AbsMesh = o.transformToAbsolute()
}

func (o Object2d) transformToAbsolute() Mesh {
	absoluteCoords := make(Mesh, len(o.MeshTemplate.Mesh))
	for i := range o.MeshTemplate.Mesh {
		absoluteCoords[i] = o.Position.Sum(o.MeshTemplate.Mesh[i])
	}
	return absoluteCoords
}

func (o Object2d) CollidesWith(other Object2d) bool {
	boundsCollide := (o.AbsMaxBound.X >= other.AbsMinBound.X) && (other.AbsMaxBound.X >= o.AbsMinBound.X) &&
		(o.AbsMaxBound.Y >= other.AbsMinBound.Y) && (other.AbsMaxBound.Y >= o.AbsMinBound.Y)
	if !boundsCollide {
		return false
	}

	if o.MeshTemplate.Simple && other.MeshTemplate.Simple {
		return true
	}

	for _, c1 := range o.AbsMesh {
		for _, c2 := range other.AbsMesh {
			if c1 == c2 {
				return true
			}
		}
	}
	return false
}

func (o Object2d) CollidesWithAnyObject(objects []Object2d) bool {
	for i := len(objects) - 1; i >= 0; i-- {
		if o.CollidesWith(objects[i]) {
			return true
		}
	}
	return false
}

func GetJetEffect(jetPattern byte) int {
	changeX := 0
	if jetPattern == '>' {
		changeX = 1
	} else if jetPattern == '<' {
		changeX = -1
	}
	return changeX
}

func simulateRocks(numRocks int, rockShapes []MeshTemplate, chamberWidth int, spawnOffset util.Coordinate, jetPatterns string) int {
	chamberLeftVerticalPlaneX := 0
	chamberRightVerticalPlaneX := chamberLeftVerticalPlaneX + chamberWidth + 1
	chamberBottomHorizontalPlaneY := 0
	var rocks []Object2d
	jpIdx := 0
	lenJetPatterns := len(jetPatterns)
	towerHeight := 0

	for i := 0; i < numRocks; i++ {
		rockShape := rockShapes[i%5]
		rock := Object2d{MeshTemplate: rockShape, Position: util.Coordinate{X: chamberLeftVerticalPlaneX + spawnOffset.X + 1, Y: towerHeight + spawnOffset.Y + 1}}
		stopped := false
		step := 0
		for !stopped {
			// apply jet effect
			jetPattern := jetPatterns[jpIdx%lenJetPatterns]
			jetEffect := GetJetEffect(jetPattern)

			changeX := jetEffect
			minAbsX := rock.Position.X + changeX
			maxAbsX := rock.Position.X + rock.MeshTemplate.MaxBound.X + changeX
			if maxAbsX >= chamberRightVerticalPlaneX {
				changeX = 0
			}
			if minAbsX <= chamberLeftVerticalPlaneX {
				changeX = 0
			}

			if changeX != 0 {
				if step < spawnOffset.Y {
					rock.Position = util.Coordinate{X: rock.Position.X + changeX, Y: rock.Position.Y}
				} else {
					oldPos := rock.Position
					rock.Position = util.Coordinate{X: oldPos.X + changeX, Y: oldPos.Y}
					rock.PersistAbsoluteCoords()
					if rock.CollidesWithAnyObject(rocks) {
						rock.Position = oldPos
					}
				}
			}

			// apply gravity
			changeY := -1
			minAbsY := rock.Position.Y + changeY
			if minAbsY <= chamberBottomHorizontalPlaneY {
				changeY = 0
			}

			if changeY != 0 {
				if step < spawnOffset.Y {
					rock.Position = util.Coordinate{X: rock.Position.X, Y: rock.Position.Y + changeY}
				} else {
					oldPos := rock.Position
					rock.Position = util.Coordinate{X: oldPos.X, Y: oldPos.Y + changeY}
					rock.PersistAbsoluteCoords()
					if rock.CollidesWithAnyObject(rocks) {
						rock.Position = oldPos
						rock.PersistAbsoluteCoords()
						if rock.AbsMaxBound.Y > towerHeight {
							towerHeight = rock.AbsMaxBound.Y
						}
						stopped = true
					}
				}
			} else {
				rock.PersistAbsoluteCoords()
				if rock.AbsMaxBound.Y > towerHeight {
					towerHeight = rock.AbsMaxBound.Y
				}
				stopped = true
			}

			jpIdx++
			step++
		}
		rocks = append(rocks, rock)
	}
	return towerHeight
}

func main() {
	const numRocks = 2022

	lines := util.ReadFileLinesAsArray(inputFile)
	jetPatterns := lines[0]
	rockShapes := CreateRockShapes()
	towerHeight := simulateRocks(numRocks, rockShapes, ChamberWidth, SpawnOffset, jetPatterns)
	fmt.Println(towerHeight)
}
