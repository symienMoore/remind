import { Component } from '@angular/core';
import {FormsModule} from '@angular/forms';
import { Reminder as ReminderService } from '../../services/reminder';

@Component({
  selector: 'app-reminder',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './reminder.html',
  styleUrl: './reminder.css'
})
export class Reminder {
  message: string = "wow"
  new_message = ""
  constructor(private reminderService: ReminderService) {
    // this.reminderService.getdata();
    this.reminderService.testApi();
  }


}
