import { Component } from '@angular/core';
import { ROUTER_DIRECTIVES} from '@angular/router';
import { HTTP_PROVIDERS } from '@angular/http';
import { VulnListComponent } from '../vulnerabilities/vuln-list.component';
import { VulnerabilityService } from '../vulnerabilities/vulnerability.service';

@Component({
    selector: 'w3af-result-organizer',
    directives: [VulnListComponent],
    providers: [VulnerabilityService, HTTP_PROVIDERS],
    template:
    `
        <h1>{{title}}</h1>
        <vulnerability-list>

    `
})
export class AppComponent {
  title: String = "W3af result organizer";
}
