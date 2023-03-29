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
  }
  cookieValue: String | undefined;
  name: String | undefined;
  ngOnInit(){
    this.account.validate();
    this.cookieValue = document.cookie
    .split("; ")
    .find((row) => row.startsWith("token="))
    ?.split("=")[1];
    if(this.cookieValue != undefined)
      this.name = this.account.decodeToken(this.cookieValue)['uname'] as string
  }
}
