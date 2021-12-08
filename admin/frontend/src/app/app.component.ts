import { Component } from '@angular/core';

export type Menu = {
  text:string;
  url:string;
}

@Component({
  selector: 'div[app-root]',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  menus:Menu[]=[
    {url:"start",text:"开始"},
    {url:"wnd",text:"弹窗"}
  ]
}
