import React from 'react'
import constants from './constants'
import StackingContext from './StackingContext'

export type StackProps = {
    children: (zIndex: number) => React.ReactNode
    value?: number
}

export const Stack: React.FC<StackProps> = ({
    children,
    value = constants.STACKING_CONTEXT
}) => {
    return (
        <StackingContext.Consumer>
            {previousValue => {
                const currentValue = Math.max(value, previousValue)
                const nextValue = currentValue + 1

                return (
                    <StackingContext.Provider value={nextValue}>
                        {children(currentValue)}
                    </StackingContext.Provider>
                )
            }}
        </StackingContext.Consumer>
    )
}