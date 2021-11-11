import { NgModule } from '@angular/core';
import { GongtenkspecificComponent } from './gongtenkspecific.component';
import { TenkMapComponent } from './tenk-map/tenk-map.component';



@NgModule({
  declarations: [
    GongtenkspecificComponent,
    TenkMapComponent
  ],
  imports: [
  ],
  exports: [
    GongtenkspecificComponent
  ]
})
export class GongtenkspecificModule { }
