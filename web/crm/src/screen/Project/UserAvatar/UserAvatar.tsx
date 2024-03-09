import { UserContext, UserContextStatus } from "context/auth/useAuthContext"

import { Avatar, AvatarExtendedProps } from "grommet"

export function UserAvatar ({ user }: {user: UserContext} & AvatarExtendedProps) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return (
                <Avatar src={user.avatar} />
            )
        default:
            return null
    }
}

export default UserAvatar
