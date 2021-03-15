import { Component, Input, OnInit, ViewChild, ComponentFactoryResolver, OnDestroy } from '@angular/core';
import { AnchorDirective } from './dynamic-fields.directive';
import { ColumnFieldComponent } from './column-field.component';

@Component({
  selector: 'dynamic-fields',
  template: `
              <div>
                <ng-template dynamicFld></ng-template>
              </div>
            `
})
export class DynamicFieldsComponent implements OnInit {
  @Input() fields;
  @ViewChild(AnchorDirective, {static: true}) dynamicFld: AnchorDirective;

  constructor(private componentFactoryResolver: ComponentFactoryResolver) { }

  ngOnInit() {
    this.dynamicComponentLoad();
  }

  dynamicComponentLoad() {
    const item = this.fields[1]; // TODO: In a loop
    const componentFactory = this.componentFactoryResolver.resolveComponentFactory(ColumnFieldComponent);
    const viewContainerRef = this.dynamicFld.viewContainerRef;
    viewContainerRef.clear();
    const componentRef = viewContainerRef.createComponent(componentFactory);
    componentRef.instance.data = item;
  }
}


