import { Component } from '@angular/core';
import { AccountService } from 'app/account.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent {
  constructor(private account: AccountService) { 
    this.name = "Guest"
    this.article = 0
    this.WPM = 0
    this.acc = 0
  }
  cookieValue: String | undefined;
  name: String | undefined;
  article: number | undefined;
  WPM: number | undefined;
  acc: number | undefined;

  ngOnInit(){
    this.account.validate();
    this.cookieValue = document.cookie
    .split("; ")
    .find((row) => row.startsWith("token="))
    ?.split("=")[1];
    if(this.cookieValue != undefined) {
      this.name = this.account.decodeToken(this.cookieValue)['uname'] as string
      //this.article = this.account.decodeToken(this.cookieValue)['articles'] as number
      //this.WPM = this.account.decodeToken(this.cookieValue)['wpm'] as number
      //this.acc = (this.account.decodeToken(this.cookieValue)['charhit'] as number) / ((this.account.decodeToken(this.cookieValue)['charhit'] as number) + (this.account.decodeToken(this.cookieValue)['charmiss'] as number))
    }
  }
}
