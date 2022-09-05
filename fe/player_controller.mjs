import { Vector2D } from "./Libs/Math/vector.mjs"

export class PlayerController {

    constructor(pressedKeys, api) {
        this.pressedKeys = pressedKeys
        this.api = api
    }

    checkPlayerIntput() {
        let direction = new Vector2D(0, 0)
        if (this.pressedKeys["w"]) {
            direction.y++
        }
        if (this.pressedKeys["s"]) {
            direction.y--
        }
        if (this.pressedKeys["a"]) {
            direction.x--
        }
        if (this.pressedKeys["d"]) {
            direction.x++
        }
        
        direction.normalize()

        if (!direction.equals(0, 0)){
            this.api.playerMoveInput(direction)
        }
    }
}
