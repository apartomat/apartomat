import React from "react";

import UserEmail from "./UserEmail";

import { useAuthContext } from "../auth/useAuthContext";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

import Box from "@material-ui/core/Box";
import Avatar from "@material-ui/core/Avatar";
import AvatarGroup from '@material-ui/lab/AvatarGroup';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
// import useWorkspace from "./useWorkspace";

import { useWorkspace, WorkspaceUsersResult, WorkspaceProjectsListResult } from "./useWorkspace";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    title: {
        flexGrow: 1,
    },
    workspace: {
        flexGrow: 1,
        display: "flex",
        alignItems: "center"
    },
    appBarSpacer: theme.mixins.toolbar,
    content: {
        flexGrow: 1,
        overflow: "auto"
      },
  }),
);

export function Workspace () {
    const classes = useStyles();
    const { user, concreteUser: { defaultWorkspaceId } } = useAuthContext();
    const { data, loading, error } = useWorkspace(defaultWorkspaceId);


    if (loading) {
        return (
            <Box>
                <p>Loading workspace...</p>
            </Box>
        );
    }

    if (error) {
        return (
            <Box>
                <h1>Errpr</h1>
                <p>Can't get workspace: {error}</p>
            </Box>
        );
    }

    switch (data?.workspace.__typename) {
        case "Workspace":
            const { workspace } = data;
            return (
                <div>
                    <AppBar position="fixed" color="transparent" elevation={0}>
                        <Toolbar>
                            <Typography variant="h6" component="h1" noWrap className={classes.title}>apartomat</Typography>
                            <UserEmail user={user}/>
                        </Toolbar>
                    </AppBar>
                    <main className={classes.content}>
                        <Typography variant="h4" component="h1">{workspace.name}</Typography>
                        <WorkspaceUsers users={workspace.users} />
                        <Projects projects={workspace.projects.list} />
                    </main>
                </div>
            );
        default:
            return (
                <div>{data?.workspace.__typename} case in not handled yet</div>
            );
    }
}

function WorkspaceUsers({ users }: {users: WorkspaceUsersResult}) {
    switch (users.__typename) {
        case "WorkspaceUsers":
            return (
                <AvatarGroup max={2}>
                    {users.items.map(user => <Avatar key={user.id} src={user.profile.gravatar.url} alt={user.profile.email}/>)}
                </AvatarGroup>
            )
        default:
            return <div>n/a</div>
    }
}

function Projects({ projects }: { projects: WorkspaceProjectsListResult }) {
    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            return (
                <ul>
                    {projects.items.map(project => <li key={project.id}>{project.name}</li>)}
                </ul>
            )
        default:
            return <div>n/a</div>
    }
}


export default Workspace;