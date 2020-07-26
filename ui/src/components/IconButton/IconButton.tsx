import React from 'react'
import { IconName, Icon } from 'components/Icon/Icon'
import cn from 'classnames'

type Props = {
    label: string
    iconName: IconName
    disabled?: boolean
    variant?: 'danger' | 'default'
} & React.HTMLAttributes<HTMLButtonElement>

export class IconButton extends React.Component<Props> {
    render() {
        const {
            iconName,
            disabled,
            variant,
            ...rest
        } = this.props

        return (
            <button
                {...rest}
                className={
                    cn('inline-block cursor-pointer border-none p-0 m-0 bg-transparent', {
                        'text-red-600': variant === 'danger'
                    })
                }
                type="button"
                disabled={disabled}
            >
                <Icon
                    iconName={iconName}
                />
            </button>
        )
    }
}