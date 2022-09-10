import { Vector2D } from "./Libs/Math/vector.mjs";

export class Character {

    constructor(x, y, color, id, m) {
        this.position = new Vector2D(x, y)
        this.color = color
        this.id = id
        this.m = m
    }

}