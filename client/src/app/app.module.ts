import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CustomerViewComponent } from './customer-view/customer-view.component';
import { TextFieldComponent } from './field-framework/field-viewmodel/text-field/text-field.component';
import { DynamicFieldsComponent } from './customer-view/dynamic-fields/dynamic-fields.component';
import { AnchorDirective } from './customer-view/dynamic-fields/dynamic-fields.directive';
import { ColumnFieldComponent } from './customer-view/dynamic-fields/column-field.component';

@NgModule({
  declarations: [
    AppComponent,
    CustomerViewComponent,
    TextFieldComponent,
    AnchorDirective,
    DynamicFieldsComponent,
    ColumnFieldComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
