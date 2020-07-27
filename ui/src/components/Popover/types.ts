import React from 'react'
import { Placement } from '@popperjs/core'

export type TriggerProps = {
    ref: React.Ref<HTMLElement>;
    'aria-controls'?: string;
    'aria-expanded': boolean;
    'aria-haspopup': boolean;
}

export type PopupProps = {
    /** HTML Id for testing etc */
    id?: string;
    /** Open State of the Dialog */
    isShown?: boolean
    /** Component used to anchor the popup to your content. Usually a button used to open the popup */
    trigger: React.ReactElement
    placement?: Placement
    /** Formatted like "0, 8px" — how far to offset the Popper from the Reference. Changes automatically based on the placement */
    offset?: number | string
    children: any
}