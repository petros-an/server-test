//import {State} from './state.js';


m = document.getElementById("life").getContext("2d");
const width = 800
const height = 800

const draw = (x, y, c, w, h) => {
  m.fillStyle = c;
  m.fillRect(x, y, w, h);
};

const currentState = new State()
//currentState.characters[0] = new character(100,100)
//currentState.characters[1] = new character(400,400)

const update = () => {
    m.clearRect(0, 0, width, height);
    draw(0, 0, "black", width, height);
    requestAnimationFrame(update);
};


update();