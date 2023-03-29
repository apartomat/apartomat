import { Navigate } from "react-router-dom"

import useAuthContext from "../useAuthContext"

function RedirectToDefaultWorkspace() {
    const { user } = useAuthContext()

    return (
        <Navigate to={user.status === "LOGGED" ? `/${user.defaultWorkspaceId}` : "/login"} />
    )
}

export default RedirectToDefaultWorkspace