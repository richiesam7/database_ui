import { Component, OnInit } from '@angular/core';
import { Customer } from '../field-framework/entityClasses/customer';

@Component({
  selector: 'app-customer-view',
  templateUrl: './customer-view.component.html',
  styleUrls: ['./customer-view.component.less']
})
export class CustomerViewComponent implements OnInit {

  public customer;
  public dynamicFields;


  constructor() {
    this.showCustomer();
  }

  ngOnInit(): void {
    this.showCustomer();
  }

  fetchCustomer() {
    this.customer = {
      CustomerID: 1111,
      CustomerName: 'Richie Sam',
    }; // replace with API fetch
  }

  // basic view model implementation - will be moved to framework
  showCustomer() {
    let classDefn = new Customer();
    let classFacts = classDefn.getEntity();
    // get template for text field from field-viewmodel
    for (let fact in classFacts) {
      // classFacts[fact] - TODO
      // add to array here
    }
    this.dynamicFields = [{name: 'Customer ID', value: '11221', type: 'number'}, {name: 'Taxation Category', value: 'GST', type: 'text'}];
  }

}
