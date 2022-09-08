package main

import (
	"fmt"
	"time"
)

type GameState struct {
	Characters []*Character
}

func (s GameState) Repr() string {
	str := ""
	for _, c := range s.Characters {
		str += fmt.Sprintf("[x:%f, y:%f, vx: %f, vy: %f, id: %s],", c.Position.X, c.Position.Y, c.Velocity.X, c.Velocity.Y, c.Id)
	}
	return str
}

type StateInput struct {
	NewCharacter   *Character
	VelocityUpdate *VelocityUpdate
}

type VelocityUpdate struct {
	CharacterId CharacterId
	Velocity    Vector2D
}

const FPS = 50.0
const DT = 1 / FPS

const sendTickerSeconds = 0.01

func gameStateMaintainer(
	output chan GameState,
	input chan StateInput,
	stopper chan bool,
) {
	gameState := GameState{
		Characters: []*Character{},
	}

	outputTicker := time.NewTicker(time.Duration(int64(sendTickerSeconds*1000)) * time.Millisecond)
	gameLoopTicker := time.NewTicker(time.Duration(int64(DT*1000)) * time.Millisecond)

	for {
		select {
		case <-outputTicker.C:
			//log.Println("[M] Sending new state")
			//log.Println(gameState.Characters)
			output <- gameState
		case stateInput := <-input:
			// log.Printf("[M] Received state update ")
			// log.Println(stateInput)
			gameState = applyStateUpdate(gameState, stateInput)
			//log.Println(gameState)
		case <-gameLoopTicker.C:
			gameState = applyGameLoopUpdate(gameState)
		}
	}
}

func applyNewCharacterUpdate(oldState GameState, newCharacter Character) GameState {
	for _, char := range oldState.Characters {
		if char.Id == newCharacter.Id {
			return oldState
		}
	}
	oldState.Characters = append(oldState.Characters, &newCharacter)
	return oldState
}

func applyVelocityUpdate(oldState GameState, velUpdate VelocityUpdate) GameState {
	for _, char := range oldState.Characters {
		if char.Id == velUpdate.CharacterId {
			char.Velocity = velUpdate.Velocity
			break
		}
	}

	//fmt.Println(oldState.Repr())
	return oldState
}

func applyStateUpdate(oldState GameState, input StateInput) GameState {

	state := oldState

	if input.NewCharacter != nil {
		state = applyNewCharacterUpdate(oldState, *input.NewCharacter)
	}

	if input.VelocityUpdate != nil {
		state = applyVelocityUpdate(oldState, *input.VelocityUpdate)
	}

	return state

}

func applyGameLoopUpdate(state GameState) GameState {
	for _, c := range state.Characters {
		c.move()
	}
	return state
}
