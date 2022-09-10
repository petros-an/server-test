import { Character } from './character.mjs';

export class State {

    constructor(m, id) {
        this.m = m;
        this.characters = [];
        this.id = id
    }



    updateState(parsed) {
        // console.log(event)
        console.log(parsed)
        this.characters = []
        for (let i = 0; i < parsed['Characters'].length; i++) {
            let characterData = parsed['Characters'][i]
            this.characters[i] = new Character(
                characterData['RigidBody']['LocalPosition']['x'],
                characterData['RigidBody']['LocalPosition']['y'],
                characterData['Color'],
                characterData['Id'],
                this.m
            )
        }
    }

}