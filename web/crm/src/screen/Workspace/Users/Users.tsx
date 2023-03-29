import React from "react"

import { WorkspaceScreenUsersFragment } from "../useWorkspace"

import { Avatar, Box, Button, Tip } from "grommet"

export default function Users({ users }: { users: WorkspaceScreenUsersFragment }) {
    switch (users.list.__typename) {
        case "WorkspaceUsersList":
            return (
                <Box direction="row">
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
                                            {user.profile.abbr || "ПО"}
                                        </Avatar>
                                    }
                                />
                            </Tip>
                        )
                    })}
                </Box>
            )
        default:
            return null
    }
}