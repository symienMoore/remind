import { Routes } from '@angular/router';
import { Home } from './pages/home/home'
import { Reminder } from './compnents/reminder/reminder';


export const routes: Routes = [
  {
    path: 'reminder',
    component: Reminder
  },
  {
    path: '',
    component: Home
  }
];
