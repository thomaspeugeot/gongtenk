import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {

  view = 'Carto view'
  carto = 'Carto view'
  data = 'Data view'
  leaflet = 'Leaflet view'
  sim = 'Sim view'
  diagrams = 'Diagrams view'
  views: string[] = [this.carto, this.data, this.leaflet, this.sim, this.diagrams];

  title = 'ng';
}
