import React, { StrictMode } from "react"
import { Grommet } from "grommet"
import ReactDOM from "react-dom"
import {BrowserRouter as Router, Switch, Route} from "react-router-dom"

import AuthProvider from "./common/context/auth/AuthProvider/AuthProvider"
import PrivateRoute from "./common/context/auth/PrivateRoute/PrivateRoute"
import RedirectToDefaultWorkspace from "./common/context/auth/RedirectToDefaultWorkspace/RedirectToDefaultWorkspace"

import Login from "screen/Login/Login";
import Logout from "./logout/Logout";
import Confirm from "screen/Confirm/Confirm";
import Workspace from './workspace/Workspace';
import Project from 'screen/Project/Project'
import Files from 'screen/Files/Files';
import Spec from 'screen/Spec/Spec';

import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { setContext } from '@apollo/client/link/context';
import { createUploadLink } from "apollo-upload-client";

const authLink = setContext((_, { headers }) => {
  const token = JSON.parse(localStorage.getItem("token") || `""`);
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const httpLink = createUploadLink({ uri: process.env.REACT_APP_APARTOMAT_API_URL });

const link = authLink.concat(httpLink);

const apolloClient = new ApolloClient({ cache: new InMemoryCache(), link });

const theme = {
    global: {
        font: {
            family: "Roboto",
            size: "18px",
            height: "20px",
        },
    },
}

ReactDOM.render(
    <Grommet theme={theme}>
        <StrictMode>
            <ApolloProvider client={apolloClient}>
                <AuthProvider>
                    <Router>
                        <Switch>
                            <PrivateRoute exact path="/">
                                <RedirectToDefaultWorkspace/>
                            </PrivateRoute>
                            <Route path="/login">
                                <Login/>
                            </Route>
                            <Route exact path="/logout">
                                <Logout/>
                            </Route>
                            <Route exact path="/confirm">
                                <Confirm/>
                            </Route>
                            <PrivateRoute exact path="/:id">
                                <Workspace/>
                            </PrivateRoute>
                            <PrivateRoute exact path="/p/:id">
                                <Project/>
                            </PrivateRoute>
                            <PrivateRoute exact path="/p/:projectId/files">
                                <Files/>
                            </PrivateRoute>
                            <PrivateRoute exact path="/p/:projectId/spec">
                                <Spec/>
                            </PrivateRoute>
                        </Switch>
                    </Router>
                </AuthProvider>
            </ApolloProvider>
        </StrictMode>
    </Grommet>,
    document.getElementById('root')
);