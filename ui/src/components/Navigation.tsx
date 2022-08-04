import React, { useState } from 'react';
import { NavigationLink } from './NavigationLink';
import { User } from 'models/user';
import { Menu } from './Menu';
import { useHistory } from 'react-router-dom';
import { authenticationService } from 'services';

interface Props {
    user: User
}

interface State {
    userDropdown: boolean
}

export const Navigation: React.FC<Props> = ({
    user
}) => {
    const { push } = useHistory()
    const [ userDropdown, setUserDropdown ] = useState<boolean>(false)
    
    const toggleUserDropdown = () => {
        setUserDropdown(state => !state)
    }

    const onProfileClick = () => {
        setUserDropdown(false)
        push('/admin/profile')
    }

    const SignOut = () => {
        authenticationService.logout()
    }

    return (
        <div className="bg-gray-300">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex items-center justify-between h-16">
                    <div className="flex items-center">
                        <div className="flex-shrink-0">
                            <svg id="smart-logo" className="w-8 h-8" data-name="Ebene 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 121.93 186.12">
                                <defs>
                                    <style>.cls-1,.cls-3{fill:#1a171b;}.cls-2{fill:#005ea8;}.cls-3{fill-rule:evenodd;}</style>
                                </defs>
                                <title>Smart Testsolutions</title>
                                <path class="cls-1" d="M77.66,26.2a13,13,0,1,1,13,13,13,13,0,0,1-13-13"/>
                                <path class="cls-2" d="M11.78,74A52.79,52.79,0,0,1,7.32,62.7c3.72.67,7.85,1.42,10.68,5.39l.7-.54-1.27-3.11c7.46,4,16.72,12.17,24.78,15.51A51.32,51.32,0,0,0,61.67,85V0H.91V118.17H61.67v-9.68a49.14,49.14,0,0,1-8.59-1.34C34.35,103,22.35,89.44,11.78,74ZM32.42,14l13.9,24.08H18.51Z"/>
                                <path class="cls-2" d="M115.69,56.15c-1.46-.47-2.66-2.83-4.7-1.53-1.08,7.18-5.88,13.68-11.77,18A54.13,54.13,0,0,1,61.67,85v23.47c19.73,1.36,40.5-8.65,52.51-24-.29-3.5,3.24-6.12,3.81-10.05C115.79,68.55,118.81,62,115.69,56.15Z"/>
                                <path class="cls-1" d="M16.18,133.33a11,11,0,0,0-7.33-3,7.81,7.81,0,0,0-7.82,8c0,4.16,2.52,5.9,6,7.32,2.1.82,4.66,1.6,4.66,4.38A3.58,3.58,0,0,1,8.21,154C6,154,4.09,152.11,3.38,150L0,153.35c1.85,3.1,4.55,5.23,8.25,5.23,5.48,0,8.82-3.81,8.82-9.17,0-9-10.64-7.11-10.64-11.63a3.19,3.19,0,0,1,3.28-3,4.83,4.83,0,0,1,4,2.21Z"/>
                                <polygon class="cls-1" points="45.7 157.87 50.97 157.87 46.95 130.34 43.18 130.34 36.85 148.02 30.23 130.34 26.54 130.34 22.02 157.87 27.29 157.87 29.67 141.23 29.74 141.23 35.6 158.22 37.42 158.22 43.5 141.23 43.57 141.23 45.7 157.87"/>
                                <path class="cls-3" d="M71.91,152.22,74,157.87h5.58L69.39,130.34H65.33L54.87,157.87h5.51l2.21-5.65Zm-1.46-4.27H64.09l3-9.75h.07Z"/>
                                <path class="cls-3" d="M95.5,146.56c4.22-.75,5.47-4.73,5.47-8.14,0-5.73-3.66-8.08-9-8.08H84.94v27.53h5.22V147.06h.08l6.51,10.81h6.33Zm-5.34-11.95h.72c2.77,0,4.88.64,4.88,3.95s-2,4.12-4.91,4.12h-.68Z"/>
                                <polygon class="cls-1" points="116.53 134.89 121.93 134.89 121.93 130.34 105.89 130.34 105.89 134.89 111.3 134.89 111.3 157.87 116.53 157.87 116.53 134.89"/>
                                <path class="cls-1" d="M5.66,181.59H3.38v-9.71H1v-2H8v2H5.66Z"/>
                                <path class="cls-1" d="M11.91,171.89v2.57h3.63v2H11.91v3.16h3.77v2h-6V169.9h6v2Z"/>
                                <path class="cls-1" d="M23.47,172.48a2.11,2.11,0,0,0-1.75-1,1.39,1.39,0,0,0-1.43,1.32c0,2,4.63,1.15,4.63,5.07a3.69,3.69,0,0,1-3.84,4,4.14,4.14,0,0,1-3.6-2.28L19,178.18a2.4,2.4,0,0,0,2.11,1.71,1.56,1.56,0,0,0,1.5-1.69c0-1.21-1.12-1.55-2-1.91-1.5-.62-2.61-1.38-2.61-3.19a3.4,3.4,0,0,1,3.41-3.5,4.8,4.8,0,0,1,3.19,1.3Z"/>
                                <path class="cls-1" d="M30.7,181.59H28.42v-9.71H26.06v-2h7v2H30.7Z"/>
                                <path class="cls-2" d="M40.17,172.48a2.11,2.11,0,0,0-1.75-1A1.39,1.39,0,0,0,37,172.83c0,2,4.63,1.15,4.63,5.07a3.69,3.69,0,0,1-3.84,4,4.14,4.14,0,0,1-3.6-2.28l1.47-1.44a2.4,2.4,0,0,0,2.11,1.71,1.56,1.56,0,0,0,1.5-1.69c0-1.21-1.12-1.55-2-1.91-1.5-.62-2.61-1.38-2.61-3.19a3.4,3.4,0,0,1,3.41-3.5,4.8,4.8,0,0,1,3.2,1.3Z"/>
                                <path class="cls-2" d="M43,175.73a5.8,5.8,0,1,1,5.78,6.17A5.9,5.9,0,0,1,43,175.73Zm2.36-.09c0,1.77,1.49,4,3.43,4s3.43-2.25,3.43-4a3.68,3.68,0,0,0-3.43-3.8A3.68,3.68,0,0,0,45.39,175.64Z"/>
                                <path class="cls-2" d="M58.77,179.61H62v2h-5.5V169.9h2.28Z"/>
                                <path class="cls-2" d="M63.43,169.9h2.28v6.54c0,1.44.12,3.47,2.06,3.47s2.06-2,2.06-3.47V169.9h2.28v7c0,2.85-1.21,5-4.34,5s-4.34-2.15-4.34-5Z"/>
                                <path class="cls-2" d="M78.3,181.59H76v-9.71H73.66v-2h7v2H78.3Z"/>
                                <path class="cls-2" d="M84.54,181.59H82.26V169.9h2.28Z"/>
                                <path class="cls-2" d="M86.42,175.73a5.79,5.79,0,1,1,5.78,6.17A5.9,5.9,0,0,1,86.42,175.73Zm2.36-.09c0,1.77,1.49,4,3.43,4s3.43-2.25,3.43-4a3.44,3.44,0,1,0-6.85,0Z"/>
                                <path class="cls-2" d="M99.87,169.59h1.64l6.15,8.17h0V169.9H110v11.92h-1.64l-6.15-8.17h0v7.94H99.87Z"/>
                                <path class="cls-2" d="M117.91,172.48a2.11,2.11,0,0,0-1.75-1,1.39,1.39,0,0,0-1.43,1.32c0,2,4.63,1.15,4.63,5.07a3.69,3.69,0,0,1-3.84,4,4.14,4.14,0,0,1-3.6-2.28l1.47-1.44a2.4,2.4,0,0,0,2.11,1.71,1.56,1.56,0,0,0,1.5-1.69c0-1.21-1.12-1.55-2-1.91-1.5-.62-2.6-1.38-2.6-3.19a3.4,3.4,0,0,1,3.41-3.5,4.79,4.79,0,0,1,3.19,1.3Z"/>
                            </svg>
                        </div>
                        <div className="block">
                            <div className="ml-10 flex items-baseline">
                                <NavigationLink
                                    exact
                                    to="/admin/dashboard"
                                >Home</NavigationLink>
                                <NavigationLink
                                    exact
                                    to="/admin/links"
                                >Links</NavigationLink>
                            </div>
                        </div>
                    </div>
                    <div className="block">
                        <div className="ml-4 flex items-center md:ml-6">
                            {/** profile dropdown */}
                            <div className="ml-3 relative">
                                <button
                                    className="max-w-xs flex items-center text-sm rounded-full text-black focus:outline-none focus:shadow-solid"
                                    id="user-menu"
                                    aria-label="User Menu"
                                    aria-haspopup="true"
                                    onClick={() => toggleUserDropdown()}
                                >
                                    {user.email}
                                </button>

                                {userDropdown && (
                                    <div
                                        className="origin-top-right absolute right-0 mt-0 mt-2 w-48 rounded-md shadow-lg"
                                    >
                                        <Menu>
                                            <Menu.Group title="group">
                                                <Menu.Item
                                                    onClick={() => onProfileClick()}
                                                >
                                                    Profil
                                                </Menu.Item>
                                            </Menu.Group>
                                            <Menu.Group>
                                                <Menu.Item
                                                    onClick={() => SignOut()}
                                                >
                                                    Sign out
                                                    </Menu.Item>
                                            </Menu.Group>
                                        </Menu>
                                    </div>

                                )}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
