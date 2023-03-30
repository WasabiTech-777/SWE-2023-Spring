import { NgModule } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { CdTimerComponent, CdTimerModule } from 'angular-cd-timer';

import { DemoPageComponent } from './demo-page.component';

describe('DemoPageComponent', () => {
  let component: DemoPageComponent;
  let fixture: ComponentFixture<DemoPageComponent>;
  let snackBar: MatSnackBar;

  beforeEach(async () => {
    // config testing module
    // module -- import components, providers, routing, ect 
    await TestBed.configureTestingModule({
      declarations: [ DemoPageComponent ],
      imports: [CdTimerModule, FormsModule, MatSnackBar, MatSnackBarModule],
      providers: [ MatSnackBar ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DemoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    snackBar = TestBed.inject(MatSnackBar);
  });

  it('Current Wiki Prompt is Japan', () => {
    const expectedString = 'Japan';
    const resultString = component.wiki.name;

    expect(resultString).toEqual(expectedString);
  });

});

describe('compare() function test', () => {
  let component: DemoPageComponent;
  let fixture: ComponentFixture<DemoPageComponent>;

  beforeEach(async () => {
    // config testing module
    // module -- import components, providers, routing, ect 
    await TestBed.configureTestingModule({
      declarations: [ DemoPageComponent ],
      imports: [CdTimerModule, FormsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DemoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('return correct', () => {

    const char1 = 'i';
    const char2 = 'i';
    const strIndex = 0; 
    const resultString = component.compare(char1, char2, strIndex);
    const numCorrect = component.correct;

    expect(resultString).toEqual('correct');
    expect(numCorrect).toEqual(1);
  });

  it('return incorrect', () => {

    const char1 = 'i';
    const char2 = 'j';
    const strIndex = 0; 
    const resultString = component.compare(char1, char2, strIndex);
    const numIncorrect = component.errors;

    expect(resultString).toEqual('incorrect');
    expect(numIncorrect).toEqual(1);
    
  });

  it('return pending', () => {

    const char1 = 'i';
    const char2 = '';
    const strIndex = 0; 
    const resultString = component.compare(char1, char2, strIndex);

    expect(resultString).toEqual('pending');
    
  });

  it('when wiki.text.length is reached execute onTimeOver()', () => {
    const char1 = 'i';
    const char2 = 'i';
    const strIndex = 5; 

    component.wiki.text = '012345';
    component.correct = 5;

    const resultString = component.compare(char1, char2, strIndex);

    expect(component.WPM).toEqual(6*5/15);


  });

});

describe('onTimeOver() function test', () => {
  let component: DemoPageComponent;
  let fixture: ComponentFixture<DemoPageComponent>;

  beforeEach(async () => {
    // config testing module
    // module -- import components, providers, routing, ect 
    await TestBed.configureTestingModule({
      declarations: [ DemoPageComponent ],
      imports: [CdTimerModule, FormsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DemoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('if no user prompt typed when time ends', () => {

    const char1 = 'i';
    const char2 = 'i';
    const strIndex = 0; 
    const resultString = component.onTimeOver();
    const numCorrect = component.correct;

    expect(component.WPM).toEqual(-1);
    expect(numCorrect).toEqual(0);
  });

});
