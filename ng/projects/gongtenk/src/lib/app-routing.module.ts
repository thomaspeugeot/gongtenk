import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CitysTableComponent } from './citys-table/citys-table.component'
import { CityDetailComponent } from './city-detail/city-detail.component'
import { CityPresentationComponent } from './city-presentation/city-presentation.component'

import { CountrysTableComponent } from './countrys-table/countrys-table.component'
import { CountryDetailComponent } from './country-detail/country-detail.component'
import { CountryPresentationComponent } from './country-presentation/country-presentation.component'

import { IndividualsTableComponent } from './individuals-table/individuals-table.component'
import { IndividualDetailComponent } from './individual-detail/individual-detail.component'
import { IndividualPresentationComponent } from './individual-presentation/individual-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongtenk_go-citys', component: CitysTableComponent, outlet: 'gongtenk_go_table' },
	{ path: 'gongtenk_go-city-adder', component: CityDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-city-adder/:id/:originStruct/:originStructFieldName', component: CityDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-city-detail/:id', component: CityDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-city-presentation/:id', component: CityPresentationComponent, outlet: 'gongtenk_go_presentation' },
	{ path: 'gongtenk_go-city-presentation-special/:id', component: CityPresentationComponent, outlet: 'gongtenk_gocitypres' },

	{ path: 'gongtenk_go-countrys', component: CountrysTableComponent, outlet: 'gongtenk_go_table' },
	{ path: 'gongtenk_go-country-adder', component: CountryDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-country-adder/:id/:originStruct/:originStructFieldName', component: CountryDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-country-detail/:id', component: CountryDetailComponent, outlet: 'gongtenk_go_editor' },
	{ path: 'gongtenk_go-country-presentation/:id', component: CountryPresentationComponent, outlet: 'gongtenk_go_presentation' },
	{ path: 'gongtenk_go-country-presentation-special/:id', component: CountryPresentationComponent, outlet: 'gongtenk_gocountrypres' },

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
