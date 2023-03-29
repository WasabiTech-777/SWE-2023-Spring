import { Component } from '@angular/core';
import { AccountService } from 'app/account.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent {
  constructor(private account: AccountService) { }
  ngOnInit(){
    this.account.validate();
    console.log(document.cookie)
  }
}
