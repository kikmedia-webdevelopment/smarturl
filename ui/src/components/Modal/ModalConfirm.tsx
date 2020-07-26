import React from 'react'
import { Modal } from './Modal';
import { Button } from 'components/Button';

interface Props {
    title?: string
    isShown: boolean
    onClose?: (event: React.MouseEvent<Element, MouseEvent> | React.KeyboardEvent<Element>) => void
    onAfterOpen?: ReactModal.OnAfterOpenCallback
    shouldCloseOnEscapePress?: boolean
    shouldCloseOnOverlayClick?: boolean
    testId?: string
    confirmLabel?: string | false
    isConfirmLoading?: boolean
    onConfirm: () => void
}


export class ModalConfirm extends React.Component<Props> {
    static defaultProps = {
        testId: 'jk-ui-modal-confirm'
    }

    render() {
        const {
            title,
            isShown,
            onClose,
            testId,
            shouldCloseOnOverlayClick,
            shouldCloseOnEscapePress,
            children,
            confirmLabel,
            onConfirm
        } = this.props
        return (
            <Modal
                isShown={isShown}
                testId={testId}
                onClose={onClose}
                shouldCloseOnOverlayClick={shouldCloseOnOverlayClick}
                shouldCloseOnEscapePress={shouldCloseOnEscapePress}
            >
                <Modal.Header label={title ?? ''} />
                <Modal.Content>
                    {children}
                </Modal.Content>
                <Modal.Controls>
                    <Button
                        onClick={() => onConfirm()}
                    >
                        {confirmLabel ?? 'Bestätigen'}
                    </Button>
                </Modal.Controls>
            </Modal>
        )
    }
}