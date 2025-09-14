import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class Auth {
  private http = inject(HttpClient);
  userLoggedIn: boolean = false;
  userToken: string = "";
  userLoggedin: boolean = false;
  // apiURL: string
  //TODO: get token from response at http://localhost:8080/auth/login {username: test, password: test}
  constructor() {

  }

  doUserLogin({username, password}: any) {
    return this.http.post("http://localhost:8080/auth/login", {username, password})
  }

}
