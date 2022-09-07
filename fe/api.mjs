export class Api {

    constructor(socket) {
        this.socket = socket;
    }


    playerMoveInput(direction) {
        let message = JSON.stringify({
            "Velocity": {
                "VX": direction.x, "VY": direction.y
            }
        })
        this.message = message
    }


}

