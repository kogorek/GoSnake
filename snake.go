package main

import "github.com/EngoEngine/glm"
import "container/list"

type snake struct {
	bodyparts *list.List
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
	var lastBodypart = this.bodyparts.Back().Value.(glm.Vec2)
	var newBodypart = lastBodypart
	this.bodyparts.PushBack(newBodypart)
}

func (this *snake) checkSelfCollision() bool {
	for i := this.bodyparts.Front(); i != nil; i = i.Next() {
		for j := this.bodyparts.Front(); j != nil; j = j.Next() {
			if &i == &j{
				continue
			}
			first := i.Value.(glm.Vec2)
			second := j.Value.(glm.Vec2)

			if first.Equal(&second){
				return true
			}
			
		}
	}
	return false
}

func (this *snake) collisionWithPoint(point glm.Vec2) bool {
	var head = this.bodyparts.Front().Value.(glm.Vec2)
	if head.Equal(&point) {
		return true
	}
	return false
}

func (this *snake) Update() {
	var head = this.bodyparts.Front().Value.(glm.Vec2)
	head.AddWith(&this.direction)
	this.bodyparts.Front().Value = head

	for i := this.bodyparts.Back(); i != this.bodyparts.Front(); i = i.Prev() {
		// var currentPart = i.Value.(glm.Vec2)
		var nextPart = i.Prev().Value.(glm.Vec2)
		i.Value = nextPart
	}
}

func NewSnake() *snake {
	var bodyparts = list.New()
	var direction = glm.Vec2{0.0, 0.0}
	var head = glm.Vec2{20.0, 20.0}
	bodyparts.PushBack(head)

	return &snake{bodyparts, direction, 0.001}
}