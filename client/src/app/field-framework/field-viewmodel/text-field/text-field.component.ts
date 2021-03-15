import { Component, OnInit, Input } from '@angular/core';
import { Customer } from '../../entityClasses/customer';
import { GenericClass } from '../../entityClasses/genericClass';
import { Column } from '../../column';

@Component({
  selector: 'text-field',
  templateUrl: './text-field.component.html',
  styleUrls: ['./text-field.component.less']
})
export class TextFieldComponent implements OnInit {
  public entity: GenericClass;
  private isEditing: boolean;

  get fieldValue() {
    return this.entity[this.fieldName];
  }

  set fieldValue(val) {
    if (this.entity[this.fieldName] !== val) {
      this.entity[this.fieldName] = val;
      // emit event on field change
    }
  }

  @Input() fieldName: string;
  @Input() fieldInputValue: string;

  constructor() { 
    this.entity = new Customer(); // should be column-entity when ready
  }

  ngOnInit(): void {
    this.setField();
    // subscribe to change in field event
  }

  public onFieldChange($event) {
    this.fieldValue = $event.target.value;
  }

  public isReadOnly(): boolean {
    const field : Column  = this.entity.getField(this.fieldName);
    return false; // TODO : use correct field
  }

  private setField(): void {
    const entity = this.entity.getField(this.fieldName);;
    if (entity !== this.entity) {
      // this.entity = entity;
      this.fieldValue = this.fieldInputValue;
    }
  }
}
