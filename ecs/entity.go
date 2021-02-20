// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecs

import (
	"sync/atomic"
)

type EntityId uint32

var eId uint32

func getEntityId() EntityId {
	return EntityId(atomic.AddUint32(&eId, 1))
}

type Entity interface {
	Container
	EntityId() EntityId
}

type Container interface {
	AddComponent(tag ComponentTag, component interface{}) Container
	HasComponent(tag ComponentTag) bool
	// geef component terug die voldoet aan complete tag bitmask
	Component(tag ComponentTag) interface{}
	// geef components terug die aan een of meerdere tags voldoen, 1 component per tag
	Components(tag ...ComponentTag) map[ComponentTag]interface{}
}

type entity struct {
	manager    *Manager
	entityId   EntityId
	components []ComponentTag
}

func (e *entity) EntityId() EntityId { return e.entityId }

func (e *entity) AddComponent(tag ComponentTag, component interface{}) Container {
	e.manager.addComponent(e.entityId, tag, component)
	e.components = append(e.components, tag)
	return Container(e)
}

func (e *entity) HasComponent(tag ComponentTag) bool {
	for _, ct := range e.components {
		if ct == tag {
			return true
		}
	}
	return false
}

func (e *entity) Component(tag ComponentTag) interface{} {
	return e.manager.getComponent(e.entityId, tag)
}

// todo: get components from manager
func (e *entity) Components(tag ...ComponentTag) map[ComponentTag]interface{} {
	result := make(map[ComponentTag]interface{})
	return result
}
