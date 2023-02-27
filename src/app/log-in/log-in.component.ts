import { Component } from '@angular/core';
import { AccountService } from 'app/account.service';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.css']
})
export class LogInComponent {
  constructor(private account: AccountService) { }
  hide = true;
  public user = new FormControl();
  public pass = new FormControl();
  onSubmit() {
    this.account.login(this.user.value, this.pass.value);
  }
}
