export class character{
    
    constructor(x, y, m){
        this.position = {x: x, y: y}
        this.m = m
    }

    draw() {
        this.m.fillStyle = "red";
        this.m.fillRect(this.position.x, this.position.y, 10, 10);
    }

}