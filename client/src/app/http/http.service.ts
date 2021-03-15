import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HttpHandler {

  private REST_API_SERVER = "http://localhost:8080/headphones";

  constructor(private httpClient: HttpClient) { }

  public getRequest(){
    return this.httpClient.get(this.REST_API_SERVER);
  }
}