import {Component, OnInit} from '@angular/core';
import {HelloWorldService} from './hello-world.service';
import { AccountService } from 'app/account.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  title: any;

  //constructor(private hw: HelloWorldService) {}

  constructor(private account: AccountService, private hw: HelloWorldService) { 
    this.name = "Guest"
  }
  cookieValue: String | undefined;
  name: String | undefined;
  ngOnInit(){
    this.hw.getTitle().subscribe(response => {this.title = (response as any).title});
    this.account.validate();
    this.cookieValue = document.cookie
    .split("; ")
    .find((row) => row.startsWith("token="))
    ?.split("=")[1];
    if(this.cookieValue != undefined)
      this.name = this.account.decodeToken(this.cookieValue)['uname'] as string
  }
  //ngOnInit() {
    //this.hw.getTitle().subscribe(response => {this.title = (response as any).title});
  //}

}
