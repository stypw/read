import { Component } from '@angular/core';

export type Menu = {
  text:string;
  url:string | string[];
}

@Component({
  selector: 'div[router-workbench]',
  templateUrl: './workbench.component.html',
  styleUrls: ['./workbench.component.scss']
})
export class WorkbenchComponent {
  menus:Menu[]=[
    {url:["/workbench","start"],text:"开始"},
    {url:["/workbench","wnd"],text:"弹窗"}
  ]
}
