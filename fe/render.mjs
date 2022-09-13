import { Character } from "./character.mjs";
import { Vector2D } from "./Libs/Math/vector.mjs";

export const m = document.getElementById("life").getContext("2d");


const screenWidth = 800
const screenHeight = 800
const widthToHeightRatio = screenWidth / screenHeight
const widthDiv2 = screenWidth / 2
const heightDiv2 = screenHeight / 2

const cameraWidth = 80
const cameraHeight = 80
const cameraHalfWidth = cameraWidth / 2
const cameraHalfHeight = cameraHeight / 2
const cameraInverseHalfWidth = 1 / cameraHalfWidth
const cameraInverseHalfHeight = 1 / cameraHalfHeight

const pixelToWorldRatio = screenWidth / cameraWidth
const worldOriginOnCanvasRatio = new Vector2D(0.5, 0.5)

export function render(currentState) {
    m.clearRect(0, 0, screenWidth, screenHeight);
    drawWorld(m)

    for (let i = 0; i < currentState.characters.length; i++) {
        drawCharacter(currentState.characters[i])
    }
}

function drawWorld(m) {
    const img = document.getElementById('world')
    m.drawImage(
        img, 0, 0
    )
}

function draw(x, y, c, w, h) {
    m.fillStyle = c;
    m.fillRect(x, y, w, h);
};


function worldToCanvas(worldPosition) {
    let pixelPos = worldPosition.add(worldOriginOnCanvasRatio.mul(screenWidth, screenHeight).div(pixelToWorldRatio)).mul(pixelToWorldRatio)
    pixelPos.y = screenHeight - pixelPos.y
    return pixelPos
}


function drawCharacter(character) {
    character.m.fillStyle = `rgb(${character.color.R}, ${character.color.G}, ${character.color.B})`;
    let pos = worldToCanvas(character.position)
    character.m.fillRect(pos.x - 10 / 2, pos.y - 10 / 2, 10, 10);
    m.font = '20px serif';
    character.m.fillText(character.id, pos.x - 15, pos.y);
}