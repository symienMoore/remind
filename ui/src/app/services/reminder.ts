import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { filter, map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class Reminder {
  private http = inject(HttpClient);
  constructor() {

  }
  getdata = () => {
    this.http.get('https://jsonplaceholder.typicode.com/todos/')
    .pipe(
      map((todos: any) =>
        todos.filter((todo: any) =>
      todo.completed))
    ).subscribe(x => {
      console.log(x)
    })
  }

  testApi = () => {
    this.http.get('http://localhost:8080/ping').subscribe(x => {
      console.log(x)
    })
  }
}
