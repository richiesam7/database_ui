import { Directive, ViewContainerRef } from '@angular/core';
@Directive({
  selector: '[dynamicFld]',
})
export class AnchorDirective {
  constructor(public viewContainerRef: ViewContainerRef) { }
}
