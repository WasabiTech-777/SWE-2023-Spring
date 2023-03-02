import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient) { }

  login(uname: any, pass: any) {
    var userInfo = {
      "uname":uname,
      "pass":pass,
      "articles":0,
      "charhit":0,
      "charmiss":0
    }
    var valInfo = this.http.post(`${environment.serverUrl}/login`, userInfo);
    console.log(userInfo);
    valInfo.subscribe();
  }
  createAccount(uname: any, pass: any) {
    var userInfo = {
      "uname":uname,
      "pass":pass,
      "articles":0,
      "charhit":0,
      "charmiss":0
    }
    var valInfo = this.http.post(`${environment.serverUrl}/users`, userInfo);
    console.log(userInfo);
    valInfo.subscribe();
    return valInfo;
  }
}
