import { Component } from '@angular/core';
import { Reminder as ReminderService } from '../../services/reminder';

@Component({
  selector: 'app-reminder',
  standalone: true,
  imports: [],
  templateUrl: './reminder.html',
  styleUrl: './reminder.css'
})
export class Reminder {
  constructor(private reminderService: ReminderService) {
    // this.reminderService.getdata();
    this.reminderService.testApi();
  }
}
