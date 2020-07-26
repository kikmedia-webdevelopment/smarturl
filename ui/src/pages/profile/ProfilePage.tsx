import React from 'react'
import { User } from 'models/user'
import { authenticationService } from 'services'

type Props = {}
type State = {
    currentUser: User | null
}

export class ProfilePage extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props)

        this.state = {
            currentUser: null
        }
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
            </React.Fragment>
            
        )
    }
}