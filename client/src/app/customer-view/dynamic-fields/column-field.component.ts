import { Component, Input } from '@angular/core';

@Component({
  template: `
    <div>
      <h4>{{data.name}}</h4>
    </div>
    <div *ngIf="true">
    <text-field fieldName="{{data.name}}" fieldInputValue="{{data.value}}"></text-field>
    </div>
  `
})
export class ColumnFieldComponent {
  @Input() data: any;
}
