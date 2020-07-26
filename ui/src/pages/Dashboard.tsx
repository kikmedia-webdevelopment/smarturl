import React from 'react';
import { statsService } from 'services';
import { Stats } from 'models/stats';
import { Loading } from 'components';

interface Props {}
interface State {
    loading: boolean
    stats: Stats
}

export class Dashboard extends React.Component<Props, State> {

    constructor(props: Props) {
        super(props)

        this.state = {
            loading: true,
            stats: {
                entries: 0,
                visits: 0
            }
        }
    }

    componentDidMount() {
        statsService.getStats()
            .then(stats => {
                this.setState({
                    stats,
                })
            }).finally(() => this.setState({ loading: false }))
    }

    render() {
        const { stats, loading } = this.state

        if (loading) {
            return <Loading />
        }
        return (
            <React.Fragment>
                <header className="bg-white shadow">
                    <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex items-center justify-between">
                        <h1 className="text-3xl font-bold leading-tight text-gray-900">Dashboard</h1>
                        <div>
                        </div>
                    </div>
                </header>
                <div className="py-6">
                    <div className="max-w-7xl mx-auto sm:px-6 lg:px-8">
                        <div className="flex flex-wrap -mx-6">
                            <div className="w-full px-6 sm:w-1/2 xl:w-1/3">
                                {/** card start */}
                                <div className="flex items-center px-5 py-6 shadow rounded-md bg-white">
                                    <div className="p-3 rounded-full bg-blue-500 bg-opacity-75">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="44" height="44" viewBox="0 0 24 24" strokeWidth="1.5" stroke="#FFFFFF" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                            <path stroke="none" d="M0 0h24v24H0z" />
                                            <path d="M11 7h-5a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-5" />
                                            <line x1="10" y1="14" x2="20" y2="4" />
                                            <polyline points="15 4 20 4 20 9" />
                                        </svg>
                                    </div>
                                    <div className="mx-5">
                                        <h4 className="text-2xl font-semibold text-gray-700">{stats.entries}</h4>
                                        <div className="text-gray-500">Links</div>
                                    </div>
                                </div>
                                {/** card end */}
                            </div>
                            <div className="w-full px-6 sm:w-1/2 xl:w-1/3">
                                {/** card start */}
                                <div className="flex items-center px-5 py-6 shadow rounded-md bg-white">
                                    <div className="p-3 rounded-full bg-indigo-600 bg-opacity-75">
                                        <svg xmlns="http://www.w3.org/2000/svg" width="44" height="44" viewBox="0 0 24 24" strokeWidth="1.5" stroke="#FFFFFF" fill="none" strokeLinecap="round" strokeLinejoin="round">
                                            <path stroke="none" d="M0 0h24v24H0z" />
                                            <line x1="3" y1="12" x2="6" y2="12" />
                                            <line x1="12" y1="3" x2="12" y2="6" />
                                            <line x1="7.8" y1="7.8" x2="5.6" y2="5.6" />
                                            <line x1="16.2" y1="7.8" x2="18.4" y2="5.6" />
                                            <line x1="7.8" y1="16.2" x2="5.6" y2="18.4" />
                                            <path d="M12 12l9 3l-4 2l-2 4l-3 -9" />
                                        </svg>
                                    </div>
                                    <div className="mx-5">
                                        <h4 className="text-2xl font-semibold text-gray-700">{stats.visits}</h4>
                                        <div className="text-gray-500">Total Clicks</div>
                                    </div>
                                </div>
                                {/** card end */}
                            </div>
                        </div>
                    </div>
                    
                </div>
            </React.Fragment>
        )
    }
}