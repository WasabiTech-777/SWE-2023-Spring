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

  value = "";
  keyboard!: Keyboard;

  ngOnInit(): void {
  }

  ngAfterViewInit() {
    this.keyboard = new Keyboard({
      onChange: input => this.onChange(input),
      onKeyPress: button => this.onKeyPress(button),
      physicalKeyboardHighlight: true,
      physicalKeyboardHighlightTextColor: "blue",
    });
  }

  onChange = (input: string) => {
    this.value = input;
    console.log("Input changed", input);
  };

  onKeyPress = (button: string) => {
    console.log("Button pressed", button);

    /**
     * If you want to handle the shift and caps lock buttons
     */
    if (button === "{shift}" || button === "{lock}") this.handleShift();
  };

  onInputChange = (event: any) => {
    this.keyboard.setInput(event.target.value);
  };

  handleShift = () => {
    let currentLayout = this.keyboard.options.layoutName;
    let shiftToggle = currentLayout === "default" ? "shift" : "default";

    this.keyboard.setOptions({
      layoutName: shiftToggle
    });
  };
}
