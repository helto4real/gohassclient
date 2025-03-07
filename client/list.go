package client

import (
    "sync"
)

// List stores all entites and its states in memeory
//
// It support threadsafety by handling all writes through
// A go routine
type List struct {
    entities map[string]HassEntity
    m        sync.Mutex
}

// NewEntityList makes a new instance of entity list
func NewEntityList() List {
    return List{entities: make(map[string]HassEntity)}
}

// GetEntity returns entity given the entity id, second return value returns false if no entity exists
func (a *List) GetEntity(entityID string) (*HassEntity, bool) {
    a.m.Lock()
    defer a.m.Unlock()
    entity, ok := a.entities[entityID]
    return &entity, ok
}

// SetEntity sets the entity to the map
func (a *List) SetEntity(entity *HassEntity) {
    a.m.Lock()
    defer a.m.Unlock()
    a.entities[entity.ID] = *entity
}

// ByID sorting by the id
type ByID []HassEntity

func (e ByID) Len() int           { return len(e) }
func (e ByID) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e ByID) Less(i, j int) bool { return e[i].ID < e[j].ID }
