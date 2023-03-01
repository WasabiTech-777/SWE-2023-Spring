import { Component } from '@angular/core';
import { AccountService } from 'app/account.service';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent {
  constructor(private account: AccountService) { }
  hide = true;
  public user = new FormControl();
  public pass = new FormControl();
  onSubmit() {
    this.account.createAccount(this.user.value, this.pass.value);
  }
}
