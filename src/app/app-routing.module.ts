import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DemoPageComponent } from './demo-page/demo-page.component';
import { LogInComponent } from './log-in/log-in.component';
import { RegistrationComponent } from './registration/registration.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { ArticlesPageComponent } from './articles-page/articles-page.component';

const routes: Routes = [
  { path: '', component: LogInComponent},
  { path: 'demo-page', component: DemoPageComponent},
  { path: 'register', component: RegistrationComponent},
  { path: 'profile', component: ProfilePageComponent},
  { path: 'articles', component: ArticlesPageComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
