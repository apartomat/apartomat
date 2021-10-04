import React from "react"

import { UserContext, UserContextStatus } from "../common/context/auth/useAuthContext"

import { Avatar } from "grommet"

export function UserAvatar ({ user, className }: {user: UserContext, className?: string}) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return (
                <Avatar src={user.avatar} />
            );
        default:
            return null;
    }
}

export default UserAvatar;