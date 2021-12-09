import { Component } from "@angular/core";
import { Router } from "@angular/router";
import { AuthService } from "@common/auth";
import { tools } from "@stypw/xl"
@Component({
    selector: "div[router-auth]",
    templateUrl: "./auth.component.html",
    styleUrls: ["./auth.component.scss"]
})
export class AuthComponent {

    busy = false;
    async doLogin() {
        if (this.busy) return;
        this.busy = true;
        const ret = await this.authService.doLogin("stypw", "stypw_123456");
        if (ret) {
            await tools.sleep(2000);
            this.router.navigate(["/workbench"]);
        }
        this.busy = false;
    }
    constructor(private authService: AuthService, private router: Router) { }
}