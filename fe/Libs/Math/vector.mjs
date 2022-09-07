export var Vector2D = function (x, y) {
    this.x = x || 0;
    this.y = y || 0;
};

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
    phase: {
        get: function () {
            return Math.atan2(this.y, this.x)
        }
    },
});

Vector2D.prototype.toString = function () {
    return "(" + this.x + ", " + this.y + ")"
};

Vector2D.prototype.normalize = function () {
    if (this.x == 0 && this.y == 0) return
    let len = this.magnitude
    this.x /= len
    this.y /= len
    return this
};

Vector2D.prototype.add = function (arg1, arg2) {
    if (arg2 == undefined) {
        return new Vector2D(this.x + arg1.x, this.y + arg1.y)
    }
    else {
        return new Vector2D(this.x + arg1, this.y + arg2)
    }
}

Vector2D.prototype.addSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x += arg1.x
        this.y += arg1.y
        return this
    }
    else {
        this.x += arg1
        this.y += arg2
        return this
    }
}

Vector2D.prototype.sub = function (arg1, arg2) {
    if (arg2 == undefined) {
        return new Vector2D(this.x - arg1.x, this.y - arg1.y)
    }
    else {
        return new Vector2D(this.x - arg1, this.y - arg2)
    }
}

Vector2D.prototype.subSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x -= arg1.x
        this.y -= arg1.y
        return this
    }
    else {
        this.x -= arg1
        this.y -= arg2
        return this
    }
}

Vector2D.prototype.mul = function (arg1, arg2) {
    if (arg2 == undefined) {
        return new Vector2D(this.x * arg1, this.y * arg1)
    }
    else {
        return new Vector2D(this.x * arg1, this.y * arg2)
    }
}

Vector2D.prototype.mulSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x *= arg1
        this.y *= arg1
        return this
    }
    else {
        this.x *= arg1
        this.y *= arg2
        return this
    }
}

Vector2D.prototype.div = function (arg1, arg2) {
    if (arg2 == undefined) {
        return new Vector2D(this.x / arg1, this.y / arg1)
    }
    else {
        return new Vector2D(this.x / arg1, this.y / arg2)
    }
}

Vector2D.prototype.divSelf = function (arg1, arg2) {
    if (arg2 == undefined) {
        this.x /= arg1
        this.y /= arg1
        return this
    }
    else {
        this.x /= arg1
        this.y /= arg2
        return this
    }
}

Vector2D.prototype.equals = function (arg1, arg2) {
    if (arg2 == undefined) {
        return this.x == arg1.x && this.y == arg1.y
    }
    else {
        return this.x == arg1 && this.y == arg2
    }
}