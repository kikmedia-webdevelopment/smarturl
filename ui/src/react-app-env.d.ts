/// <reference types="react-scripts" />

interface Config {
    apiurl: string | null
    token: string | null
}

declare global {
    interface Window {
        config: Config
    }
}

interface Window {
    config: Config
}