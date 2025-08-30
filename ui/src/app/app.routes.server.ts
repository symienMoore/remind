import { RenderMode, ServerRoute } from '@angular/ssr';
// import { Component } from '@angular/core';
import { Reminder } from './compnents/reminder/reminder';

export const serverRoutes: ServerRoute[] = [
  {
    path: '**',
    renderMode: RenderMode.Prerender
  },
  // {
  //   path: 'reminder',
  //   renderMode: RenderMode.Prerender
  // }
];
