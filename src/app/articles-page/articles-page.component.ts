import { Component } from '@angular/core';
import {HttpClient, HttpResponse} from '@angular/common/http';
import { environment } from 'environments/environment';
import { Router } from '@angular/router';

@Component({
  selector: 'app-articles-page',
  templateUrl: './articles-page.component.html',
  styleUrls: ['./articles-page.component.css']
})
export class ArticlesPageComponent {
  constructor(private http: HttpClient, private router: Router) { }
  dataSource = ArticleList
  displayedColumns: string[] = ['title', 'url', 'length'];
  ngOnInit() {
    //this.http.get(`${environment.serverUrl}/article/0`).subscribe(response => console.log(response));
  }
  onClick() {
    this.router.navigate(['/demo-page']);
  }
}
export interface Article {
  title: string;
  url: string;
  length: number;
}
const ArticleList: Article[] = [
  {title: "Japan", url: "https://simple.wikipedia.org/wiki/Japan", length: 525}
];
