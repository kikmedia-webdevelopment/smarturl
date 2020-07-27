import React from 'react'
import { PositionerProps } from './types'
import Transition from 'react-transition-group/Transition'
import getPosition from './getPosition'
import { Stack } from 'components/Stack'
import { Portal } from 'components/Portal'

const animationEasing = {
    spring: `cubic-bezier(0.175, 0.885, 0.320, 1.175)`
}

const initialState = () => ({
    top: 0,
    left: 0,
    transformOrigin: ''
})

const getCSS = ({ initialScale, animationDuration }: { initialScale: number, animationDuration: number}) => ({
    position: 'fixed',
    opacity: 0,
    transitionTimingFunction: animationEasing.spring,
    transitionDuration: `${animationDuration}ms`,
    transitionProperty: 'opacity, transform',
    transform: `scale(${initialScale}) translateY(-1px)`,
    '&[data-state="entering"], &[data-state="entered"]': {
        opacity: 1,
        visibility: 'visible',
        transform: `scale(1)`
    },
    '&[data-state="exiting"]': {
        opacity: 0,
        transform: 'scale(1)'
    }
})

type State = {
    left: number
    top: number
    transformOrigin: string
}

export const defaultProps = {
    position: 'bottom',
    bodyOffset: 6,
    targetOffset: 6,
    initialScale: 0.9,
    animationDuration: 300,
    innerRef: () => { },
    onOpenComplete: () => { },
    onCloseComplete: () => { }
}

export class Positioner extends React.Component<PositionerProps, State> {
    latestAnimationFrame!: number
    targetRef: React.RefObject<HTMLElement>
    positionerRef: React.RefObject<HTMLElement>

    static defaultProps: Partial<PositionerProps> = {
        position: 'bottom',
        bodyOffset: 6,
        targetOffset: 6,
        initialScale: 0.9,
        animationDuration: 300,
        innerRef: () => {},
        onOpenComplete: () => {},
        onCloseComplete: () => {}
    }

    constructor(props: PositionerProps, context: any) {
        super(props, context)

        this.targetRef = React.createRef<HTMLElement>()
        this.positionerRef = React.createRef<HTMLElement>()
        this.state = initialState()
    }

    componentWillUnmount() {
        if (this.latestAnimationFrame) {
            cancelAnimationFrame(this.latestAnimationFrame)
        }
    }

    getTargetRef = (ref: React.RefObject<HTMLElement>): React.RefObject<HTMLElement> => {
        this.targetRef = ref
        return this.targetRef
    }

    getRef = (ref: React.RefObject<HTMLElement>) => {
        this.positionerRef = ref
        if (this.props.innerRef) {
            this.props.innerRef(ref)
        }
    }

    handleEnter = () => {
        this.update()
    }

    update = (prevHeight: number = 0, prevWidth: number = 0) => {
        if (!this.props.isShown || !this.targetRef.current || !this.positionerRef.current) return

        const targetRect = this.targetRef.current.getBoundingClientRect()
        const hasEntered =
            this.positionerRef.current.getAttribute('data-state') === 'entered'

        const viewportHeight = document.documentElement.clientHeight
        const viewportWidth = document.documentElement.clientWidth

        let height: number
        let width: number

        if (hasEntered) {
            // Only when the animation is done should we opt-in to `getBoundingClientRect`
            const positionerRect = this.positionerRef.current.getBoundingClientRect()

            // https://github.com/segmentio/evergreen/issues/255
            // We need to ceil the width and height to prevent jitter when
            // the window is zoomed (when `window.devicePixelRatio` is not an integer)
            height = Math.round(positionerRect.height)
            width = Math.round(positionerRect.width)
        } else {
            // When the animation is in flight use `offsetWidth/Height` which
            // does not calculate the `transform` property as part of its result.
            // There is still change on jitter during the animation (although unoticable)
            // When the browser is zoomed in — we fix this with `Math.max`.
            height = Math.max(this.positionerRef.current.offsetHeight, prevHeight)
            width = Math.max(this.positionerRef.current.offsetWidth, prevWidth)
        }

        const { rect, transformOrigin } = getPosition({
            position: this.props.position ?? 'bottom',
            targetRect,
            targetOffset: this.props.targetOffset ?? 6,
            dimensions: {
                height,
                width
            },
            viewport: {
                width: viewportWidth,
                height: viewportHeight
            },
            viewportOffset: this.props.bodyOffset ?? 6
        })

        this.setState(
            {
                left: rect.left,
                top: rect.top,
                transformOrigin
            },
            () => {
                this.latestAnimationFrame = requestAnimationFrame(() => {
                    this.update(height, width)
                })
            }
        )
    }

    handleExited = () => {
        this.setState(
            () => {
                return {
                    ...initialState()
                }
            },
            () => {
                if (this.props.onCloseComplete) {
                    this.props.onCloseComplete()
                }
            }
        )
    }

    render() {
        const {
            target,
            isShown,
            children,
            initialScale,
            targetOffset,
            animationDuration
        } = this.props

        const { left, top, transformOrigin } = this.state

        return (
            <Stack>
                {zIndex => {
                    return (
                        <React.Fragment>
                            {/**
                             * {target({ getRef: this.getTargetRef, isShown })}
                             */}

                            <Transition
                                appear
                                in={isShown}
                                timeout={animationDuration}
                                onEnter={this.handleEnter}
                                onEntered={this.props.onOpenComplete}
                                onExited={this.handleExited}
                                unmountOnExit
                            >
                                {(state) => (
                                    <Portal>
                                        {children({
                                            top,
                                            left,
                                            state,
                                            zIndex,
                                            css: getCSS({
                                                initialScale,
                                                animationDuration
                                            }),
                                            style: {
                                                transformOrigin,
                                                left,
                                                top,
                                                zIndex
                                            },
                                            getRef: this.getRef,
                                            animationDuration
                                        })}
                                    </Portal>
                                )}
                            </Transition>
                            
                        </React.Fragment>
                    )
                }}
            </Stack>
        )
    }
}