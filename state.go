package main

import (
	"fmt"
	"log"
	"time"
)

const EVICTION_INTERVAL = 20 * time.Second

type GameState struct {
	Characters []*Character
	Players    map[PlayerId]*Player
}

func (s GameState) Repr() string {
	str := ""
	for _, c := range s.Characters {
		str += fmt.Sprintf("[x:%f, y:%f, vx: %f, vy: %f, id: %s],", c.RigidBody.LocalPosition.X, c.RigidBody.LocalPosition.Y, c.RigidBody.Velocity.X, c.RigidBody.Velocity.Y, c.Tag)
	}
	return str
}

type InputMessage struct {
	Type      MessageType
	PlayerId  PlayerId
	NewPlayer *Player
	Direction Vector2D
}

type MessageType byte

const (
	O_STATE MessageType = iota
)

const (
	I_NEW MessageType = iota
	I_DISCONNECT
	I_DIRECTION
)

type OutputMessage struct {
	Type         MessageType
	CurrentState GameState
}

const FPS = 50.0
const DT = 1 / FPS

const sendTickerSeconds = 0.01

const worldBorderX1 = -40
const worldBorderX2 = 40
const worldBorderY1 = -40
const worldBorderY2 = 40

func gameStateMaintainer(
	outputChannel chan OutputMessage,
	inputChannel chan InputMessage,
	stopper chan bool,
) {
	gameState := GameState{
		Characters: []*Character{},
		Players:    make(map[PlayerId]*Player),
	}

	outputTicker := time.NewTicker(time.Duration(int64(sendTickerSeconds*1000)) * time.Millisecond)
	gameLoopTicker := time.NewTicker(time.Duration(int64(DT*1000)) * time.Millisecond)
	// go evictor(inputChannel)

	for {
		select {
		case <-outputTicker.C:
			outputChannel <- OutputMessage{Type: O_STATE, CurrentState: gameState}

		case stateInput := <-inputChannel:
			applyStateUpdate(&gameState, stateInput)
			applyVitalsUpdate(&gameState, stateInput.PlayerId)

		case <-gameLoopTicker.C:
			applyGameLoopUpdate(&gameState)
		}
	}
}

func evictor(inputChannel chan InputMessage) {
	tck := time.NewTicker(time.Second)
	for range tck.C {
		inputChannel <- InputMessage{Type: I_DISCONNECT}
	}

}

func applyStateUpdate(currentState *GameState, input InputMessage) {

	switch input.Type {
	case I_DISCONNECT:
		applyRemoveInactivePlayersUpdate(currentState)
	case I_DIRECTION:
		applyVelocityUpdate(currentState, input.Direction, input.PlayerId)
	case I_NEW:
		applyNewPlayerUpdate(currentState, input.NewPlayer, input.PlayerId)
	}
}

func applyNewPlayerUpdate(state *GameState, newPlayer *Player, id PlayerId) {
	if _, exists := state.Players[id]; exists {
		return
	}
	state.Players[id] = newPlayer
	state.Characters = append(state.Characters, newPlayer.Character)
}

func applyVelocityUpdate(state *GameState, Direction Vector2D, id PlayerId) {
	state.Players[id].Character.MoveDirection = Direction
}

func applyVitalsUpdate(state *GameState, id PlayerId) {
	if _, exists := state.Players[id]; !exists {
		return
	}
	state.Players[id].LastVital = time.Now()
}

func applyGameLoopUpdate(state *GameState) {
	for _, c := range state.Characters {
		forEachCharacter(c)
	}
}

func applyRemoveInactivePlayersUpdate(state *GameState) {
	for id, p := range state.Players {
		if time.Now().Sub(p.LastVital) > EVICTION_INTERVAL {
			removeElementFromSlice(&state.Characters, p.Character)
			delete(state.Players, id)
			log.Println("kicking due to connection timeout: " + id)
		}
	}
}

func forEachCharacter(c *Character) {
	c.Update()
	worldBorder(c)
}

func worldBorder(c *Character) {
	pos := c.Position()
	oldPos := pos
	if pos.X < worldBorderX1 {
		pos.X = worldBorderX1
	}
	if pos.X > worldBorderX2 {
		pos.X = worldBorderX2
	}
	if pos.Y < worldBorderY1 {
		pos.Y = worldBorderY1
	}
	if pos.Y > worldBorderY2 {
		pos.Y = worldBorderY2
	}
	if !oldPos.Equals(pos) {
		c.SetPosition(pos)
	}
}
