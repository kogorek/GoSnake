package main

import "github.com/EngoEngine/glm"
import "container/list"

type snake struct {
	bodyparts *list.List
	direction glm.Vec2
}

func (this *snake) MoveUp() {
	this.direction = glm.Vec2{0.0, 1.0}
}

func (this *snake) MoveDown() {
	this.direction = glm.Vec2{0.0, -1.0}
}

func (this *snake) MoveLeft() {
	this.direction = glm.Vec2{-1.0, 0.0}
}

func (this *snake) MoveRight() {
	this.direction = glm.Vec2{1.0, 0.0}
}

func (this *snake) AddBodypart() {
	var lastBodypart = this.bodyparts.Back()
	var newBodypart = lastBodypart
	this.bodyparts.PushBack(newBodypart)
}

func NewSnake() snake {
	var bodyparts = list.New()
	var direction = glm.Vec2{0.0, 0.0}
	var head = glm.Vec2{10.0, 10.0}
	bodyparts.PushBack(head)

	return snake{bodyparts, direction}
}