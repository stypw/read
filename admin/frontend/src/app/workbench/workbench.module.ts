
import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { XlModule } from "@stypw/xl";
import { WorkbenchComponent } from './workbench.component';

import { routeComponents, routes } from "./workbench.router.module";

@NgModule({
    declarations: [
        WorkbenchComponent,
        ...routeComponents
    ],
    imports: [
        XlModule,
        CommonModule,
        RouterModule,
        RouterModule.forChild(routes)
    ]
})
export class WorkbenchModule { }
