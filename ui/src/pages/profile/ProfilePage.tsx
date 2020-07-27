import React from 'react'
import { User } from 'models/user'
import { authenticationService, userService } from 'services'
import { Button, Modal, Form, FormLabel, TextInput } from 'components'

type Props = {}
type State = {
    currentUser: User | null
    showPasswordModal: boolean
    password: {
        currentPassword: string 
        newPassword: string
        newPasswordConfirmed: string
    }
}

export class ProfilePage extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props)

        this.state = {
            currentUser: null,
            showPasswordModal: false,
            password: {
                currentPassword: '',
                newPassword: '',
                newPasswordConfirmed: '',
            }
        }
    }

    showPasswordModal() {
        this.setState(({
            showPasswordModal: true
        }))
    }

    hidePasswordModal() {
        this.setState(({
            showPasswordModal: false,
            password: {
                currentPassword: '',
                newPassword: '',
                newPasswordConfirmed: '',
            }
        }))
    }

    onSubmitPasswordChange() {
        const { password } = this.state
        userService.updatePassword(password)
            .finally(() => {
                this.setState(({
                    showPasswordModal: false,
                    password: {
                        currentPassword: '',
                        newPassword: '',
                        newPasswordConfirmed: '',
                    }
                }))
            })

    }

    componentDidMount() {
        authenticationService.currentUser.subscribe(x => this.setState({ currentUser: x }))
    }

    render() {
        const { currentUser } = this.state
        return (
            <React.Fragment>
                <header className="bg-white shadow">
                    <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex items-center justify-between">
                        <h1 className="text-3xl font-bold leading-tight text-gray-900">Profil</h1>
                        <div>
                            <Button
                                onClick={() => this.showPasswordModal()}
                            >
                                Change Password
                            </Button>
                        </div>
                    </div>
                </header>
                <div className="max-w-2xl p-8 mx-auto w-full">
                    <div
                        className="mt-4 bg-white rounded-lg shadow-md"
                    >
                        <div className="px-6 py-4">
                            <dl>
                                <dt className="font-sans block font-bold">
                                    E-Mail
                                </dt>
                                <dd>
                                    {currentUser?.email}
                                </dd>
                            </dl>
                        </div>
                </div>
                </div>
                <Modal
                    isShown={this.state.showPasswordModal}
                    onClose={() => this.hidePasswordModal()}
                >
                    <Modal.Header label="Change Password" />
                    <Modal.Content>
                        <Form>
                            <div>
                                <FormLabel htmlFor="currentPassword" required>
                                    Current Password
                                </FormLabel>
                                <TextInput
                                    id="currentPassword"
                                    name="currentPassword"
                                    type="password"
                                    value={this.state.password.currentPassword ?? ''}
                                    onChange={(e: React.FormEvent<HTMLInputElement>) => {
                                        const value = (e.target as HTMLInputElement).value
                                        this.setState((prevState) => ({
                                            password: {
                                                ...prevState.password,
                                                currentPassword: value
                                            }
                                        }))
                                    }}
                                    required
                                />
                            </div>

                            <div>
                                <FormLabel htmlFor="newPassword" required>
                                    New Password
                                </FormLabel>
                                <TextInput
                                    type="password"
                                    id="newPassword"
                                    name="newPassword"
                                    value={this.state.password.newPassword ?? ''}
                                    onChange={e => {
                                        const value = (e.target as HTMLInputElement).value
                                        this.setState((prevState) => ({
                                            password: {
                                                ...prevState.password,
                                                newPassword: value
                                            }
                                        }))
                                    }}
                                    required
                                />
                            </div>

                            <div>
                                <FormLabel htmlFor="newPasswordConfirmed" required>
                                    New Password Confirmed
                                </FormLabel>
                                <TextInput
                                    type="password"
                                    id="newPasswordConfirmed"
                                    name="newPasswordConfirmed"
                                    value={this.state.password.newPasswordConfirmed ?? ''}
                                    onChange={e => {
                                        const value = (e.target as HTMLInputElement).value
                                        this.setState((prevState) => ({
                                            password: {
                                                ...prevState.password,
                                                newPasswordConfirmed: value
                                            }
                                        }))
                                    }}
                                    required
                                />
                            </div>
                        </Form>
                    </Modal.Content>
                    <Modal.Controls>
                        <div className="grid gap-2 grid-flow-col">
                            <Button onClick={() => this.onSubmitPasswordChange()}>Save</Button>
                            <Button>Cancel</Button>
                        </div>
                    </Modal.Controls>
                </Modal>
            </React.Fragment>
            
        )
    }
}