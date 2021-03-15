import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CustomerViewComponent } from './customer-view/customer-view.component';
import { DetailComponent } from './product-detail/detail.component';

const routes: Routes = [
  {
    path: 'customer',
    component: CustomerViewComponent
  },
  {
    path: 'detail',
    component: DetailComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
