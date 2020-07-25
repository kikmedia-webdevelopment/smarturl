import React from 'react'

interface Props extends React.HTMLProps<HTMLButtonElement> {}

export class Button extends React.Component<Props, {}> {
    render() {
        const { type, ...props} = this.props
        return (
            <button
                {...props}
                type={type as 'submit' | 'button' | 'reset' | undefined }
                className="inline-flex items-center justify-center w-full px-5 py-3 text-base font-medium leading-6 text-white transition duration-150 ease-in-out bg-secondary border border-transparent rounded-lg hover:bg-brand"
                
            />
        )
    }
}