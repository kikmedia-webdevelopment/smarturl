import { StackProps } from 'components/Stack/Stack';
import { TransitionStatus } from 'react-transition-group/Transition'
import { defaultProps } from './Positioner';

export type Position = 'top' | 'top-left' | 'top-right' | 'bottom' | 'bottom-left' | 'bottom-right' | 'left' | 'right'
export type PositionState = TransitionStatus

export type PositionerProps = {
    position?: Position
    isShown?: boolean
    children: (params: {
        top: number,
        left: number,
        zIndex: NonNullable<StackProps['value']>,
        style: {
            transformOrigin: string,
            left: number,
            top: number,
            zIndex: NonNullable<StackProps['value']>,
        },
        getRef: (ref: React.RefObject<HTMLElement>) => void,
        animationDuration: PositionerProps['animationDuration'],
        state: PositionState
        css: any
    }) => React.ReactNode
    innerRef?: (ref: React.RefObject<HTMLElement>) => void
    bodyOffset?: number
    targetOffset?: number
    target: (params: { getRef: () => React.RefObject<HTMLElement>, isShown: boolean }) => React.ReactNode
    initialScale?: number
    animationDuration?: number
    onCloseComplete?: () => void
    onOpenComplete?: () => void
} & typeof defaultProps