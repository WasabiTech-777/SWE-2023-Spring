import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { AccountService } from './account.service';
import { environment } from 'environments/environment';

describe('AccountService', () => {
  let service: AccountService;
  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({imports: [HttpClientTestingModule]});
    service = TestBed.inject(AccountService);
    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  it('should validate a token using validate()', () => {
    service.validate();
    const req = httpTestingController.match({
        method: 'POST',
        url: `${environment.serverUrl}/token`
    })
    expect(req[0].request.method).toEqual('POST');
  });
  it('should decode a token and provide the username', () => {
    var tokenVal = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmFtZSI6ImhlbGxvMTAiLCJleHAiOjE2ODAyMTE4NjZ9.hW3tAf1mnDMGAzsLt-sAKxSg29LTZYym1W69Qa-jVRA"
    expect((service.decodeToken(tokenVal)['uname'] as string)).toBeTruthy();
  });
});
