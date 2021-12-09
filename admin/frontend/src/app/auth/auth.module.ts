
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { XlModule } from "@stypw/xl";
import { AuthComponent } from './auth.component';


export const routes: Routes = [
    { path: "", component: AuthComponent }
]

@NgModule({
    declarations: [
        AuthComponent
    ],
    imports: [
        XlModule,
        RouterModule,
        RouterModule.forChild(routes)
    ]
})
export class AuthModule { }
