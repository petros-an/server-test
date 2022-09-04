import {State} from './state.mjs';
import {character} from './character.mjs';

const m = document.getElementById("life").getContext("2d");
const width = 800
const height = 800

function draw (x, y, c, w, h) {
  m.fillStyle = c;
  m.fillRect(x, y, w, h);
};

let socket = new WebSocket("ws://localhost:8080/echo");

const currentState = new State(m)
currentState.characters[0] = new character(100,100,m)
currentState.characters[1] = new character(400,400,m)

socket.onmessage = currentState.updateState
socket.onopen = (x) => {socket.send("aaa")}


const update = () => {
    m.clearRect(0, 0, width, height);
    draw(0, 0, "black", width, height);

    for (let i = 0; i < currentState.characters.length; i++) {
      currentState.characters[i].draw(m)
    }

    requestAnimationFrame(update);
};


update();
