import { Character } from "./character.mjs";
import { Vector2D } from "./Libs/Math/vector.mjs";

export const m = document.getElementById("life").getContext("2d");

const width = 800
const height = 800
const widthToHeightRatio = width / height
const widthDiv2 = width / 2
const heightDiv2 = height / 2

const worldWidth = 400
const worldHeight = worldWidth / widthToHeightRatio
const pixelToWorldRatio = width / worldWidth
const worldOriginOnCanvasRatio = new Vector2D(0.5, 0.5)

export function render(currentState) {
    m.clearRect(0, 0, width, height);
    draw(0, 0, "black", width, height);

    for (let i = 0; i < currentState.characters.length; i++) {
        drawCharacter(currentState.characters[i])
    }
}

function draw(x, y, c, w, h) {
    m.fillStyle = c;
    m.fillRect(x, y, w, h);
};


function worldToCanvas(worldPosition) {
    return worldPosition.add(worldOriginOnCanvasRatio.mul(width, height).div(pixelToWorldRatio)).mul(pixelToWorldRatio)
}


function drawCharacter(character) {
    character.m.fillStyle = "red";
    console.log("world: " + character.position)
    let pos = worldToCanvas(character.position)
    console.log("canvas: " + pos)
    character.m.fillRect(pos.x, pos.y, 10, 10);
}