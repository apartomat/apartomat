import React, { useEffect } from "react";
import {Route, Redirect, RouteProps } from "react-router-dom";

import useAuthContext, { UserContextStatus } from "../useAuthContext";

function PrivateRoute({ children, ...rest }: RouteProps) {
    const { user, check, error } = useAuthContext();

    useEffect(() => {
        if (user.status === UserContextStatus.UNDEFINED) {
            check();
        }
    }, [user, check]);

    if (error !== undefined) {
        return (
            <div>
                <h1>Error</h1>
                <p>Can't check profile: {error}</p>
            </div>
        );
    }

    return (
        <Route {...rest} render={({ location }) => {
            switch (user.status) {
                case UserContextStatus.UNDEFINED:
                case UserContextStatus.CHEKING:
                    return (
                        <div>
                            <p>Checking...</p>
                        </div>
                    );
                case UserContextStatus.SERVER_ERROR:
                    return (
                        <div>
                            <h1>Error</h1>
                            <p>Can't check profile. Please refresh the page</p>
                        </div>
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