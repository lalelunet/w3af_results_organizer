import { Component } from '@angular/core';
import { Vulnerability } from './vulnerability';
import { VulnerabilityService } from './vulnerability.service';

@Component({
  selector: 'vulnerability-list',
  template:
  `
    <ul>
      <li *ngFor="let vuln of vulns">
        {{vuln.name}}
      </li>
    </ul>
  `
})

export class VulnListComponent{
  vulns: Vulnerability[] = [];

  constructor(private _vulnService: VulnerabilityService){}

  ngOnInit(){
    this.vulns = this._vulnService.getAll();
  }
}
