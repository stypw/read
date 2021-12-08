import { Component } from "@angular/core";
import { http } from "@stypw/xl"
@Component({
    selector: "div[example-start]",
    templateUrl: "./start.component.html",
    styleUrls: ["./start.component.scss"]
})
export class StartComponent {

    async doLogin() {
        const [err, ret] = await http.postAsync("/api/auth", { acc: "stypw", pwd: "stypw_123456" });
        console.log(err, ret);
    }


}