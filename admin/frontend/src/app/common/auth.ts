import { Injectable } from "@angular/core";
import { http } from "@stypw/xl";

type UserInfo = { acc: string; nickname: string }

@Injectable({ providedIn: "root" })
export class AuthService {
    private userInfo: UserInfo | null = null;
    async checkLogin(): Promise<boolean> {
        if (this.userInfo) return true;
        const ret = await http.getWithJson("/api/auth", null)
        if (ret && ret.code == 0) {
            this.userInfo = ret.data;
        }
        if (this.userInfo) return true;
        return false
    }
    async doLogin(acc: string, pwd: string) {
        const ret = await http.postWithJson("/api/login", { acc, pwd });
        if (ret && ret.code == 0) {
            this.userInfo = ret.data;
            return true;
        } else {
            return false;
        }
    }
    get UserInfo() {
        return this.userInfo;
    }
}
