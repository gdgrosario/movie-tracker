import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main.component';
import { UpcomingComponent } from './upcoming/upcoming.component';
import { WatchlistComponent } from './watchlist/watchlist.component';

const routes: Routes = [
  { 
    path: '', 
    component: MainComponent,
    children: [
      {
        path: 'upcoming',
        component: MainComponent
      },
      {
        path: 'watchlist',
        component: MainComponent
      },
      {
        path: 'favorites',
        component: MainComponent
      },
      {
        path: 'profile',
        component: MainComponent
      },
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MainRoutingModule { }
