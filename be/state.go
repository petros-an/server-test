package main

import "time"

type GameState struct {
	Characters []*Character
}

type StateInput struct {
	NewCharacter   *Character
	VelocityUpdate *VelocityUpdate
}

type VelocityUpdate struct {
	CharacterId CharacterId
	Vx          int
	Vy          int
}

func gameStateMaintainer(
	output chan GameState,
	input chan StateInput,
	stopper chan bool,
) {
	gameState := GameState{
		Characters: []*Character{},
	}

	outputTicker := time.NewTicker(1 * time.Second)
	gameLoopTicker := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-outputTicker.C:
			//log.Println("[M] Sending new state")
			//log.Println(gameState.Characters)
			output <- gameState
		case stateInput := <-input:
			//log.Println("[M] Received state update:")
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
			char.vx += velUpdate.Vx
			char.vy += velUpdate.Vy
			break
		}
	}
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
