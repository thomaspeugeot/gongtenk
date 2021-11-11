// insertion point sub template for components imports 
  import { CitysTableComponent } from './citys-table/citys-table.component'
  import { CitySortingComponent } from './city-sorting/city-sorting.component'
  import { CountrysTableComponent } from './countrys-table/countrys-table.component'
  import { CountrySortingComponent } from './country-sorting/country-sorting.component'
  import { IndividualsTableComponent } from './individuals-table/individuals-table.component'
  import { IndividualSortingComponent } from './individual-sorting/individual-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfCitysComponents: Map<string, any> = new Map([["CitysTableComponent", CitysTableComponent],])
  export const MapOfCitySortingComponents: Map<string, any> = new Map([["CitySortingComponent", CitySortingComponent],])
  export const MapOfCountrysComponents: Map<string, any> = new Map([["CountrysTableComponent", CountrysTableComponent],])
  export const MapOfCountrySortingComponents: Map<string, any> = new Map([["CountrySortingComponent", CountrySortingComponent],])
  export const MapOfIndividualsComponents: Map<string, any> = new Map([["IndividualsTableComponent", IndividualsTableComponent],])
  export const MapOfIndividualSortingComponents: Map<string, any> = new Map([["IndividualSortingComponent", IndividualSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["City", MapOfCitysComponents],
      ["Country", MapOfCountrysComponents],
      ["Individual", MapOfIndividualsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["City", MapOfCitySortingComponents],
      ["Country", MapOfCountrySortingComponents],
      ["Individual", MapOfIndividualSortingComponents],
    ]
  )
