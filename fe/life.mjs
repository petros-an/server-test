import { State } from './state.mjs';
import { PlayerController } from './player_controller.mjs';
import { Api } from './api.mjs';

const m = document.getElementById("life").getContext("2d");
const width = 800
const height = 800

function draw(x, y, c, w, h) {
    m.fillStyle = c;
    m.fillRect(x, y, w, h);
};

const socket = new WebSocket("ws://localhost:8080/state");
const api = new Api(socket)

let ID;
const currentState = new State(m, 'mainstate')
// currentState.characters[0] = new character(100,100,m)
// currentState.characters[1] = new character(400,400,m)

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

currentState.updateState
socket.onopen = (x) => { socket.send("aaa") }

const pressedKeys = {};
window.onkeyup = function (e) { pressedKeys[e.key] = false; }
window.onkeydown = function (e) { pressedKeys[e.key] = true; }
const playerController = new PlayerController(pressedKeys, api)

const update = () => {
    playerController.checkPlayerIntput()
    render()
    requestAnimationFrame(update);
};

function render() {
    m.clearRect(0, 0, width, height);
    draw(0, 0, "black", width, height);

    for (let i = 0; i < currentState.characters.length; i++) {
        currentState.characters[i].draw(m)
    }
}

update();
