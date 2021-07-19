import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MainRoutingModule } from './main-routing.module';
import { MainComponent } from './main.component';
import { SharedModule } from '../../shared/shared.module';
import { UpcomingComponent } from './upcoming/upcoming.component';
import { WatchlistComponent } from './watchlist/watchlist.component';


@NgModule({
  declarations: [
    MainComponent,
    UpcomingComponent,
    WatchlistComponent
  ],
  imports: [
    CommonModule,
    MainRoutingModule,
    SharedModule
  ]
})
export class MainModule { }
