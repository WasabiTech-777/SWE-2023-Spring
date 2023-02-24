import { Component, ViewEncapsulation } from '@angular/core';
import { WikiPrompt } from '../wikiPrompt';
import Keyboard from "simple-keyboard";

@Component({
  selector: 'app-demo-page',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './demo-page.component.html',
  styleUrls: ['./demo-page.component.css']
})
export class DemoPageComponent {
  wiki: WikiPrompt = { 
    name: "Japan",
    text: "The first people in Japan were the Ainu people and other Jōmon people. They were closer related to Europeans or Mongols. They were later conquered and replaced by the Yayoi people (early Japanese and Ryukyuans). The Yayoi were an ancient ethnic group that migrated to the Japanese archipelago mainly from southeastern China during the Yayoi period (300 BCE–300 CE). Modern Japanese people have primarily Yayoi ancestry at an average of 97%. The indigenous Ryukyuan and Ainu peoples have more Jōmon ancestry on the other hand.",
  };

  constructor() {}

  usrText: string = "";
  usrInpt: string = "";
  keyboard!: Keyboard;

  timeTaken: number = 0;
  correct: number = 0;
  errors: number = 0;

  WPM: number = 0;
  accuracy: number = 0;
  //timer;


  ngOnInit(): void {
  }

  ngAfterViewInit() {
    this.keyboard = new Keyboard({
      onChange: usrInpt => this.onChange(usrInpt),
      onKeyPress: button => this.onKeyPress(button),
      physicalKeyboardHighlight: true,
      physicalKeyboardHighlightTextColor: "blue",
    });
  }

  onChange = (value: string) => {
    this.usrInpt = value;
    console.log("usrInpt changed", value);
  };

  onKeyPress = (button: string) => {
    console.log("Button pressed", button);

    /**
     * If you want to handle the shift and caps lock buttons
     */
    if (button === "{shift}" || button === "{lock}") this.handleShift();
    
  };

  onInputChange = (event: any) => {
    this.keyboard.setInput(event.target.usrInpt);

  };


  handleShift = () => {
    let currentLayout = this.keyboard.options.layoutName;
    let shiftToggle = currentLayout === "default" ? "shift" : "default";

    this.keyboard.setOptions({
      layoutName: shiftToggle
    });
  };

  onUserInput(value:string){
    this.usrText = value
    if (this.usrText === this.wiki.text) {
      //this.winner = true
    }
  }

  compare(randomLetter:string, enteredLetter:string, pos: number){
    if (pos == 0) {
      this.correct = 0
      this.errors = 0
    }


    if (!enteredLetter) {
      return 'pending'
    }
    else if (randomLetter === enteredLetter) {
      this.correct += 1
      return 'correct'
    }
    else{
      this.errors += 1
      return 'incorrect'
    }
  }
}
