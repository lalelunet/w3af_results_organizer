import { Component, Input } from '@angular/core';
import { Vulnerability } from './vulnerability';

@Component({
  selector: 'vulnerability-details',
  templateUrl: 'app/templates/vuln-details.component.html'
})

export class VulnDetailsComponent{
  @Input() vuln: Vulnerability;

  saveVulnDetails(){
    alert(`${JSON.stringify(this.vuln)}`);
}

}
