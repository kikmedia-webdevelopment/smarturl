import React from 'react'
import constants from './constants'

/**
 * Context used to manage the layering of z-indexes of components.
 */
const StackingContext = React.createContext(constants.STACKING_CONTEXT)
export default StackingContext