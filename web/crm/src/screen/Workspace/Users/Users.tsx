import { useState } from "react"

import { WorkspaceScreenUsersFragment } from "../useWorkspace"
import { WorkspaceUserRoleDictionary } from "api/graphql"

import { Avatar, Box, BoxExtendedProps, Button, Tip } from "grommet"
import { Add } from "grommet-icons"

import Invite from "./Invite/Invite"


export default function Users({
    workspaceId,
    users,
    roles,
    onUserInviteSent,
    ...boxProps
}: {
    workspaceId: string,
    users: WorkspaceScreenUsersFragment,
    roles: WorkspaceUserRoleDictionary,
    onUserInviteSent?: (to: string) => void,
} & BoxExtendedProps) {
    const [ showInviteUserLayer, setShowInviteUserLayer ] = useState(false)

    return (
        <Box {...boxProps}>
            {users.list.__typename === "WorkspaceUsersList" && 
                <Box direction="row" gap="small">
                    {users.list.items.map(user => {
                        return (
                            <Tip content={user.profile.email} key={user.id}>
                                <Button
                                    plain
                                    icon={
                                        <Avatar
                                            key={user.id}
                                            src={user.profile.gravatar?.url}
                                            size="medium"
                                            background="light-1"
                                            border={{ color: "white", size: "small" }}
                                        >
                                            { user.profile.abbr || "NA" }
                                        </Avatar>
                                    }
                                />
                            </Tip>
                        )
                    })}

                    <Box justify="center">
                        <Button
                            color="brand"
                            label="Пригласить"
                            icon={<Add/>}
                            onClick={() => setShowInviteUserLayer(true)}
                        />
                    </Box>
                </Box>
            }

            {showInviteUserLayer &&
                <Invite
                    workspaceId={workspaceId}
                    roles={roles}
                    onClickClose={() => setShowInviteUserLayer(false)}
                    onInviteSent={(...args) => {
                        setShowInviteUserLayer(false)
                        onUserInviteSent && onUserInviteSent(...args)
                    }}
                />
            }
        </Box>
    )    
}

