import { Component } from '@angular/core';
import { Vulnerability } from './vulnerability';
import { VulnerabilityService } from './vulnerability.service';

@Component({
  selector: 'vulnerability-list',
  template:
  `
    <ul>
      <li *ngFor="let vuln of vulns" (click)="selectVulnerability(vuln)">
        {{vuln.name}}
      </li>
    </ul>

    <section *ngIf="selectedVulnerability">
      {{selectedVulnerability.url}}
    </section>
  `
})

export class VulnListComponent{
  vulns: Vulnerability[] = [];
  selectedVulnerability: Vulnerability;

  constructor(private _vulnService: VulnerabilityService){}

  ngOnInit(){
    this.vulns = this._vulnService.getAll();
  }

  selectVulnerability(vuln: Vulnerability){
    this.selectedVulnerability = vuln;
  }
}
