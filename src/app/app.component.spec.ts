import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { AppComponent } from './app.component';
import { MaterialModule } from 'app/material.module';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HelloWorldService } from './hello-world.service';
import { environment } from 'environments/environment';

describe('AppComponent', () => {
  let component: AppComponent;
  let fixture: ComponentFixture<AppComponent>;
  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;
  let service: HelloWorldService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AppComponent ],
      imports: [HttpClientTestingModule, MaterialModule, ReactiveFormsModule, BrowserAnimationsModule, RouterTestingModule],
      providers: [HelloWorldService]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
    service = TestBed.inject(HelloWorldService);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call getUsers and return a JSON of Users', () => {
    let mockUsers = {uname: "hell0"};
    service.getTitle().subscribe((response) => {expect(response).toEqual(jasmine.objectContaining(mockUsers));});
    const req = httpTestingController.match({
        method: 'GET',
        url: `${environment.serverUrl}/users`
    })
    expect(req[0].request.method).toEqual('GET');
    req[0].flush(mockUsers);
  });
	
});