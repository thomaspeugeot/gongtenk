import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ConfigurationsTableComponent } from './configurations-table/configurations-table.component'
import { ConfigurationDetailComponent } from './configuration-detail/configuration-detail.component'
import { ConfigurationPresentationComponent } from './configuration-presentation/configuration-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongtenk_go-configurations', component: ConfigurationsTableComponent, outlet: 'gongtenk_go_table' },
	{ path: 'gongtenk_go-configuration-adder', component: ConfigurationDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-configuration-adder/:id/:originStruct/:originStructFieldName', component: ConfigurationDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-configuration-detail/:id', component: ConfigurationDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-configuration-presentation/:id', component: ConfigurationPresentationComponent, outlet: 'gongtenk_go_presentation' },
	{ path: 'gongtenk_go-configuration-presentation-special/:id', component: ConfigurationPresentationComponent, outlet: 'gongtenk_goconfigurationpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
