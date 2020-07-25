import React from 'react'
import ReactModal from 'react-modal'
import { ModalHeader } from './ModalHeader'
import { ModalContent } from './ModalContent'
import { ModalControls } from './ModalControls'

export type ModalProps = {
    isShown: boolean
    onClose?: (event: React.MouseEvent<Element, MouseEvent> | React.KeyboardEvent<Element>) => void
    onAfterOpen?: ReactModal.OnAfterOpenCallback
    shouldCloseOnEscapePress?: boolean
    shouldCloseOnOverlayClick?: boolean
    testId?: string
} & typeof defaultProps

const defaultProps = {
    shouldCloseOnEscapePress: true,
    shouldCloseOnOverlayClick: true,
    testId: 'jk-ui-modal',
}

export class Modal extends React.Component<ModalProps> {
    static defaultProps = defaultProps
    static Header = ModalHeader
    static Content = ModalContent
    static Controls = ModalControls

    render() {
        const {
            isShown,
            onClose,
            onAfterOpen,
            shouldCloseOnEscapePress,
            shouldCloseOnOverlayClick,
            children
        } = this.props
        return (
            <ReactModal
                ariaHideApp={false}
                onRequestClose={onClose}
                isOpen={isShown}
                onAfterOpen={onAfterOpen}
                shouldCloseOnEsc={shouldCloseOnEscapePress}
                shouldCloseOnOverlayClick={shouldCloseOnOverlayClick}
                closeTimeoutMS={100}
                htmlOpenClassName="Modal__html--open"
                bodyOpenClassName="Modal__body--open"
                portalClassName="block"
                overlayClassName={{
                    base: "z-40 bg-modal-bg  flex items-center justify-center inset-0 opacity-0 delay-100 transition-opacity ease-in-out fixed overflow-y-auto text-center",
                    afterOpen: "opacity-100",
                    beforeClose: "opacity-0"
                }}
                className={{
                    base: "relative z-50  relative p-0 inline-block my-0 text-left outline-none top-0",
                    afterOpen: "scale-100 opacity-100",
                    beforeClose: "opacity-50 scale-75"
                }}
            >
                <div
                    className="bg-white rounded shadow max-w-screen-md flex flex-col w-full"
                    style={{ width: 700 }}
                >
                    {children}
                    
                </div>
            </ReactModal>
        )
    }
}