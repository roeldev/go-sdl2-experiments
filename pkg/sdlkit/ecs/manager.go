// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecs

import (
	"sync"
)

type Manager struct {
	mutex      sync.RWMutex
	entities   map[EntityId]Entity
	components map[ComponentTag]*components
}

type components struct {
	sync.RWMutex
	list map[EntityId]interface{}
}

func NewManager() *Manager {
	return &Manager{
		entities:   make(map[EntityId]Entity, 64),
		components: make(map[ComponentTag]*components, 64),
	}
}

func (em *Manager) Create(dest *Entity) Entity {
	entity := &entity{
		manager:  em,
		entityId: getEntityId(),
	}

	em.mutex.Lock()
	em.entities[entity.entityId] = entity
	em.mutex.Unlock()

	if dest != nil {
		*dest = entity
	}
	return entity
}

func (em *Manager) Register(container Container) Entity {
	entity := &containerEntity{
		Container: container,
		entityId:  getEntityId(),
		manager:   em,
	}

	em.mutex.Lock()
	em.entities[entity.entityId] = entity
	em.mutex.Unlock()

	return entity
}

func (em *Manager) Entity(id EntityId) Entity {
	em.mutex.RLock()
	res := em.entities[id]
	em.mutex.RUnlock()
	return res
}

// Entities returns all entities which have components with the given tags.
func (em *Manager) Entities(tag ...ComponentTag) []Entity {
	res := make([]Entity, 0, len(em.entities))

	em.mutex.RLock()
	for _, entity := range em.entities {
		var skip bool
		for _, ct := range tag {
			if !entity.HasComponent(ct) {
				skip = true
				break
			}
		}
		if !skip {
			res = append(res, entity)
		}
	}
	em.mutex.RUnlock()
	return res
}

func (em *Manager) Component(id EntityId, tag ComponentTag) interface{} {
	entity := em.Entity(id)
	if entity == nil {
		return nil
	}

	return entity.Component(tag)
}

func (em *Manager) Components(tag ComponentTag) []interface{} {
	em.mutex.RLock()
	cc, ok := em.components[tag]
	em.mutex.RUnlock()

	if !ok {
		return nil
	}

	cc.RLock()
	res := make([]interface{}, 0, len(cc.list))
	for _, comp := range cc.list {
		res = append(res, comp)
	}
	cc.RUnlock()
	return res
}

func (em *Manager) addComponent(id EntityId, tag ComponentTag, comp interface{}) {
	tags := make(map[ComponentTag]*components)
	em.mutex.RLock()
	for _, ct := range tag.Flags() {
		tags[ct] = em.components[ct]
	}
	em.mutex.RUnlock()

	newTags := make([]ComponentTag, 0, len(tags))
	for ct, cc := range tags {
		if nil == cc {
			newTags = append(newTags, ct)
			tags[ct] = &components{
				list: map[EntityId]interface{}{id: comp},
			}
		} else {
			cc.Lock()
			cc.list[id] = comp
			cc.Unlock()
		}
	}

	em.mutex.Lock()
	for _, ct := range newTags {
		em.components[ct] = tags[ct]
	}
	em.mutex.Unlock()
}

func (em *Manager) getComponent(id EntityId, tag ComponentTag) interface{} {
	em.mutex.RLock()
	cc, ok := em.components[tag]
	em.mutex.RUnlock()

	if !ok {
		return nil
	}

	cc.RLock()
	res := cc.list[id]
	cc.RUnlock()

	return res
}

type containerEntity struct {
	Container
	entityId EntityId
	manager  *Manager
}

func (e *containerEntity) EntityId() EntityId { return e.entityId }
