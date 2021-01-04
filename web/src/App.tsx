import React from "react";
import {BrowserRouter as Router, Switch, Route} from "react-router-dom";

import AuthProvider from "./auth/AuthProvider/AuthProvider";
import PrivateRoute from "./auth/PrivateRoute/PrivateRoute";
import RedirectToDefaultWorkspace from "./auth/RedirectToDefaultWorkspace/RedirectToDefaultWorkspace";
import Login from "./auth/Login/Login";
import Logout from "./auth/Logout/Logout";
import Confirm from "./auth/Confirm/Confirm";
import Workspace from './workspace/Workspace';

function App() {
    return (
        <AuthProvider>
            <Router>
                <Switch>
                    <PrivateRoute exact path="/">
                        <RedirectToDefaultWorkspace/>
                    </PrivateRoute>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <Route path="/logout">
                        <Logout/>
                    </Route>
                    <Route path="/confirm">
                        <Confirm/>
                    </Route>
                    <PrivateRoute path="/:id">
                        <Workspace/>
                    </PrivateRoute>
                </Switch>
            </Router>
        </AuthProvider>
    ); 
}

export default App;