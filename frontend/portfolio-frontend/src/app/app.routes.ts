import { Routes } from '@angular/router';
import { Skills } from './pages/skills/skills';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'skills',
    pathMatch: 'full',
  },

  {
    path: 'skills',
    component: Skills,
  },
];