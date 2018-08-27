package main

import "github.com/EngoEngine/glm"

type snake struct {
	bodyparts []glm.Vec2
	direction glm.Vec2
	speed float32
}

func (this *snake) MoveUp() {
	this.direction = glm.Vec2{0.0, -this.speed}
}

func (this *snake) MoveDown() {
	this.direction = glm.Vec2{0.0, this.speed}
}

func (this *snake) MoveLeft() {
	this.direction = glm.Vec2{-this.speed, 0.0}
}

func (this *snake) MoveRight() {
	this.direction = glm.Vec2{this.speed, 0.0}
}

func (this *snake) AddBodypart() {
	var totalBodyParts int = len(this.bodyparts)
	var lastBodypart = this.bodyparts[totalBodyParts-1]
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
	var totalBodyParts int = len(this.bodyparts)
	for i := 0; i < totalBodyParts-1; i++ {
		this.bodyparts[i] = this.bodyparts[i+1]
	}
	this.bodyparts[totalBodyParts-1].AddWith(&this.direction)
}

func NewSnake() *snake {
	var bodyparts = make([]glm.Vec2, 0)
	var direction = glm.Vec2{0.0, 0.0}
	var head = glm.Vec2{20.0, 20.0}
	bodyparts = append(bodyparts, head)

	return &snake{bodyparts, direction, 0.001}
}