import { Vector2D } from "./Libs/Math/vector.mjs";

export class character {

    constructor(x, y, m) {
        this.position = new Vector2D(x, y)
        this.m = m
    }

    draw() {
        this.m.fillStyle = "red";
        this.m.fillRect(this.position.x, this.position.y, 10, 10);
    }

}