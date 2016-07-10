// currently not in use
// to activate it uncomment the include in the main.js
import { provideRouter, RouterConfig } from '@angular/router';
import { VulnListComponent } from './vulnerabilities/vuln-list.component';
import { VulnDetailsComponent } from './vulnerabilities/vuln-details.component';

const routes: RouterConfig =[
  {
    path: 'vulns',
    component: VulnListComponent
  },
  {
    path: 'vulns/:id',
    component: VulnDetailsComponent
  },
  {
    path: '',
    redirectTo: '/vulns',
    pathMatch: 'full'
  }
];

export const APP_ROUTER_PROVIDERS = [
  provideRouter(routes)
];
