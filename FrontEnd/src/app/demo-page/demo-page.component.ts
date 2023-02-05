import { Component } from '@angular/core';
import { WikiPrompt } from '../wikiPrompt';

@Component({
  selector: 'app-demo-page',
  templateUrl: './demo-page.component.html',
  styleUrls: ['./demo-page.component.css']
})
export class DemoPageComponent {
  wiki: WikiPrompt = { 
    name: "Japan",
    text: "The first people in Japan were the Ainu people and other Jōmon people. They were closer related to Europeans or Mongols. They were later conquered and replaced by the Yayoi people (early Japanese and Ryukyuans). The Yayoi were an ancient ethnic group that migrated to the Japanese archipelago mainly from southeastern China during the Yayoi period (300 BCE–300 CE). Modern Japanese people have primarily Yayoi ancestry at an average of 97%. The indigenous Ryukyuan and Ainu peoples have more Jōmon ancestry on the other hand.",
  };


  ngOnInit(): void {
  }
}
