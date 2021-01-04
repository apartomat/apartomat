import React from "react";

import { UserContext, UserContextStatus } from "../auth/useAuthContext";

import Avatar from "@material-ui/core/Avatar";

export function UserEmail ({ user }: {user: UserContext}) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return (
                <Avatar src={user.avatar}/>
            );
        default:
            return null;
    }
}

export default UserEmail;