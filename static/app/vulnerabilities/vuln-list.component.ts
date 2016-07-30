import { Component } from '@angular/core';
import { ROUTER_DIRECTIVES } from '@angular/router';
import { Vulnerability } from './vulnerability';
import { VulnerabilityService } from './vulnerability.service';
import { VulnDetailsComponent } from './vuln-details.component';

@Component({
  selector: 'vulnerability-list',
  directives: [VulnDetailsComponent],
  templateUrl: 'app/templates/vuln-list.component.html'
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
