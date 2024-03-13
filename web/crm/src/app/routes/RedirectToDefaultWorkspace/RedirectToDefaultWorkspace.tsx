import { Navigate } from "react-router-dom"

import { useAuthContext } from "shared/context/auth/context"

export function RedirectToDefaultWorkspace() {
    const { user } = useAuthContext()

    return <Navigate to={user.status === "LOGGED" ? `/${user.defaultWorkspaceId}` : "/login"} />
}
