import React from 'react';
import { Table, Modal, Form, TextInput, FormLabel, Button, IconButton, ModalConfirm, Menu, Icon } from 'components';
import { linkService } from 'services/link.service';
import { Link } from 'models/link';
import { uid } from 'react-uid'
import { Popover } from 'components/Popover/Popover';
import config from 'config';

interface Props {}
interface State {
    loading: boolean
    links: Link[]
    editModal: boolean
    createModal: boolean
    showDeleteConfirm: boolean
    currentLink: Link | null
}

export class LinkList extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props)

        this.state = {
            loading: true,
            links: [],
            editModal: false,
            createModal: false,
            showDeleteConfirm: false,
            currentLink: null
        }
    }
    componentDidMount() {
        this.fetchAndSet()
    }

    fetchAndSet = async () => {
        const res = await linkService.list().finally(() => this.setState({ loading: false }))
        if (res) {
            this.setState({
                links: res
            })
        }
    }

    onCreateClick = () => {
        this.setState({
            createModal: true,
            currentLink: {
                id: '',
                url: '',
                VisitCount: 0
            }
        })
    }

    onCreateClose = () => {
        this.setState({
            currentLink: null,
            createModal: false
        })
    }

    onEditClick = (item: Link) => {
        this.setState({
            editModal: true,
            currentLink: item
        })
    }

    onCloseEdit = () => {
        this.setState({
            currentLink: null,
            editModal: false
        })
    }

    handleUrlUpdate = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value
        const { currentLink } = this.state
        if (currentLink) {
            currentLink.url = value
            this.setState({
                currentLink: currentLink
            })
        }
    }

    onSaveClick = async () => {
        const { currentLink } = this.state
        if (!currentLink) return

        this.setState({
            loading: true
        })
        linkService.create(currentLink)
            .then((res: Link) => {
                const { links } = this.state
                links.push(res)
                this.setState({
                    links: links
                })
            }).finally(() => this.setState({ loading: false, createModal: false }))
    }
    
    onUpdateClick = async () => {
        const { currentLink } = this.state
        if (!currentLink) return

        this.setState({
            loading: true
        })
        linkService.update(currentLink)
            .then((res: Link) => {
                const { links } = this.state
                const index = links.findIndex(l => l.id === res.id)
                if (index !== -1) {
                    links[index] = res
                    this.setState({
                        links: links,
                    })
                }
            }).finally(() => this.setState({ loading: false, editModal: false }))
    }

    onDeleteClose = () => {
        this.setState({
            currentLink: null,
            showDeleteConfirm: false,
        })
    }
    onDeleteConfirm = () => {
        const { currentLink: item } = this.state
        if (!item) {
            this.setState({
                showDeleteConfirm: false
            })
            return
        }

        this.setState({
            loading: true,
        })

        linkService.destroy(item)
            .then(() => {
                let { links } = this.state
                links = links.filter(l => l.id !== item.id)
                this.setState({
                    links: links
                })
            })
            .finally(() => this.setState({ loading: false, currentLink: null, showDeleteConfirm: false }))
    }

    onDeleteClick = (item: Link) => {
        this.setState({
            showDeleteConfirm: true,
            currentLink: Object.assign({}, item)
        })
    }

    onGetQRCodeClick = (item: Link) => {
        const code = encodeURIComponent(`${config.baseUrl}/${item.id}`)
        const url = `${config.baseUrl}/qrcode/${code}/256/1`
        if (window) {
            let win = window.open(url, '_blank')
            win?.focus()
        }
    }

    render() {
        const { links } = this.state
        return (
            <React.Fragment>
                <header className="bg-white shadow">
                    <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex items-center justify-between">
                        <h1 className="text-3xl font-bold leading-tight text-gray-900">Links</h1>
                        <div>
                            <Button onClick={this.onCreateClick}>Erstellen</Button>
                        </div>
                    </div>
                </header>
                <div className="py-6">
                    <div className="max-w-7xl mx-auto sm:px-6 lg:px-8">
                        <Table>
                            <Table.Head>
                                <Table.Row>
                                    <Table.HeadItem>
                                        Ziel
                                </Table.HeadItem>
                                    <Table.HeadItem>
                                        Code
                                </Table.HeadItem>
                                    <Table.HeadItem>
                                        VisitCount
                                </Table.HeadItem>
                                    <Table.HeadAction />
                                </Table.Row>
                            </Table.Head>
                            <Table.Body>
                                {links.map(link => (
                                    <Table.Row key={uid(link)}>
                                        <Table.Item>
                                            {link.url}
                                        </Table.Item>
                                        <Table.Item>
                                            {link.id}
                                        </Table.Item>
                                        <Table.Item>
                                            {link.VisitCount}
                                        </Table.Item>
                                        <Table.Item>
                                            <div className="grid gap-1 grid-flow-col">
                                                <Popover
                                                    trigger={
                                                        <span className="cursor-pointer select-none">...</span>
                                                    }
                                                >
                                                    <Menu>
                                                        <Menu.Group>
                                                            <Menu.Item
                                                                onClick={() => this.onGetQRCodeClick(link)}
                                                            >
                                                                QR Code Anzeigen
                                                            </Menu.Item>
                                                            <Menu.Item
                                                                onClick={() => this.onEditClick(link)}
                                                            >
                                                                Bearbeiten
                                                            </Menu.Item>
                                                        </Menu.Group>
                                                        <Menu.Group>
                                                            <Menu.Item
                                                                onClick={() => this.onDeleteClick(link)}
                                                            >
                                                                Löschen...
                                                            </Menu.Item>
                                                        </Menu.Group>
                                                    </Menu>
                                                </Popover>
                                            </div>

                                        </Table.Item>
                                    </Table.Row>
                                ))}
                            </Table.Body>
                        </Table>
                    </div>
                    {this.state.showDeleteConfirm && this.state.currentLink && (
                        <ModalConfirm
                            title="Bist du Sicher?"
                            isShown={this.state.showDeleteConfirm}
                            onClose={() => this.onDeleteClose()}
                            onConfirm={() => this.onDeleteConfirm()}
                        >
                            Du bist gerade dabei
                            `{this.state.currentLink.url}` mit dem Code
                            `{this.state.currentLink.id}` zu löschen
                        </ModalConfirm>
                    )}
                    {this.state.editModal && this.state.currentLink && (
                        <Modal
                            isShown={this.state.editModal}
                            onClose={() => this.onCloseEdit()}
                        >
                            <Modal.Header label="Bearbeiten" />
                            <Modal.Content>
                                <Form>
                                    <div>
                                        <FormLabel
                                            htmlFor="url"
                                            required
                                        >
                                            Ziel
                                    </FormLabel>
                                        <TextInput
                                            value={this.state.currentLink.url}
                                            name="url"
                                            onChange={this.handleUrlUpdate}
                                        />
                                    </div>
                                    <div>
                                        <FormLabel
                                            htmlFor="url"
                                        >
                                            Code
                                    </FormLabel>
                                        <TextInput
                                            value={this.state.currentLink.id}
                                            disabled
                                            name="id"
                                            onChange={() => { }}
                                        />
                                    </div>
                                </Form>

                            </Modal.Content>
                            <Modal.Controls>
                                <Button onClick={() => this.onUpdateClick()}>Speichern</Button>
                            </Modal.Controls>
                        </Modal>
                    )}
                    {this.state.createModal && this.state.currentLink && (
                        <Modal
                            isShown={this.state.createModal}
                            onClose={() => this.onCloseEdit()}
                        >
                            <Modal.Header label="Neuen Link erstellen" />
                            <Modal.Content>
                                <Form>
                                    <div>
                                        <FormLabel htmlFor="url" required>URL</FormLabel>
                                        <TextInput
                                            value={this.state.currentLink.url}
                                            name="url"
                                            onChange={this.handleUrlUpdate}
                                        />
                                    </div>
                                    <div>
                                        <FormLabel htmlFor="code" required>CODE</FormLabel>
                                        <TextInput
                                            value={this.state.currentLink.id}
                                            name="code"
                                            onChange={(e) => {
                                                const { currentLink } = this.state
                                                if (currentLink) {
                                                    this.setState({
                                                        currentLink: {
                                                            ...currentLink,
                                                            id: e.target.value
                                                        }
                                                    })
                                                }
                                            }}
                                        />
                                    </div>
                                </Form>
                            </Modal.Content>
                            <Modal.Controls>
                                <Button onClick={() => this.onSaveClick()}>Erstellen</Button>
                            </Modal.Controls>
                        </Modal>
                    )}
                </div>
            </React.Fragment>
        )
    }
}