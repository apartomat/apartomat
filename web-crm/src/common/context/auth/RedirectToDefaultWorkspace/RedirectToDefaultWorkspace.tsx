import React from "react"
import { Redirect, RouteProps} from "react-router-dom"

import useAuthContext from "../useAuthContext"

function RedirectToDefaultWorkspace({ children, ...rest }: RouteProps) {
    const { user } = useAuthContext();

    switch (user.status) {
        case "LOGGED":
            return (
                <Redirect to={{ pathname: `/${user.defaultWorkspaceId}`}}/>
            );
        default:
            return (
                <div>Error...</div>
            );
    }
}

export default RedirectToDefaultWorkspace;