import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { XlModule } from "@stypw/xl";
import { AppComponent } from './app.component';

import { routes,routeComponents } from "./app.routers.module";


@NgModule({
  declarations: [
    AppComponent,
    ...routeComponents
  ],
  imports: [
    BrowserModule,
    XlModule,
    RouterModule,
    RouterModule.forRoot(routes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
