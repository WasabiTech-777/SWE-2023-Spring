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
    this.article = -1
    this.WPM = -1
    this.acc = -1
  }
  cookieValue: String | undefined;
  name: String | undefined;
  article: number | undefined;
  WPM: number | undefined;
  acc: number | undefined;
  user: any;

  ngOnInit(){
    this.account.validate();
    this.cookieValue = document.cookie
    .split("; ")
    .find((row) => row.startsWith("token="))
    ?.split("=")[1];
    if(this.cookieValue != undefined) {
      this.name = this.account.decodeToken(this.cookieValue)['uname'] as string
      this.account.getUserInfo(this.name).subscribe(response => 
        {
          this.user = response;
          this.article = this.user.articles;
          if(this.user.charhit + this.user.charmiss > 0)
            this.acc = this.user.charhit/(this.user.charhit + this.user.charmiss);
          else
            this.acc = 0;
        })
    }
  }
}
