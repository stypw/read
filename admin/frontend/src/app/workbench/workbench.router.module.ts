import { Routes } from "@angular/router";
import { StartComponent } from "./start/start.component";
import { WndComponent } from "./wnd/wnd.component";
import { WorkbenchComponent } from "./workbench.component";


export const routeComponents = [
    WorkbenchComponent,
    StartComponent,
    WndComponent
];

export const routes: Routes = [
    {
        path: "",
        component: WorkbenchComponent,
        children: [
            { path: "", pathMatch: "full", redirectTo: "start" },
            { path: "start", component: StartComponent },
            { path: "wnd", component: WndComponent },
            { path: "**", pathMatch: "full", redirectTo: "start" }
        ]
    }
]