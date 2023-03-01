import { NgModule } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CdTimerComponent, CdTimerModule } from 'angular-cd-timer';
import { DemoPageComponent } from './demo-page.component';

describe('DemoPageComponent', () => {
  let component: DemoPageComponent;
  let fixture: ComponentFixture<DemoPageComponent>;

  beforeEach(async () => {
    // config testing module
    // module -- import components, providers, routing, ect 
    await TestBed.configureTestingModule({
      declarations: [ DemoPageComponent, CdTimerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DemoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
