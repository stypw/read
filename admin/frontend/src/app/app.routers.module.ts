import { Routes } from "@angular/router";
import { StartComponent } from "./start/start.component";
import { WndComponent } from "./wnd/wnd.component";
export const routeComponents = [
    StartComponent,
    WndComponent
];

export const routes: Routes = [
    { path: "", pathMatch: "full", redirectTo: "start" },
    { path: "start", component: StartComponent },
    { path: "wnd", component: WndComponent },
    { path: "**", pathMatch: "full", redirectTo: "start" }
]