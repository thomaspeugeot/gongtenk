// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ConfigurationDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	NumberOfCitiesToDisplay: number = 0
	NumberOfCitiesToDisplay_real: number = 0

	// insertion point for other declarations
}
