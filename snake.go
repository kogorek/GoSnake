package main

import "github.com/EngoEngine/glm"

type snake struct {
	bodyparts []glm.Vec2
	direction glm.Vec2
	speed float32
	moveTimer float32
}

func (this *snake) MoveUp() {
	this.direction = glm.Vec2{0.0, -1.0}
}

func (this *snake) MoveDown() {
	this.direction = glm.Vec2{0.0, 1.0}
}

func (this *snake) MoveLeft() {
	this.direction = glm.Vec2{-1.0, 0.0}
}

func (this *snake) MoveRight() {
	this.direction = glm.Vec2{1.0, 0.0}
}

func (this *snake) AddBodypart() {
	var lastBodypart = this.bodyparts[0]
	var newBodypart glm.Vec2 = lastBodypart
	this.bodyparts = append(this.bodyparts, newBodypart)
}

func (this *snake) checkSelfCollision() bool {
	for i, first := range this.bodyparts {
		for j, second := range this.bodyparts {
			if(i == j){
				continue
			}
			if first.EqualThreshold(&second, 0.7) {
				return true
			}
		}
	}
	return false
}

func (this *snake) collisionWithPoint(point glm.Vec2) bool {
	var totalBodyParts int = len(this.bodyparts)
	var head = this.bodyparts[totalBodyParts-1]
	if head.EqualThreshold(&point, 0.7) {
		return true
	}
	return false
}

func (this *snake) Update() {

	if this.moveTimer >= 1 {
		var totalBodyParts int = len(this.bodyparts)
		this.bodyparts[totalBodyParts-1].AddWith(&this.direction)
		this.moveTimer = 0.0

		for i := 0; i < totalBodyParts-1; i++ {
			this.bodyparts[i] = this.bodyparts[i+1]
		}
	}

	this.moveTimer += this.speed
}

func NewSnake(position glm.Vec2) *snake {
	var bodyparts = make([]glm.Vec2, 0)
	var direction = glm.Vec2{0.0, 0.0}
	var head = position
	bodyparts = append(bodyparts, head)

	return &snake{bodyparts, direction, 0.001, 0.0}
}