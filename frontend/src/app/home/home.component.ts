import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  userName: String = '';
  userId: String = '';
  favColor: String = '';

  constructor() { }

  ngOnInit(): void {
    // Get saved userName and userId
    const userName = localStorage.getItem('userName');
    const userId = localStorage.getItem('userId');
    const favColor = localStorage.getItem('favColor')

    this.userName = userName ? userName : '';
    this.userId = userId ? userId : '';
    this.favColor = favColor ? favColor : '';
  }

}
