import {character} from './character.mjs';
import { convertStateStream } from './reciever.mjs';

export class State{

    constructor(m){
        this.m = m;
        this.characters = [];
    }

    updateState(event){
        console.log(event)
        console.log(event.data)
//        convertStateStream(event.data)
    }

}