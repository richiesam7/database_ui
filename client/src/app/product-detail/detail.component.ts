import { Component, OnInit } from '@angular/core';
import { Customer } from '../field-framework/entityClasses/customer';
import { HttpHandler } from '../http/http.service';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.less']
})
export class DetailComponent implements OnInit {

  public customer;
  public dynamicFields;
  public products = [];
  public url = '';

  constructor(private httpHandler: HttpHandler) { }

  ngOnInit() {
    this.httpHandler.getRequest().subscribe((data: any[])=>{
      console.log(data);
      this.products = data;
      this.products.forEach(product => {
        this.url = product.image_url
      });
      this.url = this.products[0].image_url
    })  
    this.showCustomer();
  }

  fetchCustomer() {
    this.customer = {
      CustNodeID: 1111,
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
    this.dynamicFields = [{name: 'Customer ID', value: '11221'}, {name: 'Taxation Category', value: 'GST'}];
  }

}
