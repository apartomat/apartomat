import { UserContext, UserContextStatus } from "shared/context/auth/context"

import { Avatar, AvatarExtendedProps } from "grommet"

export function UserAvatar({ user }: { user: UserContext } & AvatarExtendedProps) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return <Avatar src={user.avatar} />
        default:
            return null
    }
}

export default UserAvatar
