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

const ID = (Math.random() + 1).toString(16).substring(2);
console.log(ID)

const sendTicker = 0.03
socket.onopen = (x) => {
    socket.send(ID)
    sendRepeat()
}

function sleep(time) {
    return new Promise((resolve) => setTimeout(resolve, time));
}

const sendRepeat = () => {
    while (true) {
        if (api.message != undefined) {
            socket.send(api.message)
            sleep(sendTicker * 1000).then(sendRepeat)
            return;
        }
    }
}

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
