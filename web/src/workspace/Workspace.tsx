import React from "react";

import UserEmail from "./UserEmail";

import { useAuthContext } from "../auth/useAuthContext";

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

import Box from "@material-ui/core/Box";

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
// import useWorkspace from "./useWorkspace";

import { useWorkspace } from "./useWorkspace";

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
            return (
                <div>
                    <AppBar position="fixed" color="transparent" elevation={0}>
                        <Toolbar>
                            <Typography variant="h6" component="h1" noWrap className={classes.title}>apartomat</Typography>
                            <Box className={classes.workspace}>
                                <Typography variant="h4" component="h1">{data?.workspace.name}</Typography>
                            </Box>
                            <UserEmail user={user}/>
                        </Toolbar>
                    </AppBar>
                    <main className={classes.content}>
                    </main>
                </div>
            );
        default:
            return (
                <div>{data?.workspace.__typename} case in not handled yet</div>
            );
    }

    
}

export default Workspace;