import { Component } from '@angular/core';
import { VulnListComponent } from '../vulnerabilities/vuln-list.component';
import { VulnerabilityService } from '../vulnerabilities/vulnerability.service';

@Component({
    selector: 'w3af-result-organizer',
    template:
    `
        <h1>{{title}}</h1>
        <vulnerability-list>

    `,
    directives: [VulnListComponent],
    providers: [VulnerabilityService]
})
export class AppComponent {
  title: String = "W3af result organizer";
}
