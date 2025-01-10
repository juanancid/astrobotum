package ecs

import "reflect"

type World struct {
	nextEntityID Entity
	components   map[reflect.Type]map[Entity]Component
	systems      []System
}

type Entity int

type Component interface{}

type System interface {
	Update(w *World, dt float64)
}

func NewWorld() *World {
	return &World{
		nextEntityID: 0,
		components:   make(map[reflect.Type]map[Entity]Component),
		systems:      make([]System, 0),
	}
}

func (w *World) AddEntity() Entity {
	id := w.nextEntityID
	w.nextEntityID++
	return id
}

func (w *World) AddComponent(entity Entity, component Component) {
	componentType := reflect.TypeOf(component)
	if w.components[componentType] == nil {
		w.components[componentType] = make(map[Entity]Component)
	}
	w.components[componentType][entity] = component
}

func (w *World) GetComponent(entity Entity, componentType reflect.Type) Component {
	return w.components[componentType][entity]
}

func (w *World) GetComponents(componentType reflect.Type) map[Entity]Component {
	return w.components[componentType]
}

func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) UpdateSystems(dt float64) {
	for _, s := range w.systems {
		s.Update(w, dt)
	}
}

func (w *World) GetSystems() []System {
	return w.systems
}
