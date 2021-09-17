import React, { StrictMode } from "react";
import { Grommet } from "grommet";
import ReactDOM from "react-dom";
import {BrowserRouter as Router, Switch, Route} from "react-router-dom";

import AuthProvider from "./common/context/auth/AuthProvider/AuthProvider";
import PrivateRoute from "./common/context/auth/PrivateRoute/PrivateRoute";
import RedirectToDefaultWorkspace from "./common/context/auth/RedirectToDefaultWorkspace/RedirectToDefaultWorkspace";

import Index from "./common/ui/Index"
import Login from "./login/Login";
import Logout from "./logout/Logout";
import Confirm from "./confirm/Confirm";
import Workspace from './workspace/Workspace';
import Project from './project/Project';
import Files from './files/Files';
import Spec from './spec/Spec';

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

const httpLink = createUploadLink({ uri: "http://localhost:8010/graphql" });

const link = authLink.concat(httpLink);

const apolloClient = new ApolloClient({ cache: new InMemoryCache(), link });

const theme = {
    global: {
        font: {
            family: 'Roboto',
            size: '18px',
            height: '20px',
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
                            <Route exact path="/ui-kit">
                                <Index/>
                            </Route>
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