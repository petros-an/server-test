export class PlayerController{

    constructor(pressedKeys){
        this.pressedKeys = pressedKeys
    }

    checkPlayerIntput(){
        if(this.pressedKeys["w"]){
            console.log("up")
        }
        if(this.pressedKeys["s"]){
            console.log("down")
        }
        if(this.pressedKeys["a"]){
            console.log("left")
        }
        if(this.pressedKeys["d"]){
            console.log("right")
        }
    }
}
