import React, { useEffect } from "react";
import {Route, Redirect, RouteProps } from "react-router-dom";

import useAuthContext, { UserContextStatus } from '../useAuthContext';

import Box from "@material-ui/core/Box";

function PrivateRoute({ children, ...rest }: RouteProps) {
    const { user, check, error } = useAuthContext();

    useEffect(() => {
        if (user.status === UserContextStatus.UNDEFINED) {
            check();
        }
    }, [user, check]);

    if (error !== undefined) {
        return (
            <Box>
                <h1>Error</h1>
                <p>Can't check profile: {error}</p>
            </Box>
        );
    }

    return (
        <Route {...rest} render={({ location }) => {
            switch (user.status) {
                case UserContextStatus.UNDEFINED:
                case UserContextStatus.CHEKING:
                    return (
                        <Box>
                            <p>Checking...</p>
                        </Box>
                    );
                case UserContextStatus.SERVER_ERROR:
                    return (
                        <Box>
                            <h1>Error</h1>
                            <p>Can't check profile. Please refresh the page</p>
                        </Box>
                    );
                case UserContextStatus.LOGGED:
                    return children;
                default:
                    return <Redirect to={{ pathname: "/login", state: { from : location }}}/>;
            }
        }} />
    );
}

export default PrivateRoute;