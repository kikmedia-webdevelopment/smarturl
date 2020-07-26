import React from 'react'
import cn from 'classnames'
import styles from './loading.module.css'

export const Loading: React.FC = () => {
    return (
        <div
            className="flex-1 flex items-center justify-center"
        >
            <div className="bg-white border py-2 px-5 rounded-lg flex items-center flex-col">
                <div className={cn('loader-dots block relative w-20 h-5 mt-2', styles['loader-dots'])}>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-brand"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-brand"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-brand"></div>
                    <div className="absolute top-0 mt-1 w-3 h-3 rounded-full bg-brand"></div>
                </div>
                
            </div>
        </div>
    )
}