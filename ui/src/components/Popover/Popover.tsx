import React, { useEffect } from 'react'
import { usePopper } from 'react-popper'
import { PopupProps } from './types'
import { Portal } from 'components/Portal'


export const Popover: React.FC<PopupProps> = ({
    id,
    trigger,
    isShown,
    placement,
    children
}) => {
    const [triggerRef, setTriggerRef] = React.useState<HTMLElement | null>(null)
    const [popperRef, setPopperRef] = React.useState<HTMLDivElement | null>(null)
    const { styles } = usePopper(triggerRef, popperRef, {
        placement: placement ?? 'auto'
    })
    const [isOpen, setOpen] = React.useState<boolean>(isShown ?? false)

    useEffect(() => {
        const handleClickOutside = (ev: MouseEvent) => {
            if (triggerRef && triggerRef.contains(ev.target as Element)) {
                return
            }
            if (popperRef && !popperRef.contains(ev.target as Element)) {
                if (isOpen) {
                    setOpen(false)
                }
            }
        }
        if (isOpen) {
            document.addEventListener('mousedown', handleClickOutside)
        }
        return () => {
            document.removeEventListener('mousedown', handleClickOutside)
        }
    }, [isOpen, popperRef, triggerRef])

    const Toggle = () => {
        setOpen(open => !open)
    }

    return (
        <>
            {React.cloneElement(trigger, {
                ref: setTriggerRef,
                'aria-controls': id,
                'aria-expanded': isOpen,
                'aria-haspopup': true,
                onClick: () => Toggle()
            })}
            {isOpen && (
                <Portal>
                    <div
                        className="shadow bg-white rounded"
                        ref={setPopperRef}
                        style={styles.popper}
                        
                    >
                        {children}
                    </div>
                </Portal>
            )}
        </>
    )
}

Popover.displayName = 'Popover'