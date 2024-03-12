import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  favColor: string = "#1680f7";

  constructor(
    private userService: UserService,
    private router: Router,
  ) { }

  ngOnInit(): void { }

  onFormSubmit(form: NgForm) {
    const name = form.value.name;
    const email = form.value.email;
    const color = this.favColor;

    this.userService.postRegister(name, email, color).subscribe((data) => {
      // Save info to use in home page
      localStorage.setItem('userName', name)
      localStorage.setItem('userId', data.user_id)
      localStorage.setItem('favColor', color)

      // console.log(this.favColor)

      // Once we've received a response, take the user to the home page
      this.router.navigateByUrl('/home');
    })
  }
  // onColorChange(event:any){
  //   this.favColor = event.target.value;
  //   console.log(this.favColor)
  // }
}
