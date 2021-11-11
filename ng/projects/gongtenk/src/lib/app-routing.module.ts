import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { IndividualsTableComponent } from './individuals-table/individuals-table.component'
import { IndividualDetailComponent } from './individual-detail/individual-detail.component'
import { IndividualPresentationComponent } from './individual-presentation/individual-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongtenk_go-individuals', component: IndividualsTableComponent, outlet: 'gongtenk_go_table' },
	{ path: 'gongtenk_go-individual-adder', component: IndividualDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-individual-adder/:id/:originStruct/:originStructFieldName', component: IndividualDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-individual-detail/:id', component: IndividualDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-individual-presentation/:id', component: IndividualPresentationComponent, outlet: 'gongtenk_go_presentation' },
	{ path: 'gongtenk_go-individual-presentation-special/:id', component: IndividualPresentationComponent, outlet: 'gongtenk_goindividualpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
