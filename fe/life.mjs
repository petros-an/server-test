import { State } from './state.mjs';
import { PlayerController } from './player_controller.mjs';
import { Api } from './api.mjs';
import { render, m } from './render.mjs';



const socket = new WebSocket("ws://localhost:8080/state");
const api = new Api(socket)

const currentState = new State(m, 'mainstate')

socket.onmessage = (event) => {
    let parsed = JSON.parse(event.data)
    if (!isNaN(parsed)) {
        ID = parsed
        //console.log("ID: ", ID)
    }
    else {
        currentState.updateState(parsed)
    }
}

const ID =  (Math.random() + 1).toString(16).substring(2);
console.log(ID)
socket.onopen = (x) => { socket.send(ID) }

const pressedKeys = {};
window.onkeyup = function (e) { pressedKeys[e.key] = false; }
window.onkeydown = function (e) { pressedKeys[e.key] = true; }
const playerController = new PlayerController(pressedKeys, api)

const update = () => {
    playerController.checkPlayerIntput()
    render(currentState)
    requestAnimationFrame(update);
};


update();
