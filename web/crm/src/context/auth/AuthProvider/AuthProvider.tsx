import { ReactNode } from "react"

import { authContext, useAuthProvider } from "../useAuthContext"

export function AuthProvider({ children }: { children: ReactNode }) {
    const auth = useAuthProvider()

    return <authContext.Provider value={auth}>{children}</authContext.Provider>
}

export default AuthProvider
