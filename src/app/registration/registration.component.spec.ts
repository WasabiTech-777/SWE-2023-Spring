import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { RegistrationComponent } from './registration.component';
import { MaterialModule } from 'app/material.module';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { environment } from 'environments/environment';
import { AccountService } from 'app/account.service';

describe('RegistrationComponent', () => {
  let component: RegistrationComponent;
  let fixture: ComponentFixture<RegistrationComponent>;
  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;
  let service: AccountService

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegistrationComponent ],
      imports: [HttpClientTestingModule, MaterialModule, ReactiveFormsModule, BrowserAnimationsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegistrationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
    service = TestBed.inject(AccountService)
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call postUsers and return the new user info', () => {
    let mockUsers = {uname: "hell02"};
    service.createAccount("hell02", "test").subscribe((response) => {expect(response).toEqual(jasmine.objectContaining(mockUsers));});
    const req = httpTestingController.match({
        method: 'POST',
        url: `${environment.serverUrl}/users`
    })
    expect(req[0].request.method).toEqual('POST');
    req[0].flush(mockUsers);
  });
});
