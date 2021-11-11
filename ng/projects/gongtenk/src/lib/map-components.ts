// insertion point sub template for components imports 
  import { IndividualsTableComponent } from './individuals-table/individuals-table.component'
  import { IndividualSortingComponent } from './individual-sorting/individual-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfIndividualsComponents: Map<string, any> = new Map([["IndividualsTableComponent", IndividualsTableComponent],])
  export const MapOfIndividualSortingComponents: Map<string, any> = new Map([["IndividualSortingComponent", IndividualSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Individual", MapOfIndividualsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Individual", MapOfIndividualSortingComponents],
    ]
  )
