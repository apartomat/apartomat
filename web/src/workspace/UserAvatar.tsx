import React from "react";

import { UserContext, UserContextStatus } from "../common/context/auth/useAuthContext";

import Avatar from "../common/ui/Avatar";

export function UserAvatar ({ user, className }: {user: UserContext, className?: string}) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return (
                <Avatar src={user.avatar} className={className} />
            );
        default:
            return null;
    }
}

export default UserAvatar;