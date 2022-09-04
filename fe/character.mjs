export class character{
    
    constructor(x, y){
        this.position = {x: x, y: y}
    }

    draw(m) {
        m.fillStyle = "red";
        m.fillRect(this.position.x, this.position.y, 10, 10);
    }

}