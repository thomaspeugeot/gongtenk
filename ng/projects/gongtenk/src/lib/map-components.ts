// insertion point sub template for components imports 
  import { ConfigurationsTableComponent } from './configurations-table/configurations-table.component'
  import { ConfigurationSortingComponent } from './configuration-sorting/configuration-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfConfigurationsComponents: Map<string, any> = new Map([["ConfigurationsTableComponent", ConfigurationsTableComponent],])
  export const MapOfConfigurationSortingComponents: Map<string, any> = new Map([["ConfigurationSortingComponent", ConfigurationSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Configuration", MapOfConfigurationsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Configuration", MapOfConfigurationSortingComponents],
    ]
  )
