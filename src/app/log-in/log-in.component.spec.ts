import { HttpClient } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MaterialModule } from 'app/material.module';
import { ReactiveFormsModule } from '@angular/forms';

import { LogInComponent } from './log-in.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AccountService } from 'app/account.service';
import { environment } from 'environments/environment';
import { RouterTestingModule } from '@angular/router/testing';
import { ProfilePageComponent } from 'app/profile-page/profile-page.component';

describe('LogInComponent', () => {
  let component: LogInComponent;
  let fixture: ComponentFixture<LogInComponent>;
  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;
  let service: AccountService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LogInComponent ],
      imports: [HttpClientTestingModule, MaterialModule, ReactiveFormsModule, BrowserAnimationsModule, 
        RouterTestingModule.withRoutes(
        [{path: 'profile', component: ProfilePageComponent}]
      )]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LogInComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
    service = TestBed.inject(AccountService)
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call login and allow a valid user to login', () => {
    let mockUsers = {uname: "hello10", pass: "world10"};
    service.login("hello10", "world10");
    const req = httpTestingController.match({
        method: 'POST',
        url: `${environment.serverUrl}/login`
    })
    expect(req[0].request.method).toEqual('POST');
    req[0].flush(mockUsers);
  });
});
