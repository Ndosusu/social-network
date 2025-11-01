import { NgClass } from '@angular/common';
import { Component} from '@angular/core';

@Component({
  selector: 'app-not-found',
  imports: [NgClass],
  templateUrl: './not-found.html',
  styleUrl: './not-found.css',
})
export class NotFound {
  darkMode = localStorage.getItem("darkmode") == "1" ? true : false

  toggleLightMode() {
    console.log("ok")
    this.darkMode = !this.darkMode
    localStorage.setItem("darkmode", this.darkMode? "1": "0")
  }
}
