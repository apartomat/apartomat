import { Avatar, AvatarExtendedProps } from "grommet"

import { useAuthContext, UserContextStatus } from "shared/context/auth/context"

export function UserAvatar(props: AvatarExtendedProps) {
    const { user } = useAuthContext()

    switch (user.status) {
        case UserContextStatus.LOGGED:
            return <Avatar src={user.avatar} {...props} />
        default:
            return null
    }
}
