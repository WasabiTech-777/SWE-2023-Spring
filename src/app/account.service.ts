import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import {environment} from '../environments/environment';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  constructor(private http: HttpClient, private router: Router) { }

  login(uname: any, pass: any) {
    var userInfo = {
      "uname":uname,
      "pass":pass,
      "articles":0,
      "charhit":0,
      "charmiss":0
    }
    var valInfo = this.http.post<HttpResponse<any>>(`${environment.serverUrl}/login`, userInfo, {observe: 'response'});
    console.log(userInfo);
    valInfo.subscribe(
      (data: HttpResponse<any>) => {
      if (data.status === 200) {
        this.router.navigate(['/profile']);
      }});
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
  validate() {
    
  }
}
