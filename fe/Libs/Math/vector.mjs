export class Vector2D {

    constructor(x, y) {
        this.x = x
        this.y = y
    }

    toString() {
        return "(" + this.x + ", " + this.y + ")"
    }

    normalize() {
        if (this.x == 0 && this.y == 0) return
        let len = this.magnitude
        this.x /= len
        this.y /= len
    }

    get magnitude() {
        return Math.sqrt(this.x * this.x + this.y * this.y)
    }

    get magnitudeSq() {
        return this.x * this.x + this.y * this.y
    }

    get phase() {
        return Math.atan2(this.y, this.x)
    }

    add(arg1, arg2) {
        if (arg2 == undefined) {
            return new Vector2D(this.x + arg1.x, this.y + arg1.y)
        }
        else {
            return new Vector2D(this.x + arg1, this.y + arg2)
        }
    }

    addSelf(arg1, arg2) {
        if (arg2 == undefined) {
            this.x += arg1.x
            this.y += arg1.y
        }
        else {
            this.x += arg1
            this.y += arg2
        }
    }

    equals(arg1, arg2) {
        if (arg2 == undefined) {
            return this.x == arg1.x && this.y == arg1.y
        }
        else {
            return this.x == arg1 && this.y == arg2
        }
    }
}