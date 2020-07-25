import React from 'react'

class TableHeadAction extends React.Component {
    render() {
        const { children } = this.props
        return (
            <th className="px-6 py-3 border-b border-gray-200 bg-gray-50">
                {children}
            </th>
        )
    }
}

class TableHeadItem extends React.Component {
    render() {
        const { children } = this.props
        return (
            <th className="px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
                {children}
            </th>
        )
    }
}

class TableRow extends React.Component {
    render() {
        const { children } = this.props
        return (
            <tr>
                {children}
            </tr>
        )
    }
}

class TableItem extends React.Component {
    render() {
        const { children } = this.props
        return (
            <td className="px-6 py-4 whitespace-no-wrap border-b border-gray-200">
                {children}
            </td>
        )
    }
}

class TableBody extends React.Component {
    render() {
        const { children } = this.props
        return (
            <tbody className="bg-white">
                {children}
            </tbody>
        )
    }
}

class TableHead extends React.Component {
    render() {
        const { children } = this.props
        return (
            <thead>
                {children}
            </thead>
        )
    }
}

export class Table extends React.Component {
    static Head = TableHead
    static Body = TableBody
    static Row = TableRow
    static Item = TableItem
    static HeadItem = TableHeadItem
    static HeadAction = TableHeadAction

    render() {
        const { children } = this.props
        return (
            <div className="flex flex-col">
                <div className="-my-2 py-2 overflow-x-auto sm:-mx-6 sm:px-6 lg:-mx-8 lg:px-8">
                    <div className="align-middle inline-block min-w-full shadow overflow-hidden sm:rounded-lg border-b border-gray-200">
                        <table className="min-w-full">
                            {children}
                        </table>
                    </div>
                </div>
            </div>
        )
    }
}