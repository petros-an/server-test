import { Vector2D } from "./Libs/Math/vector.mjs";

export class Character {

    constructor(x, y, m) {
        this.position = new Vector2D(x, y)
        this.m = m
    }

}