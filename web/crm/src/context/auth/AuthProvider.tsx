import { ReactNode } from "react"

import { authContext, useAuthProvider } from "context/auth/context"

export function AuthProvider({ children }: { children: ReactNode }) {
    const auth = useAuthProvider()

    return <authContext.Provider value={auth}>{children}</authContext.Provider>
}
