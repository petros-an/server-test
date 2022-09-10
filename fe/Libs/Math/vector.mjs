export var Vector2D = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x = arg1.x
        this.y = arg1.y
    }
    else {
        this.x = arg1
        this.y = arg2
    }
}

Object.defineProperties(Vector2D.prototype, {
    magnitude: {
        get: function () {
            return Math.sqrt(this.x * this.x + this.y * this.y)
        }
    },
    magnitudeSq: {
        get: function () {
            return this.x * this.x + this.y * this.y
        }
    },
    angleR: {
        get: function () {
            return Math.atan2(this.y, this.x)
        }
    },
    angleD: {
        get: function () {
            return Math.atan2(this.y, this.x) * 180 / Math.PI
        }
    },
    normalized: {
        get: function () {
            return new Vector2D(this) / this.magnitude
        }
    },
    conj: {
        get: function () {
            return new Vector2D(this.x, -this.y)
        }
    },
})

Vector2D.prototype.toString = function () {
    return "(" + this.x + ", " + this.y + ")"
}

Vector2D.prototype.equals = function (arg1, arg2) {
    if (arg2 == undefined) {
        return this.x == arg1.x && this.y == arg1.y
    }
    else {
        return this.x == arg1 && this.y == arg2
    }
}

Vector2D.prototype.normalize = function () {
    if (this.x == 0 && this.y == 0) return
    let len = this.magnitude
    this.x /= len
    this.y /= len
    return this
}

Vector2D.prototype.conjSelf = function () {
    this.y = -this.y
    return this
}

Vector2D.prototype.add = function (arg1, arg2) {
    return new Vector2D(this).addSelf(arg1, arg2)
}

Vector2D.prototype.addSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x += arg1.x
        this.y += arg1.y
    }
    else {
        this.x += arg1
        this.y += arg2
    }
    return this
}

Vector2D.prototype.sub = function (arg1, arg2) {
    return new Vector2D(this).subSelf(arg1, arg2)
}

Vector2D.prototype.subSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x -= arg1.x
        this.y -= arg1.y
    }
    else {
        this.x -= arg1
        this.y -= arg2
    }
    return this
}

Vector2D.prototype.opp = function () {
    return new Vector2D(-this.x, -this.y)
}

Vector2D.prototype.oppSelf = function () {
    this.x, this.y = -this.x, -this.y
    return this
}

Vector2D.prototype.mul = function (arg1, arg2) {
    return new Vector2D(this).mulSelf(arg1, arg2)
}

Vector2D.prototype.mulSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        if (typeof arg1 == "Vector2D") {
            this.x *= arg1.x
            this.y *= arg1.y
        }
        else {
            this.x *= arg1
            this.y *= arg1
        }
    }
    else {
        this.x *= arg1
        this.y *= arg2
    }
    return this
}

Vector2D.prototype.div = function (arg1, arg2) {
    return new Vector2D(this).divSelf(arg1, arg2)
}

Vector2D.prototype.divSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        if (typeof arg1 == "Vector2D") {
            this.x /= arg1.x
            this.y /= arg1.y
        }
        else {
            this.x /= arg1
            this.y /= arg1
        }
    }
    else {
        this.x /= arg1
        this.y /= arg2
    }
    return this
}

Vector2D.prototype.mulC = function (arg1, arg2) {
    return new Vector2D(this).mulCSelf(arg1, arg2)
}

Vector2D.prototype.mulCSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x, this.y = this.x * arg1.x - this.y * arg1.y, this.x * arg1.y + this.y * arg1.x
    }
    else {
        this.x, this.y = this.x * arg1 - this.y * arg2, this.x * arg2 + this.y * arg1
    }
    return this
}

Vector2D.prototype.mulConj = function (arg1, arg2) {
    return new Vector2D(this).mulConjSelf(arg1, arg2)
}

Vector2D.prototype.mulConjSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x, this.y = this.x * arg1.x + this.y * arg1.y, this.x * arg1.y - this.y * arg1.x
    }
    else {
        this.x, this.y = this.x * arg1 + this.y * arg2, this.x * arg2 - this.y * arg1
    }
    return this
}

Vector2D.prototype.divC = function (arg1, arg2) {
    return new Vector2D(this).divCSelf(arg1, arg2)
}

Vector2D.prototype.divCSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x, this.y = this.x * arg1.x + this.y * arg1.y, this.x * arg1.y - this.y * arg1.x
        this.div(arg1.magnitudeSq)
    }
    else {
        this.x, this.y = this.x * arg1 + this.y * arg2, this.x * arg2 - this.y * arg1
        this.div(arg1 * arg1 + arg2 * arg2)
    }
    return this
}

Vector2D.prototype.rotate90 = function () {
    return new Vector2D(-this.y, this.x)
}

Vector2D.prototype.rotate90Self = function () {
    this.x, this.y = -this.y, this.x
    return this
}