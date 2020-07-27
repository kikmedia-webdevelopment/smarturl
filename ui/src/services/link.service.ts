import { Link } from 'models/link'
import api from './api.service'
import { AxiosResponse } from 'axios'

export const linkService = {
    list,
    update,
    create,
    destroy
}

async function destroy(link: Link): Promise<AxiosResponse<void>> {
    return api.delete<void>('/links', {
        data: link
    }).then((res) => Promise.resolve(res))
}

async function list() {
    return api.get<Link[]>('/links').then(res => res.data)
}

async function update(link: Link) {
    return api.patch<Link>('/links', link).then(res => res.data)
}

async function create(link: Link) {
    return api.post<Link>('/links', link).then(res => res.data)
}