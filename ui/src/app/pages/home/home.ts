import { Component } from '@angular/core';
import { ReactiveFormsModule, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Auth } from '../../services/auth'
@Component({
  selector: 'app-home',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './home.html',
  styleUrl: './home.css'
})
export class Home {



    // login: login 
  loginForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });


  constructor(private auth: Auth) {}

  doUserLogin() {
    console.log('Payload sent to backend:', this.loginForm.value);

    console.log(this.loginForm.value);
    this.auth.doUserLogin(this.loginForm.value).subscribe({
      next: res => console.log(res),
      error: err => console.error(err)
    })
  }

}
