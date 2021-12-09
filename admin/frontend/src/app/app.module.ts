import { Injectable, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { CanLoad, Router, RouterModule, Routes } from '@angular/router';
import { AuthService } from '@common/auth';
import { XlModule, http, mysql } from "@stypw/xl";
import { AppComponent } from './app.component';

@Injectable()
export class CanWorkbenchLoad implements CanLoad {
  async canLoad() {
    if (await this.authService.checkLogin()) {
      return true;
    }
    this.router.navigate(["/auth"]);
    return false
  }
  constructor(private router: Router, private authService: AuthService) { }
}


export const routes: Routes = [
  { path: "", pathMatch: "full", redirectTo: "workbench" },
  { path: "workbench", loadChildren: () => import("./workbench/workbench.module").then(mod => mod.WorkbenchModule), canLoad: [CanWorkbenchLoad] },
  { path: "auth", loadChildren: () => import("./auth/auth.module").then(mod => mod.AuthModule) },
  { path: "**", pathMatch: "full", redirectTo: "workbench" }
]


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    XlModule,
    RouterModule,
    RouterModule.forRoot(routes)
  ],
  providers: [AuthService, CanWorkbenchLoad],
  bootstrap: [AppComponent]
})
export class AppModule { }
