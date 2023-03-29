import {Component, OnInit} from '@angular/core';
import {HelloWorldService} from './hello-world.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  title: any;

  constructor(private hw: HelloWorldService) {}

  ngOnInit() {
    this.hw.getTitle().subscribe(response => {this.title = (response as any).title});
  }

}
