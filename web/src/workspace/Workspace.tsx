import React, { useState } from 'react';

import { Clipboard, ProductList, Product } from './Clipboard/Clipboard';

import { useAuthContext, UserContext, UserContextStatus } from '../auth/useAuthContext';

import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Avatar from "@material-ui/core/Avatar";
import Container from "@material-ui/core/Container";
import Box from '@material-ui/core/Box';
import MoreVertIcon from '@material-ui/icons/MoreVert';
import IconButton from '@material-ui/core/IconButton';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';

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

    const { user } = useAuthContext();
    const [ list, setList ] = useState<Product[]>([]);

    function add(prod: Product) {
        setList(list.concat([prod]));
    }

    return (
        <div>
            <AppBar position="fixed" color="transparent" elevation={0}>
                <Toolbar>
                    <Typography variant="h6" component="h1" noWrap className={classes.title}>Apartomat</Typography>
                    <Box className={classes.workspace}>
                        <Typography variant="h4" component="h1">Workspace</Typography>
                    </Box>
                    
                    <UserEmail user={user}/>
                </Toolbar>
            </AppBar>
            <main className={classes.content}>
                <Container>
                </Container>
            </main>
        </div>
    );
}

function UserEmail ({ user }: {user: UserContext}) {
    switch (user.status) {
        case UserContextStatus.LOGGED:
            return (
                <Avatar src={user.avatar}/>
            );
        default:
            return null;
    }
}

export default Workspace;